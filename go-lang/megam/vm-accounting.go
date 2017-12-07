//## to get all vms size and add 
package main 

import (
   "fmt"
  "encoding/xml"
   "github.com/megamsys/opennebula-go/metrics"
   "github.com/megamsys/opennebula-go/api"
  "time"
)

func main() {
var total int64
  cl,_ := api.NewClient(map[string]string{"password":"CepGerojFed0","userid":"oneadmin","endpoint":"http://213.32.30.57:2633/RPC2"})
  im := &metrics.Accounting{
  Api: cl,
  StartTime: time.Now().Add(-10* time.Minute).Unix(),
  EndTime: time.Now().Unix(),
  }

  ims, e := im.Get()
  ons := &metrics.OpenNebulaStatus{}
	e = xml.Unmarshal([]byte(ims.(string)), ons)
	if e != nil {
        fmt.Println("Error: ",e)
        return 
	}
 
for _,v := range ons.History_Records {
  fmt.Println( v.VM.Name ,"  - ", v.DiskSize())
  total  = total + v.DiskSize()
}
fmt.Println("Total size of VMs :", total)
}

