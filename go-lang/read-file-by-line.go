// https://gist.github.com/tomcatzh/5d1d0d9a95cecba798d1
package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
  readLine("./test.txt")
}
func readLine(path string) {
  inFile, _ := os.Open(path)
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
}
