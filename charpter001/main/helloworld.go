package main

import (
  "fmt"
  "time"
)

func main(){
  fmt.Println("Hello World!")
  var a chan string = make(chan string,2)
  var b chan string = make(chan string,2)
  var c chan string = make(chan string,2)
  c <- "C"
  time.Sleep(3*time.Second)
  a <- "A"
  time.Sleep(3*time.Second)
  b <- "B"
  for {
    //按成绩分成ABCD等级
    select {
    case a <- "A":
      fmt.Println("等级A")
    case b <- "B":
      fmt.Println("等级B")
    default:
      fmt.Println("等级C")
    }
  }
}
