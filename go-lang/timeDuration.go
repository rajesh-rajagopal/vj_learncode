package main
 
import (
"time"
"fmt"
)

func main() {

a := 1*time.Hour
b := 13*time.Minute

fmt.Println("Diff Sec :",a.Seconds()/b.Seconds())
fmt.Println("Diff Min :",a.Minutes()/b.Minutes())

}
