package main

import "fmt"

func main() {
	fmt.Println("lets discuss")
  myNumber := 20
  var ptr  = &myNumber
  println("found address",ptr);
  println("found the actual value",*ptr);

  *ptr = *ptr +2
  fmt.Println("new value is ",*ptr)
  *ptr = *ptr *2
  fmt.Println("multipication value ",*ptr)
}