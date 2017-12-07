package main

import (
    "fmt"
    "io/ioutil"
     "log"
)

func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }

    for _, f := range files {
        if f.IsDir() {
          fmt.Println("Directroy : ",f.Name())
	  fmt.Println(f.Name())
	}
            
    }
}
