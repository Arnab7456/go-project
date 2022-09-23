package main

import (
	"fmt"
	"time"
)

func main() {
 fmt.Print("lets disscus time \n")
  arnabtime := time.Now()
  fmt.Println(arnabtime)

    // for india location only
  loc, _ := time.LoadLocation("Asia/Kolkata")
    now := time.Now().In(loc)
    fmt.Println("Location : ", loc, " Time : ", now)

  fmt.Println(arnabtime.Format("01-02-2006 15:04:05 Monday"))
// "01-02-2006 15:04:05 Monday" this whole command is present in dev.go documentation ===> dispaly the actual time of any regeion

  // lets create a date 
  createdate := time.Date(2020, time.July, 5,14,45,15,15,time.UTC)
  fmt.Println(createdate)
}