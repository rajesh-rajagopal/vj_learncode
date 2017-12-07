
package main 

import (
  "github.com/megamsys/libgo/api"
  "github.com/megamsys/libgo/pairs"
  "encoding/json"
  "fmt"
)

type Assembly struct {
	Id           string                `json:"id" cql:"id"`
	OrgId        string                `json:"org_id" cql:"org_id"`
	AccountId    string                `json:"account_id" cql:"account_id"`
	Name         string                `json:"name" cql:"name"`
	JsonClaz     string                `json:"json_claz" cql:"json_claz"`
	Tosca        string                `json:"tosca_type" cql:"tosca_type"`
	Status       string                `json:"status" cql:"status"`
	State        string                `json:"state" cql:"state"`
	CreatedAt    string                `json:"created_at" cql:"created_at"`
	Inputs       pairs.JsonPairs       `json:"inputs" cql:"inputs"`
	Outputs      pairs.JsonPairs       `json:"outputs" cql:"outputs"`
	ComponentIds []string              `json:"components" cql:"components"`
}

type ApiAssembly struct {
	JsonClaz string     `json:"json_claz"`
	Results  []Assembly `json:"results"`
}


func main() {
   args := api.ApiArgs{
    Email: "xhoninooo@gmail.com",
    Url: "http://146.0.247.2:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
  }
  asm := TestGetAssembly(args, "ASM6238277329090116695")
  fmt.Println(asm)
  args.Org_Id = asm.OrgId
  res, err := testPost(args,"/assembly/update", asm)
  fmt.Println("Error :", err)
  fmt.Println("response :",string(res))

}

func testGet(args api.ApiArgs, path string) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Get()
}

func testPost(args api.ApiArgs,path string, item interface{}) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Post(item)
}


func TestGetAssembly(args api.ApiArgs, id string) Assembly {
  response, err := testGet(args,"/assembly/" + id)
  if err != nil {
  fmt.Println(err)
  }
	ac := &ApiAssembly{}
	err = json.Unmarshal(response, ac)
	if err != nil {
           fmt.Println(err)
	}
	a := ac.Results[0]
   fmt.Println(a)
   m := make(map[string][]string, 2)
   m["public_ipv4"] = []string{a.Outputs.Match("publicipv4")}
   a.Outputs.NukeAndSet(m)
   return a
}


