package main 

import (
  "github.com/megamsys/libgo/api"
 // "github.com/megamsys/libgo/pairs"
 // "encoding/json"
  "fmt"
)

func main() {
   args := api.ApiArgs{
    Email: "vino@gmail.co",
    Url: "http://13.73.206.106:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
  }
  id := "MKT4949751510137504867" 
  res, err := testGet(args,"/marketplaces/" + id)
  fmt.Println("res :", string(res))
  fmt.Println("error",err)

}

func testGet(args api.ApiArgs, path string) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Get()
}
