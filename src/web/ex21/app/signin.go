package app

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type GoogleUserId struct {
	ID            string `json:"id"`
	Email         string `json:"email"`
	VerifiedEmail bool   `json:"verified_email"`
	Picture       string `json:"picture"`
}

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
	fmt.Printf("### url: %v\n", url) // https://accounts.google.com/o/oauth2/auth?client_id=...
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// cookie에 설정할 state값 생성 및 cookie 생성/등록
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
	http.SetCookie(w, &cookie) // 쿠키 등록
	return state
}

// '/auth/google/callback' URL로 콜백 요청
func googleAuthCallback(w http.ResponseWriter, r *http.Request) {
	oauthstate, _ := r.Cookie("oauthstate")
	// 우리가 요청한 정보가 아님 (잘못된 요청)
	if r.FormValue("state") != oauthstate.Value {
		orrMsg := fmt.Sprintf("invalid google oauth state cookie: %s state: %s\n", oauthstate.Value, r.FormValue("state"))
		log.Printf("%s", orrMsg)
		// 에러 정보를 띄우면, 해커가 어떤 문제인지 파악하기 쉬움
		http.Error(w, orrMsg, http.StatusInternalServerError)
		return
	}

	data, err := getGoogleUserInfo(r.FormValue("code"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Id 등 사용자 정보를 Session cookie에 저장
	var userInfo GoogleUserId
	err = json.Unmarshal(data, &userInfo)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, err := store.Get(r, "session")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// Set some session values.
	session.Values["ID"] = userInfo.ID
	session.Values["Email"] = userInfo.Email
	session.Values["Picture"] = userInfo.Picture

	// Save it before we write to the response/return from the handler.
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// 구글 로그인한 사용자 정보 request
func getGoogleUserInfo(code string) ([]byte, error) {
	// 인증 코드를 oauth2.Token으로 변환
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return nil, fmt.Errorf("failed to Exchange %s", err.Error())
	}

	// 사용자 정보 request URL에 access token값을 붙여서 request
	resp, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to Get UserInfo %s", err.Error())
	}

	return io.ReadAll(resp.Body)
}
