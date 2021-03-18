package main

func main() {
	//1. 익명함수를 변수에 할당한 후 이 변수를 전달하는 방법
	arg := func(x int, y int, z int) int {
		return x * y * z
	}

	res1 := myfunc(arg, 118, 119, 120)
	println(res1)

	//2. 직접 다른 함수 호출 파라미터에 함수를 적는 방법
	res2 := myfunc(func(i int, j int, k int) int { return i + j + k }, 118, 119, 120)
	println(res2)

}

func myfunc(f func(int, int, int) int, a int, b int, c int) int {
	res := f(a, b, c)
	return res
}
