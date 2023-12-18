package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleOauthConfig = oauth2.Config{
	RedirectURL:  "http://localhost:3000/auth/google/callback", // 구글에서 우리쪽으로 리다이렉트할 주소
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_SECRET_KEY"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

func googleLoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateStateOauthCookie(w)
	fmt.Printf("### state: %v\n", state)
	url := googleOauthConfig.AuthCodeURL(state)
	fmt.Printf("### url: %v\n", url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func generateStateOauthCookie(w http.ResponseWriter) string {
	// 쿠키 만료 시간 설정
	expiration := time.Now().Add(1 * 24 * time.Hour)

	// state값을 임의로 만들어서 설정 (CSRF 공격 방지)
	b := make([]byte, 16)
	rand.Read(b)
	fmt.Printf("### b: %v\n", b)
	state := base64.URLEncoding.EncodeToString(b)

	// 쿠키 생성
	cookie := http.Cookie{Name: "oauthstate", Value: state, Expires: expiration}
	http.SetCookie(w, &cookie)
	return state
}

func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")
	// 우리가 요청한 정보가 아님 (잘못된 요청)
	if r.FormValue("state") != oauthstate.Value {
		log.Printf("invalid google oauth state cookie: %s state: %s\n", oauthstate.Value, r.FormValue("state"))
		// 에러 정보를 띄우면, 해커가 어떤 문제인지 파악하기 쉬움
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprint(w, string(data))
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// 구글 로그인한 사용자 정보 request
func getGoogleUserInfo(code string) ([]byte, error) {
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to Exchange %s", err.Error())
	}
	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to Get UserInfo %s", err.Error())
	}

	return io.ReadAll(resp.Body)
}

func main() {
	mux := pat.New()
	mux.HandleFunc("/auth/google/login", googleLoginHandler)
	mux.HandleFunc("/auth/google/callback", googleAuthCallback)

	n := negroni.Classic() // public 폴더 자동으로 지원
	n.UseHandler(mux)
	http.ListenAndServe(":3000", n)
}
