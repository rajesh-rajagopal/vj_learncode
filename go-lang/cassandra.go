package main

import(
    "fmt"
    "reflect"
    "github.com/megamsys/gocassa"
)

type Sale struct {
    Id          string 
    CustomerId  string
    SellerId    string
    Price       int
}



type Store struct {
	City    string
	Manager string
	Id      string
	Address string
}

var (
	tablename  = "store"
	StoreIndex = []string{"Manager","Id"}
	StorePK    = []string{"City"}
	CityKey    = StorePK[0]
	ManagerKey = StoreIndex[0]
	IdKey      = StoreIndex[1]
)

func main() {
	c, err := gocassa.Connect([]string{"192.168.0.108"}, "", "")
	ns := c.KeySpace("testing")
        tbl := ns.MultimapMultiKeyTable("stores", StorePK, StoreIndex, Store{})
	london := Store{
		City:    "London",
		Manager: "Joe",
		Id:      "12412-afa-16958",
		Address: "Somerset House",
	}
	err = tbl.Set(london).Run()
	if err != nil {
		fmt.Println(err)
	}
	field := make(map[string]interface{})
	id := make(map[string]interface{})

	field[CityKey] = "London"
	id[ManagerKey] = "Joe"
	id[IdKey] = "12412-afa-16956"

	res := &Store{}
	err = tbl.Read(field, id, res).Run()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res)
	if !reflect.DeepEqual(*res, london) {
		fmt.Println(*res, london)
	}
	err = tbl.Delete(field, id).Run()
	if err != nil {
		fmt.Println(err)
	}
	err = tbl.Read(field, id, res).Run()
	if err == nil {
		fmt.Println(res)
	}
	
	list := &[]Store{}
        op := gocassa.Options{AllowFiltering: true}
	err = tbl.List(field, nil, 20, list).WithOptions(op).Run()
	if err != nil {
		fmt.Println(err)
	}
	if len(*list) != 1 {
		fmt.Println("******",*list,"*****")
	}
}


