package main

import (
   "fmt"
   "os/exec"
   "os"
)

func main() {
    app := "/bin/ls"
    arg0 := "-l"
    cmd := exec.Command(app, arg0)
    stdout, err := cmd.Output()
    if (err != nil) {
       fmt.Fprintln(os.Stderr, err.Error())
       return
    }

    fmt.Println(string(stdout))

}

