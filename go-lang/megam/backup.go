package main

import (
  "github.com/megamsys/libgo/api"
  "fmt"
)


func main() {
   args := api.ApiArgs{
    Email: "testadmin@astimp.ro",
    Url: "http://188.240.231.84:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
  }
  args.Org_Id = ""
  bid := "BAK6537005410413029772"
  aid := "ASM7073844791887679923"  
  //show(bid,args)
  list(args)
  //listAsm(aid,args)
}

func show(b string, api.ApiArgs) {
  cl := api.NewClient(args, "/backups/show/"+b)
  res, err := cl.Get()
  fmt.Println("Error :", err)
  fmt.Println("response :",string(res))


}

func list(c api.ApiArgs) {
  cl := api.NewClient(args, "/backups")
  res, err := cl.Get()
  fmt.Println("Error :", err)
  fmt.Println("List response :",string(res))


}

func listAsm(asm string,c api.ApiArgs) {
  cl := api.NewClient(args, "/backups/"+asm)
  res, err := cl.Get()
  fmt.Println("Error :", err)
  fmt.Println("List ASM response :",string(res))


}


