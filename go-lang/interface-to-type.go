package main

import "fmt"
import "reflect"

func main() {
    var data interface{}
    data = 4

     

    switch t := data.(type) {
    case int:
        fmt.Printf("Integer: %v\n", t)
    case float64:
        fmt.Printf("Float64: %v\n", t)
    case string:
        fmt.Printf("String: %v\n", t)
    case bool:
        fmt.Printf("Bool: %v\n", t)
    case []interface {}:
        for i,n := range t {
            fmt.Printf("Item: %v= %v\n", i, n)
        }
    default:
        var r = reflect.TypeOf(t)
        fmt.Printf("Other:%v\n", r)             
    }

}
