package main

import (
  "strings"
  "github.com/pivotal-golang/bytefmt"
  "fmt"
  "regexp"
  "strconv"
)


func main() {



fmt.Println(numMemory("6 GB"))



}

func ConnumMemory(bc string) uint64 {
	if cs, err := strconv.ParseUint(trimMemory(bc), 10, 64); err != nil {
               fmt.Println(err)
		return 0
	} else {
		return cs
	}
}

func numMemory(bc string) uint64 {
	if cp, err := bytefmt.ToMegabytes(strings.Replace(bc, " ", "", -1)); err != nil {
		return 0
	} else {
		return cp
	}
}


func trimMemory(bc string) string {
	coreRegex := regexp.MustCompile("[Bb].*")
	s := strings.TrimSpace(coreRegex.ReplaceAllString(bc, ""))
	return  strings.Replace(s, " ", "", -1)
}

