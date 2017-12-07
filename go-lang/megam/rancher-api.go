
// curl -fsSL https://get.docker.com/ | sh
// sudo docker run -d --restart=unless-stopped -p 8080:8080 rancher/server

package main

import (
 "github.com/megamsys/go-rancher/v2"
  "fmt"
)

type TestSuite struct {
Client *client.RancherClient
}


func main() {
opts := client.ClientOpts{
Url: "http://192.168.0.133:8080/",
AccountId: "1a5",
Method: "post",
}
// http://185.88.29.244:8080/
//AccessKey: "B82650A88E19ACDA7FF4", // Access Key (Username)
//SecretKey: "Etj4UWdn9YYzQzHRLK1RnhFV6s2Nz9UMQVsKX999",  // Secret Key (Password)


clt, err :=client.NewRancherClient(&opts)
if err != nil {
  fmt.Println("New Rancher Error :", err)
  return
}

t := TestSuite{Client: clt}

// t.GetContainer("1i27")
// t.CreateCon()
// t.DeleteContainers("1i13","1i14")
// t.GetHost("1h1")
// t.LogsContainers("1i20")
// t.StopContainer("1i19")
// t.StartContainer("1i19")
}

func (t *TestSuite) CreateCon() {

cont := client.Container{
AccountId: "1a5",
Count: 1,
Name: "myNewCont3",
ImageUuid: "docker:tutum/hello-world",
StartOnCreate: true,
}

contan, err:= t.Client.Container.Create(&cont)
if err != nil {
  fmt.Println("Create container Error :", err)
} else {
  fmt.Printf("\nSuccess  :%T    ---- \n    %#v",contan)
}

}

func (t *TestSuite) GetContainer(id string) {
  cont, err := t.Client.Container.ById(id)
  if err != nil {
    fmt.Println("GET Error :",err)
  }
  fmt.Println("RESPONSE :",cont)
}


func (t *TestSuite) DeleteContainers(contIds ...string) {
  for _, v := range contIds {
    cont, err := t.Client.Container.ById(v)
    if err != nil {
      fmt.Println("Get Rancher Container Error :", err)
      continue
    }
    err = t.Client.Container.Delete(cont)
    if err != nil {
      fmt.Println("Delete Rancher Container Error :", err)
    } else {
      fmt.Println("Container deleted : ",v)
    }

  }
}


func (t *TestSuite) LogsContainers(v string)  {
    cont, err := t.Client.Container.ById(v)
    contlogs, err := t.Client.ContainerLogs.ById(v)
    if err != nil {
      fmt.Println("Get Rancher Container Error :", err)
      return
    }
   fmt.Println("   **********************cont log :",contlogs)
   contlogs.Lines = 200

    ht, err := t.Client.Container.ActionLogs(cont, contlogs)
    if err != nil {
      fmt.Println("action log Error :", err)
    }
    fmt.Println("Container logs :", ht)
}

func (t *TestSuite) GetHost(v string) {
    cont, err := t.Client.Host.ById(v)
    if err != nil {
      fmt.Println("Get Rancher Host Error :", err)
    }
    fmt.Printf("\nHost :%#v", cont)
}

func (t *TestSuite) StopContainer(v string) {
  cont, err := t.Client.Container.ById(v)
    insStop, err := t.Client.InstanceStop.ById(v)
    if err != nil {
      fmt.Println("Get Rancher Container Error :", err)
      return
    }
   fmt.Println("   **********************cont log :",insStop)

    instance, err := t.Client.Container.ActionStop(cont, insStop)
    if err != nil {
      fmt.Println("action log Error :", err)
    }
    fmt.Println("Container logs :", instance)
}

func (t *TestSuite) StartContainer(v string) {
cont, err := t.Client.Container.ById(v)
 if err != nil {
      fmt.Println("Get Rancher Container Error :", err)
      return
    }

    instance, err := t.Client.Container.ActionStart(cont)
    if err != nil {
      fmt.Println("action log Error :", err)
    }
 fmt.Println("Container logs :", instance)
}


func (t *TestSuite) GetLogs(v string) {
  
}
