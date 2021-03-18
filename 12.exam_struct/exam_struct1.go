package main

import (
	"fmt"
)

//User - struct 정의
type User struct {
	nickname string
	rank     int
}

func main() {

	u1 := User{}
	u1.nickname = "magic1"
	u1.rank = 1

	u2 := User{"magic2", 2}
	u3 := User{nickname: "magic3", rank: 3}
	u4 := new(User)
	u4.nickname = "magic4"
	u4.rank = 4

	fmt.Printf("u1 nickname[%s] rank[%d]\n", u1.nickname, u1.rank)
	fmt.Printf("u2 nickname[%s] rank[%d]\n", u2.nickname, u2.rank)
	fmt.Printf("u3 nickname[%s] rank[%d]\n", u3.nickname, u3.rank)
	fmt.Printf("u4 nickname[%s] rank[%d]\n", u4.nickname, u4.rank)

}
