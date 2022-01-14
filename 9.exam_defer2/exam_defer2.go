package main

import "fmt"

func main() {
	fmt.Println("magic number, counting!")

	for i := 1; i <= 15; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("main func, done")
}
