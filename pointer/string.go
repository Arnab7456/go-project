package main

import "fmt"

func changevalue(str *string) {
	*str = "changed"
}
func change(str string) {
	str = "change!"
}
func main() {
	tochange := "Arnab"
	fmt.Println(tochange)
	change(tochange)
  fmt.Println(tochange)
	changevalue(&tochange)
  fmt.Println(tochange)

  /*
    here in changevalue func ====> pass a pointer in the case of go lang we pass pointer in the main function so lets discuss the another aspect of 
  what happern with the change func ====> here we don't pass any pointer so a copy of the the string is pass on func main () 
    also we declear a value od tochange = "Arnab" 
    so change change function create a copy of tochange 
    and we found
    Arnab
    Arnab
    changed
*/
    values := "Arnab this side"
    var pointer *string = &values
  fmt.Println(*pointer)
  fmt.Println(pointer,&pointer)
}