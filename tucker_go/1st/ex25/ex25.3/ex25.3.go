package main

import (
	"fmt"
	"sync"
	"time"
)

type Car struct {
	Body  string
	Tire  string
	Color string
}

var wg sync.WaitGroup
var startTime = time.Now()

func main() {
	// [채널로 생산자 소비자 패턴 구현하기]
	// - 고루틴에서 뮤텍스 사용하지 않는 방법
	//   1) 영역 나누기
	//   2) 역할 나누기

	// Ex) 채널을 이용해서 '역할'을 나누는 방법

	tireCh := make(chan *Car)
	paintCh := make(chan *Car)

	fmt.Printf("Start Factory\n")

	wg.Add(3)
	// 고루틴 생성
	go MakeBody(tireCh)
	go InstallTire(tireCh, paintCh)
	go PaintCar(paintCh)

	wg.Wait()
	fmt.Println("Close the factory")
}

// 차체 생산
func MakeBody(tireCh chan *Car) {
	tick := time.Tick(time.Second)
	after := time.After(10 * time.Second)

	for {
		select {
		case <-tick:
			// 차체 생산 작업
			car := &Car{}
			car.Body = "포르쉐"
			tireCh <- car

		case <-after:
			close(tireCh)
			wg.Done()
			return
		}
	}
}

// 바퀴 설치
func InstallTire(tireCh, paintCh chan *Car) {
	for car := range tireCh { // for range는 채널이 닫히는 것을 감지하면 종료
		// 바퀴 설치 작업
		time.Sleep(time.Second * 1)
		car.Tire = "스노우 타이어"
		paintCh <- car
	}
	wg.Done()
	close(paintCh)
}

// 도색
func PaintCar(paintCh chan *Car) {
	// 도색 작업
	for car := range paintCh {
		time.Sleep(time.Second * 1)
		car.Color = "빨강"

		// 경과 시간 출력
		// duration := time.Now().Sub(startTime)
		duration := time.Since(startTime)
		fmt.Printf("%.2f Complete Car: %s %s %s\n", duration.Seconds(), car.Body, car.Tire, car.Color)
	}
	wg.Done()
}

// [생산자 소비자 패턴]
// - 한 쪽에서 데이터를 생성해서 넣어주면, 다른 쪽에서 생성된 데이터를 빼서 사용

// 예제에서 살펴보면...
// 1) 바퀴 설치 작업 시
//   - MakeBody()루틴이 생산자, InstallTire()루틴은 소비자
// 2) 도색 작업 시
//   - InstallTire()루틴은 생산자, PaintCar()루틴은 소비자
