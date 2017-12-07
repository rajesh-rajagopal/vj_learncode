package main 
 
import (
  "encoding/json"
  "fmt"
)

type Test struct {
  Name string 
  Age string 
  Address string 
  Phone   string  `json:"-"`
}

func main() {
  data := &Test{Name: "viajy",Age:"23",Phone:"1234697"}
	 jsonMsg, err := json.Marshal(data)
	 if err != nil {
	   fmt.Println(err)
	 }
 fmt.Println(string(jsonMsg))

}
