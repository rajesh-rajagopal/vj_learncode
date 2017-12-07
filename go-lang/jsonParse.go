package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Name string  `json:"name"`
    Age  int     `json:"age"`
    Address []string  `json:"address"`
}

func main() {
    user := &User{Name: "Frank",
            Age: 20,
            Address: []string{"31b mettu st.,","idappadi","salem"},
          }
    b, err := json.Marshal(user)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(string(b))
}

