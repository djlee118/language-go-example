package main
 
import (
    "log"
    "os"
)
 
func main() {
    f, err := os.Open("C:\\log\\process_20201112.log")

    if err != nil {
        log.Fatal(err.Error())
    }

    println(f.Name())
}