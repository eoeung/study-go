package main

import "fmt"

func main() {
	// 예제로 배우는 GO 프로그래밍
	// [채널 닫기]
	//   - 채널을 오픈한 다음, close(ch)를 사용해서 채널을 닫음
	//   - 채널을 닫으면 채널로 데이터를 보내는 건 안되지만, 수신은 가능
	//   - <-ch는 2개의 리턴값을 가짐
	//       1) 채널 메시지
	//       2) 수신 성공 여부

	ch := make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다.
	close(ch)

	// 채널 수신
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	if msg, success := <-ch; !success {
		fmt.Println(msg) // 0 (데이터가 없어서 기본값으로 출력된 듯)
		fmt.Println("더 이상 채널에 데이터가 없음")
	}

	// go by example
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		// [워커 고루틴] 1) jobs 채널(메인 고루틴 → 다른 고루틴) 값을 반복문을 돌려서 모두 출력
		// value, isSuccess := <- channel
		for {
			j, more := <-jobs
			// [워커 고루틴] 2) 받아온 3개의 job을 출력
			if more {
				fmt.Println("received job", j)
			} else { // [워커 고루틴] 3) 받아온 3개의 job을 모두 꺼낸 경우, false반환
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// [메인 고루틴] 1) jobs 채널을 통해서 워커 고루틴으로 3개의 job을 보냄
	// [jobs 채널] 1) 3개의 값을 받음 (1,2,3)
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// [메인 고루틴] 2) jobs 채널 종료
	close(jobs)
	fmt.Println("sent all jobs")
	fmt.Println("tttttttteeeeeesssssssstttttttt")

	<-done

	fmt.Println("abcdefgh")

	/*
		sent job 1
		sent job 2
		sent job 3
		sent all jobs
		tttttttteeeeeesssssssstttttttt
		received job 1
		received job 2
		received job 3
		received all jobs
		abcdefgh
	*/

	zzz := make(chan int, 2)
	zzz <- 1
	close(zzz)
	// zzz <- 2 // 에러발생 - panic: send on closed channel // 닫힌 채널로 값을 전송할 수 없음
}
