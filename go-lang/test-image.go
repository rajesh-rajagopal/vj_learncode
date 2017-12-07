package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	cmd := "convert"
	args := []string{"-resize", "75%", "foo.jpg", "foo.half.jpg"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully halved image in size")
}
