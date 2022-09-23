package main

import 
	"fmt"


func main(){
  // here we use argonomisrtic type

  arr := [3]int{1,2,3}
  fmt.Println(arr)
  fmt.Println(len(arr))
  sum :=0
  for i:=0;i<len(arr) ; i++{
    sum += arr[i]
  }
  fmt.Println(sum)
  
}