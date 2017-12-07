package main

import (
	"fmt"
	"os"
       	"io/ioutil"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
)

type MainConf struct {
 Conf Config `json:""  toml:"master"`
}

type Config struct {
	One   *Oned  `json:"oned" toml:"oned"`
	Repos *Repos `json:"repos" toml:"repos"`
}

type Oned struct {
	OneEnpoint   string `json:"one_endpoint" toml:"one_endpoint"`
	OneUser      string `json:"one_user" toml:"one_user"`
	OnePassword  string `json:"one_password" toml:"one_password"`
	OneImage     string `json:"api" toml:"api"`
	OneTemplate  string `json:"one_template" toml:"one_template"`
	OneClusterID string `json:"cluster_id" toml:"cluster_id"`
	OneVnet      string `json:"one_vnet" toml:"one_vnet"`
	VCPU         string `json:"vcpu" toml:"vcpu"`
	CPU          string `json:"cpu" toml:"cpu"`
	RAM          string `json:"memory" toml:"memory"`
	HDD          string `json:"hdd" toml:"hdd"`
}

type OneConf struct {
	ImageUrls map[string]string `json:"image_urls"`
	OneInitScript string `json:"init_scripts"`
	XMLTemplate string  `json:"XMLTemplate"`
}

type Repos struct {
	Distros []Distro `json:"distro" toml:"distro"`
}

type Distro struct {
        Name         string `json:"name" toml:"name"`
	OneRepos     string `json:"one_repos" toml:"one_repos"`
	VerticeRepos string `json:"vertice_repos" toml:"vertice_repos"`
	CephRepos    string `json:"ceph_repos" toml:"ceph_repos"`
	CQLRepos     string `json:"cassandra_repos" toml:"cassandra_repos"`
}

func main() {
       path := ""
       config := &MainConf{}
	if path == "" {
		path = "/home/alrin/code/megam/directory/obc/api.conf"
	}
	log.Warnf("Using configuration at: %s", path)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	if err := toml.Unmarshal(buf, &config); err != nil {
		panic(err)
	}
	fmt.Printf("conf     ********      :%#v",config)
	
}

