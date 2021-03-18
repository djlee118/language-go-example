package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("defer main")

	println("start> main func")
	magicFunc()
	println("end> main func")
}

func magicFunc() {
	defer fmt.Println("defer magicFunc")

	defer func() {
		if i := recover(); i != nil {
			fmt.Println("recover!")
		}
	}()

	fmt.Println("call> main_func")
	panic("call panic!")

}
