package main

import (
 "fmt"
 "net/http"
)

func main() {

client := &http.Client{}
// headers := &http.Headers{}
//headers.Set(key, value)

/* Authenticate */
req, err := http.NewRequest("GET", "http://103.56.92.58:5353/admincp/vmlist.php", nil)
req.SetBasicAuth("vpsadmin","vpsadmin")
resp, err := client.Do(req)
if err != nil {
    fmt.Printf("Error : %s", err)
}



fmt.Println("\n\nresponse")
fmt.Printf("%#v  ",resp)

}

