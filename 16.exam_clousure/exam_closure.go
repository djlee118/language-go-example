package main
 
func nextVal() func() int {
    num := 0
    return func() int {
        num++
        return num
    }
}
 
func main() {
    next1 := nextVal()
 
    println(next1())  // 1
    println(next1())  // 2
    println(next1())  // 3
 
    next2 := nextVal()
    println(next2()) // 1 다시 시작 됨.
    println(next2()) // 2
    println(next2()) // 3
}