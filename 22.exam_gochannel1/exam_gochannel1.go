package main

func main() {
	// 문자열형 Go 채널을 생성.
	gch := make(chan string)

	go func() {
		gch <- "magic" //Go 채널에 "magic" 을 전송함.
	}()

	var chstr string
	chstr = <- gch // Go 채널로부터 "magic" 을 수신함.
	println(chstr)
}
