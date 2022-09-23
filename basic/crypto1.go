package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

)
func main(){
  fmt.Print("lets start crypto command")
  myRandomnum,_ := rand.Int(rand.Reader, big.NewInt(5))

  fmt.Print(myRandomnum)
  
}