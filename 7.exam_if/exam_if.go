package main

import "fmt"

func what_company(cp string) string {
	if cp == "SK" {
		return "SK company"
	}

	return "default company"
}

func main() {
	fmt.Println(what_company("SK"))
	fmt.Println(what_company("TEST"))
}
