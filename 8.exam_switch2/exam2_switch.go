package main

import "fmt"

func checktype(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int, value[%v]\n", v)
	case bool:
		fmt.Printf("bool, value[%v]\n", v)
	case string:
		fmt.Printf("string, value[%v]\n", v)
	default:
		fmt.Printf("no type, value[%v]\n", v)
	}
}

func main() {
	checktype(100)
	checktype("magic")
	checktype(true)
}
