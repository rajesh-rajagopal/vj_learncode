package main

import (
	"bytes"
	"encoding/json"
  "fmt"
	// "net/http"
	// "runtime"
)

// Encodes the response headers into JSON format.
func encodeResponse(response interface{}) []byte {
	var bytesBuffer bytes.Buffer
//	bytesBuffer.WriteString(json.Header)
	e := json.NewEncoder(&bytesBuffer)
	e.Encode(response)
	return bytesBuffer.Bytes()
}

type Person struct {
  Name string `josn:"name"`
  Address []string `json:"address"`
}

func main() {
  p := &Person{
    Name: "krish",
    Address: []string{"madhura","yamunai street"},
  }
  fmt.Println(string(encodeResponse(p)))
}
