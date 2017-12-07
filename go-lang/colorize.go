package main

import (
   "fmt"
)

func main() {

  fmt.Println(fmt.Sprintf("\033[38;5;%dm%s\033[0m", 2,"Hello"))
}
