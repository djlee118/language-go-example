package main

func main() {
    var i interface{} = 118
 
    k := i       // i와 k 는 dynamic 타입이고, 값은 118 임.
    m := i.(int) // m 은 int 타입이고, 값은 118 을 나타냄.
 
    println(k)  // 포인터주소 출력하게 됨.
    println(m)  // 118 출력을 출력하게 됨.
}