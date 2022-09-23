package main

import (
	// "bufio"
	"fmt"
	// "os"
)

func main (){
  var num1 int
  var num2 int 
  fmt.Println("enter the first number")
        fmt.Scan(&num1)
  fmt.Println("enter the seconed number")
        fmt.Scan(&num2)
  
  fmt.Printf("the value is %d\n" , num1 + num2)
  fmt.Printf("the value is %d\n" , num1 - num2)
  fmt.Printf("the value is %d\n" , num1 * num2) 
  fmt.Printf("the value is %d\n" , num1 / num2)
}