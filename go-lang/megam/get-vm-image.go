package main

import (
   onevm  "github.com/megamsys/opennebula-go/virtualmachine"
  "github.com/megamsys/opennebula-go/api"
  "strconv"
  "fmt"
)

func main() {
	cm := make(map[string]string)
	cm[api.ENDPOINT] = "http://88.99.204.132:2633/RPC2"
	cm[api.USERID] = "oneadmin"
	cm[api.PASSWORD] = "virtengine"
 for i := 1; i <= 315; i++ {
	client, _ := api.NewClient(cm)
 	v := onevm.Vnc{T: client, VmId: strconv.Itoa(i)}
	vm, err := v.GetVm()
        if err != nil {
          fmt.Println(i," : [ERROR] : ",err)
        }
        ip := vm.GetVMIP()  
        if (ip == "88.99.200.185" || ip == "88.99.200.174") {
          fmt.Println("vm id : ",vm.Id,"  ",ip, "  : ",vm.VmTemplate.Disk.Image)
        }
 }

}
