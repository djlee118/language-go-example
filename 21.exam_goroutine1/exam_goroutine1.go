package main

import (
	"fmt"
	"runtime"
	"time"
)

func test(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s, "#", i)
	}
}

func main() {
	runtime.GOMAXPROCS(7) //7개의 CPU 사용

	// 동기적으로 실행됨.
	test("this is sync.")

	// 비동기적으로 실행됨.
	go test("async, one!")
	go test("async, two!!")
	go test("async, three!!!")

	// 3초 대기
	time.Sleep(time.Second * 3)
}
