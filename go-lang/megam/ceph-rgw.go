package main

import (
	"fmt"
        "time"
	"github.com/megamsys/go-radosgw/api"
//	"github.com/QuentinPerez/go-radosgw/pkg/api"
)


type RadosGW struct {
  Api *radosAPI.API
}

type Buckets struct {
  UserId string `json:"uid"`
  Bucket *Bucket `json:"bucket"`
  TotalSize int `json:"total_size"`
}

type Bucket struct {
  Name string `json:"enabled"`
  Owner string `json:"owner"`
  Quota *BucketQuota `json:"bucket_quota"`
  Usage *Usage `json:"usage"`
}

type BucketQuota struct {
  Enabled    bool `json:"enabled"`
  MaxObjects int  `json:"max_objects"`
  MaxSizeKb  int  `json:"max_size_kb"`
}

type Usage struct {
    NumObjects   int `json:"num_objects"`
    SizeKb       int `json:"size_kb"`
    SizeKbActual int `json:"size_kb_actual"`
}

func main() {
           
       // "user": "rajesh", "U15C54BTWCL8CL77E0QZ", "6CyKhGyPWT7MHSSUC2z4E4kMb6bAPYSv96apEaYC" has user read persion

	api := radosAPI.New("http://192.168.0.115:7480", "V3I3ID7RC0A3H3WB3GHE", "G2YRbwnUKVq0f4oOuaDsltiDR5Bh8AxCCMt2O4M3")
	//fmt.Println(api)
	// get users named JohnDoe
         //users := []string{}
         //getUsers( api ) // "vijaykanth@megam.io", "raj@raj.com", "thilak@b.com"

         fmt.Println(api)	
         getBucket("", "thilak@b.com",api)

}

func getUsers(api *radosAPI.API) {
	user, err := api.GetUser("")
	if err != nil {
		fmt.Println("get user error :",err)
	} else {
           fmt.Println("user : ", user)
        }

}


func getUser(api *radosAPI.API, users ...string) {
	
for _, u := range users {
	user, err := api.GetUser(u)
	if err != nil {
		fmt.Println("get user error :",err)
	} else {
           fmt.Println("user : ", user)
        }
        start := time.Now().Add(-24*time.Hour)
	end := time.Now()
	usage, err := api.GetUsage(radosAPI.UsageConfig{
		UID:         u,
                Start: &start,
		End: &end,
		ShowEntries: true,
    		ShowSummary: true,
	})

	if err != nil {
		fmt.Println("get usage error :",err)
	}

	fmt.Println(usage)
 }
}

func getBucket(name,user string,api *radosAPI.API) {
   	bkt := radosAPI.BucketConfig{Bucket: "", UID: user, Stats: true}
   	bucket, err := api.GetBuckets(bkt)
	if err != nil {
		fmt.Println("error bucket ",err)
	}
        for _, i := range bucket {
		fmt.Println(i.Stats.Usage)
		fmt.Println("Bucket : ",i.Stats.Bucket)
		fmt.Println("Size ",float64(float64(i.Stats.Usage.RgwMain.SizeKbActual)/1024.0)," MB")
		fmt.Println("No. of items :", i.Stats.Usage.RgwMain.NumObjects)
        }
	
} 

