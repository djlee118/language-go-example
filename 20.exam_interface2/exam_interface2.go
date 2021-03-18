package main

import "fmt"

func main() {
	var i interface{}
	i = 118
	i = "magic"

	showValue(i)
}

func showValue(i interface{}) {
	fmt.Println(i) // magic 을 출력함.
}
