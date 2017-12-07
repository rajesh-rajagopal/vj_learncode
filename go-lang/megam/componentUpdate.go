
package main 

import (
  "github.com/megamsys/libgo/api"
  "github.com/megamsys/libgo/pairs"
  "encoding/json"
  "fmt"
)

type ApiComponent struct {
	JsonClaz string    `json:"json_claz"`
	Results  []Component `json:"results"`
}

type Artifacts struct {
	Type         string          `json:"artifact_type" cql:"type"`
	Content      string          `json:"content" cql:"content"`
	Requirements pairs.JsonPairs `json:"requirements" cql:"requirements"`
}

type Repo struct {
	Rtype    string `json:"rtype" cql:"rtype"`
	Branch   string `json:"branch" cql:"branch"`
	Source   string `json:"source" cql:"source"`
	Oneclick string `json:"oneclick" cql:"oneclick"`
	Rurl     string `json:"url" cql:"url"`
}

type Component struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Tosca             string          `json:"tosca_type"`
	Artifacts         *Artifacts      `json:"artifacts"`
	Inputs            pairs.JsonPairs `json:"inputs"`
	Outputs           pairs.JsonPairs `json:"outputs"`
	Envs              pairs.JsonPairs `json:"envs"`
	Repo              Repo            `json:"repo"`
	RelatedComponents []string        `json:"related_components"`
	Status            string          `json:"status"`
	State             string          `json:"state"`
	CreatedAt         string	  `json:"created_at"`
}

func main() {
   args := api.ApiArgs{
    Email: "mvijaykanth@gmail.com",
    Url: "http://192.168.1.4:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
    Org_Id: "ORG8296783039362231677",
  }
  asm := TestGetComponent(args)
  // fmt.Println(asm)
 //  m := make(map[string][]string, 2)
//   m["provider"] = []string{"rancher"}
 //  asm.Inputs.NukeAndSet(m)
  asm.Tosca = "vertice.container.tutum/hello-world"
  res, err := testPost(args,"/components/update", asm)
  fmt.Println("Error :", err)
  fmt.Println("response :",string(res))

}

//'ORG8296783039362231677', 'ASM5976862151476671458', 'mvijaykanth@gmail.com',


func testGet(args api.ApiArgs, path string) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Get()
}

func testPost(args api.ApiArgs,path string, item interface{}) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Post(item)
}


func TestGetComponent(args api.ApiArgs) Component {
  response, err := testGet(args,"/components/COM4961547459155819617")
  if err != nil {
  fmt.Println(err)
  }
  fmt.Println("Response : ", string(response))
	ac := &ApiComponent{}
	err = json.Unmarshal(response, ac)
	if err != nil {
           fmt.Println(err)
	}
	a := ac.Results[0]
   return a
}
