package main

import "fmt"

func main() {
	fmt.Println("lets start with Array")
  var fruitlist [4] string
  fruitlist[0]="apple"
  fruitlist[1]="mango"
  fruitlist[2]="bannana"
  fruitlist[3]="pineapple"

  fmt.Println("fruit list is :=",fruitlist)
  fmt.Println("fruit list table length : =" ,len(fruitlist))

  var veglist =[3] string{"potato", "carrot","ladish finger"}
  fmt.Println("my vergeis list is :",veglist)

  // lets talk  about int
  var num[4]int
  fmt.Println(num)
  var num1=[4]int {1 ,2 ,3,4}
  fmt.Println(num1)
  // chang on the postion  of array index
  num1[0]=99
  fmt.Println(num1)
}
/*  declearion 
    var (variable) [] String
*/