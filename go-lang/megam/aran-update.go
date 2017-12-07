package main

import (
 "fmt"
 "net/http"
  "io/ioutil"
 "gitlab.com/rioos/k8s.io/api/core/v1"
  "encoding/json"
  "bytes"
)
var client *http.Client
var headers http.Header
func main() {

client = &http.Client{}
headers := http.Header{}
  headers.Set("Authorization", "Bearer kmKPcCfkpTbUATNBLv")
  headers.Set("X-AUTH-RIOOS-EMAIL", "info@megam.io")

 api := "http://192.168.1.100:9636/api/v1/"
 suffix := "assemblys/812713345915305984"
path := api + suffix
//body := bytes.NewReader(data)
body, err := get(path)
if err != nil {
    fmt.Printf("Error : %s", err)
   return
}
err = put(path, body) 
if err != nil {
    fmt.Printf("Error : %s", err)
   return
}
}

func get(path string) ([]byte, error) {
fmt.Println("URL [GET] : ",path)
var body []byte{}
req, err := http.NewRequest("GET", path, nil)
req.Header = headers
resp, err := client.Do(req)
if err != nil {
    fmt.Printf("Error : %s", err)
    return body, err
}
defer resp.Body.Close()
body, err = ioutil.ReadAll(resp.Body)
if err != nil {
    fmt.Printf("GET Error : %s", err)
    return body, err
}
fmt.Println("\n\nresponse  :")
fmt.Printf("%#v  ",string(body))
return body, nil
}

func put(path string,body []byte) error {
fmt.Println("URL [GET] : ",path)
req.Header = headers
reqbody = bytes.NewReader(body)
req, err := http.NewRequest("PUT", path, reqbody )
resp, err := client.Do(req)
if err != nil {
    fmt.Printf("PUT Error : %s", err)
    return err
}
return nil
}

func unmarshal(body []byte) (v1.Assembly, error) {
   asm := &v1.Assembly{}
   err := json.Unmarshal(body, asm)
if err != nil {
    fmt.Printf("Unmarshalling Error : %s", err)
}
return asm, err
}


