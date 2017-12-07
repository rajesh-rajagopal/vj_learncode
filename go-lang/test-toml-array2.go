package main

import(
 "github.com/BurntSushi/toml"
  "encoding/json"
 "fmt"
 "time"
)


type One struct {
 Name string `toml:"Name"`
 Server server `toml:"alpah"`
}

type server struct {
    IP     string       `toml:"ip,omitempty"`
    Config []serverConfig `toml:"config"`
}

type serverConfig struct {
    Ports    []int `toml:"Ports"`
    Location string `toml:"Location"`
    Created  time.Time `toml:"Created"`
    Net      []net `toml:"net"`
}

type net struct {
  Name string `toml:"Name"`
  Sub []sub `toml:"sub"`
}

type sub struct {
 ID string `toml:"ID"`
}

func main() {


var tomlBlob = `
# Some comments.
title = "TOML Example"
[one]
welcome= "megam"
[one.owner]
name = "Tom Preston-Werner"
organization = "GitHub"
[one.owner.database]
server = "192.168.1.1"
[one.owner.database.servers]
name = "server1"
[one.owner.database.servers.alpha]
dc = "eqdc10"
ip = "10.0.0.1"
[one.owner.database.servers.beta]
dc = "eqdc10"
ip = "10.0.0.2"
name = "server1"
[one.owner.database.servers2]
name= "server2"
[one.owner.database.servers2.alpha]
dc = "eqdc10"
ip = "10.0.0.1"
[one.owner.database.servers2.beta]
dc = "eqdc10"
ip = "10.0.0.2"
[one.owner.database2]
server = "192.168.1.1"
[one.owner.database2.servers]
name= "server1"
[one.owner.database2.servers.alpha]
dc = "eqdc10"
ip = "10.0.0.1"
[one.owner.database2.servers.beta]
dc = "eqdc10"
ip = "10.0.0.2"
[one.owner.database2.servers2]
name= "server2"
[one.owner.database2.servers2.alpha]
dc = "eqdc10"
ip = "10.0.0.1"
[one.owner.database2.servers2.beta]
dc = "eqdc10"
ip = "10.0.0.2"
`

type servers map[string]interface{}

var config servers
//var h One
if _, err := toml.Decode(tomlBlob, &config); err != nil {
    fmt.Println(err)
}

jsonString, _ := json.Marshal(config)

fmt.Println("\n\n********",jsonString,"\n\n***********")

fmt.Printf("%#v",config)


}

