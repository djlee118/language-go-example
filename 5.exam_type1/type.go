package main

import (
	"fmt"
	"math/cmplx"
)

var (
	NickName string     = "magic"
	YesNo    bool       = true
	MaxInt   uint64     = 1<<64 - 1
	Zzz      complex128 = cmplx.Sqrt(-2 + 10i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", NickName, NickName)
	fmt.Printf("Type: %T Value: %v\n", YesNo, YesNo)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", Zzz, Zzz)
}
