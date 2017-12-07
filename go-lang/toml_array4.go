
package main

import (
 "github.com/BurntSushi/toml"
 "fmt"
 "io/ioutil"
 "os"
 //"reflect"
 //"encoding/json"
//    "time"
)

// deployd.one.region.cluster.vnet

type Config struct {
 Provider    string `json:"provider" toml:"provider"`
 Deployd     deployd `json:"deployd" toml:"deployd"`
}

type deployd struct {
Name string  `json:"name" toml:"name"`
One one      `json:"one" toml:"one"`
}

type one struct {
Enabled bool    `json:"enabled" toml:"enabled"`
Region []region `json:"region" toml:"region"`
OneMasterZone string `json:"one_master_zone" toml:"one_master_zone"`
OneMasterEndpoint string `json:"one_master_endpoint" toml:"one_master_endpoint"`
OneMasterUserid     string `json:"one_master_user" toml:"one_master_user"`
OneMasterPassword string `json:"one_master_password" toml:"one_master_password"`
}

type region struct {
OneZone     string `json:"one_zone" toml:"one_zone"`
OneEndpoint    string `json:"one_endpoint" toml:"one_endpoint"`
OneUserid        string `json:"one_user" toml:"one_user"`
OnePassword    string `json:"one_password" toml:"one_password"`
OneTemplate string `json:"one_template" toml:"one_template"`
Certificate string `json:"certificate" toml:"certificate"`
Cluster []cluster      `json:"cluster" toml:"cluster"`
cluster
}

type cluster struct {
Enabled bool    `json:"enabled" toml:"enabled"`
ClusterId string  `json:"cluster_id" toml:"cluster_id"`
Vnet_pri_ipv4 string `json:"vnet_pri_ipv4" toml:"vnet_pri_ipv4"`
Vnet_pub_ipv4 string `json:"vnet_pub_ipv4" toml:"vnet_pub_ipv4"`
Vnet_pri_ipv6 string `json:"vnet_pri_ipv6" toml:"vnet_pri_ipv6"`
Vnet_pub_ipv6 string `json:"vnet_pub_ipv6" toml:"vnet_pub_ipv6"`
}


func main() {

path := "/home/alrin/code/megam/directory/vertice/vertice.conf"
f, err := os.Open(path)
    if err != nil {
        panic(err)
}
defer f.Close()
buf, err := ioutil.ReadAll(f)
if err != nil {
  panic(err)
}
var config Config

if err := toml.Unmarshal(buf, &config); err != nil {
  panic(err)
}
/*
var a interface{}

a = config.Deployd.One
fmt.Println("**********interface to json**************************")
	b, err := json.Marshal(a)
if err != nil {
	fmt.Println("error:", err)
}
os.Stdout.Write(b)
fmt.Println("**********interface to struct**************************")
   if w, ok := a.(one); ok {
            fmt.Println("given value is i =", w)
    }

*/

c := region{ 
cluster: cluster{ClusterId: "100",},
}

fmt.Println(c)

fmt.Println("\n\n one first region   :", config.Deployd.One.Region[0].OneZone)
fmt.Println("\n\n one first region clusters   :",config.Deployd.One.Region[0].Cluster)
fmt.Println("\n\n one second region clusters   :",config.Deployd.One.Region[1].Cluster)
fmt.Println("\n\n one first region first cluster first vnet   :",config.Deployd.One.Region[0].Cluster[0].ClusterId)
fmt.Println("\n\n Length of clusters in first region :",len(config.Deployd.One.Region[1].Cluster)) 


}
