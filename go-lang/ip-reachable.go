package main
 
import (
   "fmt"
   "net"
   "time"
)
 
func main() {
    host := "192.168.1.47"
   if isReachable(host) {
    fmt.Println(host, "reable")
   } else {
    fmt.Println(host, "Not reable")
  }
}
 
func isReachable(hostname string) bool {
   hostName := hostname
   portNum := "80"
   seconds := 5
   timeOut := time.Duration(seconds) * time.Second
 
   _, err := net.DialTimeout("tcp", hostName+":"+portNum, timeOut)

   if err != nil {
      fmt.Println(err)
      return false 
   }
return true
}
