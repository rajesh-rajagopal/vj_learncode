// https://gist.github.com/tomcatzh/5d1d0d9a95cecba798d1
// 
package main

import (
  "os"
  "bufio"
  "fmt"
)

func main() {
 lines, err :=  readLines("./test.txt")
  if err != nil {
   fmt.Println("Line read error : ",err)
    return 
  }
  lines = append(lines, "** Welcome to writer **")
  err = writeLines(lines, "./test2.txt")
  err = AppendLinesToFile("./test2.txt", []string{"** New appender **", "my name ", "is", "vijay"})
}

func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  fmt.Println("write success")
  return w.Flush()
}

/**
 * Append string to the end of file
 *
 * path: the path of the file
 * text: the string to be appended. If you want to append text line to file,
 *       put a newline '\n' at the of the string.
 */
func AppendStringToFile(path, text string) error {
      f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
      if err != nil {
              return err
      }
      defer f.Close()

      _, err = f.WriteString(text)
      if err != nil {
              return err
      }
      return nil
}



func AppendLinesToFile(path string, lines []string) error {
      f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
      if err != nil {
              return err
      }
      defer f.Close()
   for _, line := range lines{
      line = line + "\n"
      _, err = f.WriteString(line)
      if err != nil {
              return err
      }
   }
   fmt.Println("append success")
      return nil
}

