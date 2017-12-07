package main

import (
	"fmt"
)

type Location struct {
 x, y int
 location string
}

type dist struct {
 name string 
 locations []Location
}

type state struct {
 name string
 districts []dist
}

func PrintAll(vals map[int]interface{}) {
	for k, val := range vals {
		fmt.Println(k," - ",val)
	}
}

func main() {
	names := state{name: "tamilnadu", districts: []dist{dist{name:"salem",locations: []Location{Location{location: "attur",x: 8,y: 5},Location{location: "idappadi",x: 6,y: 5},}},dist{name:"chennai",locations: []Location{Location{location: "tnagar",x: 8,y: 5},Location{location: "velachery",x: 6,y: 5},}},},} 
	var a []string
	vals := make(map[int]interface{})
	for i, v := range names.districts {
		vals[i] = v
		a = append(a ,v.name)
	}
	var test interface{}
	test = vals
	fmt.Println("***********",test)
	PrintAll(vals)
	
	if w, ok := test.(map[int]interface{}); ok {
        fmt.Println("given value is i =", w[0])
        }
	
	fmt.Println("districts :",a)
	vals2 := make(map[int]interface{})
	for i, v := range a {
	  vals2[i] = v
	}
	PrintAll(vals2)
	
}

