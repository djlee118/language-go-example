package main

import "fmt"

func main() {
	users := map[string]string{
		"rank1": "magic1",
		"rank2": "magic2",
		"rank3": "magic3",
	}

	// map 의 key를 체크해 보자.
	val, exists := users["rank5"]

	if !exists {
		println("no rank5 user.")
	} else {
		println("val[" + val + "]")
	}

	for key, val := range users {
		fmt.Println(key, val)
	}

	delete(users, "rank2")

	for key, val := range users {
		fmt.Println(key, val)
	}

}
