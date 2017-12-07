//## to get all images size and add 
package main 

import (
   "fmt"
   "github.com/megamsys/opennebula-go/images"
   "github.com/megamsys/opennebula-go/api"
)

func main() {
var total int
  cl,err := api.NewClient(map[string]string{"password":"CepGerojFed0","userid":"oneadmin","endpoint":"http://213.32.30.57:2633/RPC2"})
  im := &images.Image{
   T: cl,
  }

  ims, err := im.List()
  fmt.Println("Error: ",err)
for _,v := range ims.Images {
  fmt.Println( v.Name ,"  - ", v.Size)
  total  = total + v.Size
}
fmt.Println("Total size of Images :", total)
}

