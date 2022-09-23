package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
  // arnab := bufio.NewReader(os.Stdin)
	fmt.Println("rate our work")
    // coma ok syntex also called error ok  syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thank you for rating ", input)
  fmt.Printf("thank you for rating %T ", input)

}
