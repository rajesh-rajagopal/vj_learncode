package main

import (
 "fmt"
 "strings"
)

func main() {

s := []string{"hello"}
s2 := "hello2"

fmt.Println("s- {\"hello\"} len(s) -",len(s))
fmt.Println("s2 - \"hello2\" len(s2)-",len(s2))
//fmt.Println("len (strings.TrimSpace(s))    - ",len(strings.TrimSpace(s)))
fmt.Println("len(strings.TrimSpace(s2))  - ",len(strings.TrimSpace(s2)))
}

