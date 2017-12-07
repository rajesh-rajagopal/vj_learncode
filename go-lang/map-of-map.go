package main

import (
    "fmt"
)

func main() {
    myData := make(map[string]map[string]string)
    path := "first"
    mm, ok := myData[path]
    if !ok {
        mm = make(map[string]string)
        mm["sub-first"] = "hello"
        myData[path] = mm
    }

    path = "second"
    mm2, ok := myData[path]
    if !ok {
        mm2 = make(map[string]string)
        mm2["sub-second"] = "megam"
        myData[path] = mm2
    }

    for k, _ := range myData {
        fmt.Println(k, ":")
        for _, j := range myData[k] {
            fmt.Println(j)
        }
    }

}
