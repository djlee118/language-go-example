package main

import "fmt"

func main() {
	defer fmt.Println("magic")

	fmt.Print("T.hub master : ")
}
