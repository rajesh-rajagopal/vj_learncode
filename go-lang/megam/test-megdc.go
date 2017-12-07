package main

import (
 "github.com/megamsys/megdc/handler"
 "fmt"
)

type BaseHost2 struct {
	Host       string            `json:"ipaddress"`
	Username   string            `json:"username"`
	Password   string            `json:"password"`
	HostInfo  bool   `json:"hostinfo"`
	PrivateKey string            `json:"privatekey"`
	Inputs  map[string]string    `json:"inputs"`
}

type BaseHost struct {
	Host string  `json:"ipaddress"`
	Port string  `json:"port"`
	Username string `json:"username"`
	SSHKey
	SSHPassword
	IsRan bool    //to verify can cleanup or not
}

type SSHKey struct {
	PrivateKey string `json:"privatekey"`
	PublicKey  string `json:"publickey"`
}

type SSHPassword struct {
  Password string `json:"password"`
}

type HostInfo struct {
 BaseHost2
 HostInfo  bool   `json:"hostinfo"`
}

func main() {
        //h := &HostInfo{BaseHost: BaseHost{Host: "127.0.0.1", Username: "vijay",SSHPassword: SSHPassword{Password: "speed"}},HostInfo: true}
        h := &BaseHost2{Host: "127.0.0.1", Username: "vijay",Password: "speed",HostInfo: true}
	f := handler.NewWrap(h)
        fmt.Println("New Wraper")
        fmt.Printf("%#v",f)
}
