package main
 
import (
    "bufio"
    "encoding/csv"
    "fmt"
    "os"
)
 
func main() {
   
    // 기존에 작성했던, CSV 파일(sample.csv)을 오픈함.
    file, _ := os.Open("C:\\work\\go\\sample.csv")
 
    // CSV reader 를 생성함.
    rdr := csv.NewReader(bufio.NewReader(file))
 
    // sample.csv 파일의 내용을 모두 읽음.
    rows, _ := rdr.ReadAll()
 
    // for loop 를 사용해서 행,열 읽음.
    for i, row := range rows {
        for j := range row {
            fmt.Printf("%s ", rows[i][j])
        }
        fmt.Println()
    }
}
