package main

import (
  "fmt"
 "reflect"
 "strings"
)

type Credentials struct {
	AccountId string
	Api_Key string
	Master_Key string
	Password string
}

type Test struct{
         AccountId string
        Api_Key string
        Master_Key string
        Password string
  Name string
  Age  int 
  Address string 
}

func main() {
  t := &Test{Name: "vijay", Age: 25, Address: "31B, Idappadi,Salem, TN"}
  fmt.Println("Map :",t.get())
}

func (c *Test) get() map[string]string {
	keys := make(map[string]string)
	s := reflect.ValueOf(c).Elem()
	typ := s.Type()
	if s.Kind() == reflect.Struct {
		for i := 0; i < s.NumField(); i++ {
			key := s.Field(i)
			value := s.FieldByName(typ.Field(i).Name)

			switch key.Interface().(type) {
			case string:
				if value.String() != "" {
					keys[strings.ToLower(typ.Field(i).Name)] = value.String()
				}
			}
		}
	}
	return keys
}
