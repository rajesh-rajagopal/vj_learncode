package main
import (
	"os"
	"log"
	//"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	cmd := `#!/bin/sh
/usr/bin/passwd test <<EOF
hello
hello
EOF
`
	d1 := []byte(cmd)
	err := ioutil.WriteFile("dat1", d1, 0755)
	_, err = exec.Command("./dat1").Output()
	if err != nil {
			log.Fatal(err)
	}
	err = os.Remove("./dat1")
}
