package main

import (
        "fmt"
        "log"
        "regexp"
        "strings"
)

func main() {
        reg, err := regexp.Compile("[^A-Za-z0-9]+")
        if err != nil {
                log.Fatal(err)
        }

        safe := reg.ReplaceAllString("a*-+fe5v9034,j*.AE6", "-")
        safe = strings.ToLower(strings.Trim(safe, "-"))
        fmt.Println(safe)   // Output: a*-+fe5v9034,j*.ae6
}
