package main

import (
 "github.com/influxdata/toml"
 "fmt"
 "io/ioutil"
 "os"
//    "time"
)
//  fruit.physical.color.taste.variety
type Config struct {
Fruit  fruit
}

type fruit struct { deplyde
Name string
Physical physical
}

type physical struct { one 
Color []color 
Shape string
}

type color struct {  region
Colour string
Taste []taste
}

type taste struct { cluster
Flavor string 
Variety []variety
} 

type variety struct {  vnet
Name string
}


func main() {


f, err := os.Open("example.conf")
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

fmt.Println("fruit[apple]-physical-shape   :",config.Fruit.Physical.Shape)
fmt.Println("fruit[apple]-physical-color-taste   :",config.Fruit.Physical.Color)
fmt.Println("fruit[apple]-physical-color[red]-taste[sweet]-flavous   :",config.Fruit.Physical.Color[0].Taste[0].Variety)
fmt.Println("Length of color :",len(config.Fruit.Physical.Color))

}
