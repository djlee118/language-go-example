package main

import (
	"bufio"
	"encoding/csv"
	"os"
)

func main() {
	// CSV 파일을 생성함.
	file, err := os.Create("C:\\work\\go\\sample.csv")

	if err != nil {
		panic(err)
	}

	// CSV writer 를 생성함.
	wr := csv.NewWriter(bufio.NewWriter(file))

	// CSV 파일에 내용을 씀.
	wr.Write([]string{"magic1", "1"})
	wr.Write([]string{"magic2", "2"})

	wr.Flush()
}
