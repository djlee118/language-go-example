package main

import "fmt"

func main() {
	s := []int{1, 3, 5, 7, 9, 0}
	s = s[2:6]

	fmt.Println(s) // 5,7,9,0 을 출력함.
}
