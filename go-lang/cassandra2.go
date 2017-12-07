package main

import(
    "fmt"
  //  "reflect"
    "github.com/megamsys/gocassa"
)

type Sale struct {
    Id          string 
    CustomerId  string
    SellerId    string
    Price       int
}



var (
	tablename  = "store"
	StoreIndex = []string{}
	StorePK    = []string{"Id"}
	CityKey    = StorePK[0]
)

func main() {
	c, err := gocassa.Connect([]string{"192.168.0.108"}, "", "")
	ns := c.KeySpace("testing")
        tbl := ns.MultimapMultiKeyTable("sales", StorePK, StoreIndex, Sale{})

	field := map[string]interface{}{"Id":"sale-1"}
	//field[CityKey] = "sale-1
	
	list := &[]Sale{}
        op := gocassa.Options{AllowFiltering: true}
	err = tbl.List(field, nil, 20, list).WithOptions(op).Run()
	if err != nil {
		fmt.Println(err)
	}
	if len(*list) != 1 {
		fmt.Println("******",*list,"*****")
	}
}


