package main

import (
"regexp"
"fmt"
"strings"
)

func main() {

r := regexp.MustCompile("[Bb].*")
s := strings.TrimSpace(r.ReplaceAllString(" 1 GB ", ""))
s1 := strings.Replace(s, " ", "", -1)
fmt.Println(s)
fmt.Println(s1)
}
