package main

import (
	"github.com/megamsys/libgo/api"
	"github.com/megamsys/libgo/pairs"
	"encoding/json"
	"fmt"
	"strconv"
)

type Quota struct {
	Id          string          `json:"id" cql:"id"`
	AccountId   string          `json:"account_id" cql:"account_id"`
	Name        string          `json:"name" cql:"name"`
	JsonClaz    string          `json:"json_claz" cql:"json_claz"`
	Allowed     pairs.JsonPairs `json:"allowed" cql:"allowed"`
	AllocatedTo string          `json:"allocated_to" cql:"allocated_to"`
	QuotaType   string          `json:"quota_type" cql:"quota_type",omitempty`
	Status      string          `json:"status" cql:"status"`
       // CreateAt    string          `json:"created_at"`
	Inputs      pairs.JsonPairs `json:"inputs" cql:"inputs"`
}

type ApiQuota struct {
	JsonClaz string  `json:"json_claz"`
	Results  []Quota `json:"results"`
}

// IST to UTC(table) - 5:30 so 
// we have to use with get from table +5:30
// fddg@ss.co | 2017-02-03 12:04:06.000000+0000 | QUO8671194369186440757

// update quotas set quota_type='vm', status='active' where account_id='fddg@ss.co' and created_at='2017-02-03T17:34:06.0000' and id='QUO8671194369186440757';
//update quotas set quota_type='vm', status='active' where account_id='fddg@ss.co' and created_at='2017-02-03T17:34:18.0000' and id='QUO5830545150944309312';


func main() {
   args := api.ApiArgs{
    Email: "mvijaykanth@megam.io",
    Url: "http://192.168.0.6:9000/v2",
    Api_Key: "",
    Master_Key: "3b8eb672aa7c8db82e5d34a0744740b20ed59e1f6814cfb63364040b0994ee3f",
    Password: "",
  }
  q := new(Quota)
  q.AccountId = args.Email
  q.Id = "QUO8671194369186000000"
  quota,err := q.get(args)
  if err != nil {
   fmt.Println("Get Error :", err)
    return 
  }
  fmt.Println(quota) 
  err = quota.update(args)
  fmt.Println("Update Error :", err)


}
 
// QUO8671194369186440757

func testGet(args api.ApiArgs, path string) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Get()
}

func testPost(args api.ApiArgs,path string, item interface{}) ([]byte, error) {
  cl := api.NewClient(args, path)
  return cl.Post(item)
}

func (q *Quota) update(args api.ApiArgs) error {
     	count, _ := strconv.Atoi(q.Allowed.Match("snapshots"))
	mm := make(map[string][]string, 1)
	mm["snapshots"] = []string{strconv.Itoa(count-1)}
	q.Allowed.NukeAndSet(mm) 		
       res, err := testPost(args,"/quotas/update", q)
	if err != nil {
		return err
	}
      fmt.Println("Result :",string(res))
	return nil
}

func (q *Quota) get(args api.ApiArgs) (*Quota, error) {
	cl := api.NewClient(args, "/quotas/"+ q.Id)
	response, err := cl.Get()
	if err != nil {
		return nil, err
	}
	ac := &ApiQuota{}
        fmt.Println(string(response))
	//log.Debugf("Response %s :  (%s)",cmd.Colorfy("[Body]", "green", "", "bold"),string(response))
	err = json.Unmarshal(response, ac)
	if err != nil {
		return nil, err
	}

	return &ac.Results[0], nil
}


