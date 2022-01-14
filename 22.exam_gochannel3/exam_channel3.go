package main
 
func main() {
    gch := make(chan int, 2)
     
    // 아래와 같이 Go 채널에 송신 함.
    gch <- 1
    gch <- 2
     
    // 해당 Go 채널을 닫음.
    close(gch)
 
    // Go 채널관련 수신
    println(<-gch)
    println(<-gch)
     
    if _, succ := <-gch; !succ {
        println("No data.")
    }
}