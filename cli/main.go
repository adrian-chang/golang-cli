package main

import (
  "fmt"
)

func main() {
  c := make(chan int)
  go func(chan int) {
    fmt.Println("Hello world")
    c <- 1
  }(c)
  <- c 
}
