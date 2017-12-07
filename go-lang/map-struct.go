package main 

import (
"fmt"
)

type group struct {
  resource string 
  version string 
}

var testmap map[group]bool 

func main() {

 testmap = map[group]bool{group{resource:"v1"}: true, group{resource:"v2"}: false}
 fmt.Println(testmap)
 if testmap[group{resource: "v1", version: ""}] {
   fmt.Println("ok")
 } else { fmt.Println("not ok")}
}
