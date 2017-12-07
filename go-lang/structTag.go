package main

import (
"fmt"
"reflect"
)

type example struct{
 Name string `json:"user_name"`
 Address string `json:"address"`
}

func main() {
b := example{Name: "vijay"}
b.PrintFields()

}
func (b example) PrintFields() {
    val := reflect.ValueOf(b)
    s := reflect.ValueOf(b).Elem().Type()
    for i := 0; i < val.Type().NumField(); i++ {
        k := val.Type().Field(i).Name
        s.FieldByName(k)
        fmt.Println(val.Type().Field(i).Tag.Get("json"), " : ",k)
        
    }
}

