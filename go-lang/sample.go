package main

import (
  "fmt"
  "time"
"strconv"
)

func main() {
fmt.Println(time.Now().Add(-10*time.Minute))
fmt.Println(time.Now())
 Times := []string{"1482350912", "1482350955", "1482413597", "1482414664", "1482429197", "1482436831", "1482437109", "1482437430", "1482437649", "1482437902", "1482438211", "1482476071", "1482479677", "1482480135"}
for _,j := range Times {
    i, err := strconv.ParseInt(j, 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)
    fmt.Println("Start at :",tm)
}
}
