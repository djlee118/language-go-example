package main

import (
	"fmt"
)

func main() {
	var p *int
	i := 42
	p = &i

	fmt.Println(*p) // p 포인터를 통해서 i 값을 읽어와서 출력
	*p = 21         // p 포인터를 통해서 i 값을 변경
	fmt.Println(i)
	fmt.Println(*p)

	// var pNum *int = new(int)

	// pNum++              // 포인터 연산은 허용하지 않으므로, 컴파일 에러가 남.
	// pNum = 0xc0999962d0 // 포인터에 메모리 주소를 직접 대입할 수 없으므로, 컴파일 에러가 남.

	// fmt.Println(pNum)
}
