package main

func main() {
	calc := func(n ...int) int { //익명함수를 정의함.
		num := 0

		for _, i := range n {
			num += i
		}

		return num
	}

	res := calc(118, 119, 120) //익명함수를 호출함.

	println(res)
}
