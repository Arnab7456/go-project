package main

import (
	"fmt"
	"math/rand"
	
	"time"
)

func main(){

  fmt.Print("welcome to go world")
/*
  var numberone int = 2;
  var numbertwo float64 = 4.5 ;
  fmt.Println("sum of this num" , numberone + int(numbertwo))
  // but where is 0.5 ?
*/
  // concept of random 
  
  rand.Seed(time.Now().UnixNano())
  // we discuss time and unix in next class
  fmt.Print(rand.Intn(5)+1)
  // why it only shows 1 because in machine language everythin based on algorithm ...so it shows 1 only 
  // rand = math function 

  
}