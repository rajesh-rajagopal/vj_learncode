
package main 

import (
  "github.com/megamsys/libgo/api"
  "time"
  "encoding/json"
  "fmt"
)


type ApiBalances struct {
	JsonClaz string     `json:"json_claz" cql:"json_claz"`
	Results  []Balances `json:"results" json:"results"`
}
type Balances struct {
	Id        string `json:"id" cql:"id"`
	AccountId string `json:"account_id" cql:"account_id"`
	Credit    string `json:"credit" cql:"credit"`
	JsonClaz  string `json:"json_claz" cql:"json_claz"`
}

type updateBalances struct {
	Id        string    `json:"id" cql:"id"`
	AccountId string    `json:"account_id" cql:"account_id"`
	Credit    string    `json:"credit" cql:"credit"`
        BillerCredit string `json:"biller_credit"`
	JsonClaz  string    `json:"json_claz" cql:"json_claz"`
	CreatedAt time.Time `json:"created_at" cql:"created_at"`
	UpdatedAt time.Time `json:"updated_at" cql:"updated_at"`
}

func main() {
   args := api.ApiArgs{
    Email: "mvijaykanth@gmail.com",
    Url: "http://146.0.247.2:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
    Org_Id: "",
  }
 
  bls := getBalances(args, []string{"djlulimix@gmail.com"})
  for _, bl := range bls {
   fmt.Println(bl)
   // _ = bl.update(args)
  }
  
}

func getBalances(args api.ApiArgs, users []string) []*Balances {
    bls := make([]*Balances,0)
    for _, email := range users {
        args.Email = email
	cl := api.NewClient(args, "/balances/" + email)
	response, err := cl.Get()
	if err != nil {
		fmt.Println( email, "*******get******",err)
	}

	ac := &ApiBalances{}
	err = json.Unmarshal(response, ac)
	if err != nil {
	   fmt.Println( email, "******parse*******",err)
	}
	bls = append(bls,&ac.Results[0])
   }
       fmt.Println("balances updated **************",len(bls))
  return bls

}

func (b *Balances) update(args api.ApiArgs) error {
         args.Email = b.AccountId
	 cl := api.NewClient(args, "/balances/bill")
	 _, err := cl.Post(b.toUpdate())
	 if err != nil {
	 	fmt.Println( args.Email, "*******post******",err)
	 }
    return nil
}


func (b *Balances) toUpdate() *updateBalances {
	return &updateBalances{
		Id:        b.Id,
		AccountId: b.AccountId,
		Credit:    b.Credit,
		JsonClaz:  b.JsonClaz,
                BillerCredit: 	"0",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

