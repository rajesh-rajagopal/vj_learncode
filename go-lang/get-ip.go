package main

import (
"fmt"
"net"

)

func main() {
	
	ifaces, err := net.Interfaces()
	if err != nil {
		fmt.Println(err)
	}
	for _, iface := range ifaces {
		ifaddress, err := iface.Addrs()
		if err != nil {
			fmt.Println(err)
		}
		for _, address := range ifaddress {
		   if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsMulticast() {
			if ip4 := ipnet.IP.To4(); ip4 != nil {
       			   if ip4[0] == 192 || ip4[0] == 10 || ip4[0] == 172 {
				fmt.Println("private ip:  ",ip4)
			   } else {
                                fmt.Println("public ip: ",ip4)
			   }
			    
			}
   		   }
	       }
	}
}


