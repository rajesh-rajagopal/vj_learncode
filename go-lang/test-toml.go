package main

import(
 "github.com/BurntSushi/toml"
 "fmt"
)

type Config struct {
	Provider    string `toml:"provider"`
	OneEndPoint string `toml:"one_endpoint"`
	OneUserid   string `toml:"one_userid"`
	OnePassword string `toml:"one_password"`
	OneTemplate string `toml:"one_template"`
	OneZone     string `toml:"one_zone"`
	Certificate string `toml:"certificate"`
	Image       string `toml:"image"`
	VCPUPercentage string `toml:"vcpu_percentage"`
}

type Conf struct {
  Con Config `toml:"common"`
}

func main() {

var cm Conf
	if _, err := toml.Decode(`
[common]	
 one_endpoint = "http://opennebula:3000/xmlrpc2"
 one_userid   = "oneadmin"
 one_password = "password"
 one_template = "megam"
 one_zone     = "plano01"
 certificate = "/etc/ssl/cert.pem"


`, &cm); err != nil {
		fmt.Println(err)
	}
fmt.Println(cm)

}
