package main

import "fmt"

func main() {
	var i []int              //슬라이스 변수 i 선언
	i = []int{1, 2, 3, 4, 5} //슬라이스에 해당 리터럴값 지정함.
	i[1] = 10                //슬라이스 내용 변경
	fmt.Println(i)           //슬라이스 내용 출력
}
