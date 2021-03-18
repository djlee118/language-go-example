package main

import "fmt"

func main() {
	sum := 1
	for sum < 11 {
		sum += sum
	}
	fmt.Println(sum)
}
