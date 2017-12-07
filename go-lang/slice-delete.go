
package main 

import (
  "fmt"
  "strings"
)

func main() {
  s1 := []int{1,2,3,4,5}
    fmt.Println("original int [] :",s1)
  fmt.Printf("\n After remove %d element  %v", 3, removeInt(s1,3))
  s2 := []string{"hi","my","name","is","Vijayakanth"}
  fmt.Println("\n\noriginal string [] :",s2)
  fmt.Printf("\n After remove %d element  %v", 3, strings.Join(removeString(s2, 3)," "))
  fmt.Println()
}

func removeString(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func removeInt(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}
