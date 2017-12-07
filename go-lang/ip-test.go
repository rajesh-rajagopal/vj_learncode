package main 

import (
 "github.com/tatsushid/go-fastping"
"fmt"
"os"
"net"
"time"
)

func main() {
  ping()
}

func ping() {
p := fastping.NewPinger()
ra, err := net.ResolveIPAddr("ip4", "192.168.2.47")
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
p.AddIPAddr(ra)
p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
	fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
}
p.OnIdle = func() {
	fmt.Println("finish")
}

fmt.Printf("\n    ------  %#v",p)
err = p.Run()
if err != nil {
	fmt.Println(err)
}
}
