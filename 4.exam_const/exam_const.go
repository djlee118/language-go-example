package main

import "fmt"

func main() {
	const LANG = "go"

	const NUM int = 30
	const STR string = "Hello"
	const MYVAL = 118

	const (
		SK      = "SK"
		SAMSUNG = "SAMSUNG"
		LG      = "LG"
	)

	const (
		JAVA   = iota // 0
		GO            // 1
		PYTHON        // 2
	)

	fmt.Println("Study Language: ", LANG)
	fmt.Println("My Favorite Number: ", MYVAL)

	fmt.Println("Company one:", SK)
	fmt.Println("Company two:", SAMSUNG)
	fmt.Println("Company three:", LG)

	fmt.Println("JAVA:", JAVA)
	fmt.Println("GO:", GO)
	fmt.Println("PYTHON:", PYTHON)

}
