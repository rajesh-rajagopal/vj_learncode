package main
import (
"fmt"
"encoding/json"

)

type Hello struct {
 Name string
 Address map[string]string
 Age int
}

func main() {
  h := &Hello{
  Name: "vijay",
  Age: 25,
  Address: map[string]string{"street":"31b kattuvalavu","town":"idappadi","district":"salem"},
  }
  
  	data, err := json.Marshal(h)
	if err != nil {
          fmt.Println(err)
	}
	h1 := &Hello{}
 	err = json.Unmarshal(data, h1)
	if err != nil {
          fmt.Println(err)
	}
	fmt.Println("Hello    :",h1)
}
