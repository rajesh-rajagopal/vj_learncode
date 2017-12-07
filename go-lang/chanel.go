package main

import (
    "fmt"
    "time"
    )
var i int = 0

    func put(c chan string){
     for {
        c <- "Tadaa!"
       }
      }

    func print(c chan string){
      for {
        message := <- c
        fmt.Println(message)
        time.Sleep(time.Second * 2)
        }
       }
     func main() {
       var c chan string = make(chan string)

       go put(c)
       go print(c)

       var input string
       fmt.Scanln(&input)
      }

