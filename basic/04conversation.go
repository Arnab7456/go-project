package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
  func main(){
    fmt.Print("welcome to our pizza shop")
    fmt.Println("rate our pizza between 1 to 5")
    reader :=bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')

      fmt.Println("thank you for rating " , input)

    numRating, err := strconv.ParseFloat(strings.TrimSpace(input),64)

    if err != nil {
      fmt.Println(err)
    } else {
        fmt.Println("added one in rating " , numRating + 1)
      }
    
  }