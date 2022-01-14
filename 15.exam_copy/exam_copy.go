package main

import "fmt"

func main() {
	src := []int{118, 119, 120}
	tgt := make([]int, len(src), cap(src)*2)
	copy(tgt, src)
	fmt.Println(tgt)            // "118, 119, 120" 출력
	println(len(tgt), cap(tgt)) // 3, 6 출력
}
