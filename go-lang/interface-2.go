package main

import "fmt"

func weird(i int) interface{} {
    if i < 0 {
        return "negative"
    }
    return i
}

func main() {
    var i = 42
    if w, ok := weird(7).(int); ok {
        i += w
        fmt.Println("given value is i =", i)
    }
    if w, ok := weird(-100).(string); ok {
        fmt.Println("given value is i= ",w)
    }
    
}

