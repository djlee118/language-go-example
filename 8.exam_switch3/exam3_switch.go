package main

import "fmt"

func main() {
	casecheck(2)
}

func casecheck(val int) {
	switch val {
	case 1:
		fmt.Println("case 1")
		fallthrough
	case 2:
		fmt.Println("case 2")
		fallthrough
	case 3:
		fmt.Println("case 3")
		fallthrough
	default:
		fmt.Println("case default")
	}
}
