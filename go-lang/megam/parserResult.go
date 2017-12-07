package main

import (
//  "errors"
	//"os/exec"
  "fmt"
	"strconv"
	"strings"
  "io/ioutil"
  //"os"
)

func main() {
  fmt.Println("************************")
  t := "space"
  raw := readFile("temp")

  lines := strings.Split(strings.TrimSpace(string(raw)), "\n")
  fmt.Println(lines)
	for _, line := range lines[1:] {
		if !strings.HasPrefix(line, "/") {
			continue
	}
		chunks := strings.Fields(line)
    fmt.Println("************************")
    fmt.Println(chunks)
		if len(chunks) >= 6 {
			tags := map[string]string{
				"file_system": chunks[0],
				"mounted_on":  chunks[5],
			}
      if v, e := strconv.ParseInt(chunks[1], 10, 64); e == nil {
				fmt.Println(t+".Total", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[2], 10, 64); e == nil {
				fmt.Println(t+".Used", v, tags)
			}
			if v, e := strconv.ParseInt(chunks[3], 10, 64); e == nil {
				fmt.Println(t+".Available", v, tags)
			}
			if v, e := strconv.ParseInt(strings.Replace(chunks[4], "%", "", 1), 10, 64); e == nil {
				fmt.Println(t+".Use", v, tags)
			}
    }
 }
}

func readFile(path string) []byte {
	b, e := ioutil.ReadFile(path)
	if e != nil {
		panic(e.Error())
	}
	return b
}
