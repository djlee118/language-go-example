package main

func main() {
	var mArray = [3][3]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9}, //그리고, 끝에 콤마를 추가해 준다.
	}
	println(mArray[2][1])
}
