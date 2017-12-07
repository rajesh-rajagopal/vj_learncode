package main

import (
    "fmt"
    "io/ioutil"
    lib "github.com/libvirt/libvirt-go"
)

// this is a comment

func main() {
    fmt.Println("Hello World")
    CreateDomain()
}

func buildConnection() *lib.Connect {
	conn, err := lib.NewConnect("qemu:///system")
	if err != nil {
	  fmt.Println(err)
	}
	return conn
}

func buildDomain() (*lib.Domain, *lib.Connect) {
	conn := buildConnection()
        dat, err := ioutil.ReadFile("/home/rajthilak/images/test.xml")
fmt.Println(string(dat))
	dom, err := conn.DomainDefineXML(string(dat))
	if err != nil {
		fmt.Println(err)
	}
	return dom, conn
}

func CreateDomain() {
	dom, conn := buildDomain()
	defer func() {
		dom.Free()
		if res, _ := conn.Close(); res != 0 {
		   fmt.Printf("Close() == %d, expected 0", res)
		}
	}()
	if err := dom.Create(); err != nil {
		fmt.Println(err)
		return
	}
}
