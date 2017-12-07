// https://github.com/QuentinPerez/go-radosgw



package main

import (
"bytes"
	"encoding/json"
	"github.com/ceph/go-ceph/rados"
	"github.com/ceph/go-ceph/rbd"
	"os/exec"
	"sort"
)

func main() {
	conn, _ := rados.NewConn()
	conn.ReadDefaultConfigFile()
	conn.Connect()
	ioctx, err := conn.OpenIOContext("one")
	if err != nil {
	  fmt.Println("Error  :", err)
        }
         
  imageNames, err := rbd.GetImageNames(ioctx)
 	if err != nil {
	  fmt.Println("Error  :", err)
        } else {
		  fmt.Println("Success  :" ,imageNames)
		}

}

