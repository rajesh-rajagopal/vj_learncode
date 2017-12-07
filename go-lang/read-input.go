package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    
    fmt.Println("Enter your next text :")
    var input string
    fmt.Scanln(&input)
    fmt.Print(input)

}
