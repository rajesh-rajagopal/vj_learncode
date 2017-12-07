package main

import (
 "fmt"
 "net"
 "strings"
)

const (
	PUBLICIPV4   = "public_ipv4"
	PRIVATEIPV4  = "private_ipv4"
	PUBLICIPV6   = "public_ipv6"
	PRIVATEIPV6  = "private_ipv6"
)

func main() {
	ips := mergeSameIPtype(findIps())
	fmt.Println("  find and setips of machine ",ips)
	
}

func mergeSameIPtype(mm map[string][]string) map[string][]string {
	for IPtype, ips := range mm {
		var sameIp string
		for _, ip := range ips {
			sameIp = sameIp + ip + ", "
		}
		if sameIp != "" {
			mm[IPtype] = []string{strings.TrimRight(sameIp, ", ")}
		}
	}
	return mm
}

// FindIps returns the non loopback local IP4 (can be public or private)
// if an iface contains a string "pub", then we consider it a public interface
func findIps() map[string][]string {
	var ips = make(map[string][]string)
        var ipv4s = make(map[string]string, 0)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ips
	}
	for _, iface := range ifaces {
		ifaddress, err := iface.Addrs()
		if err != nil {
			return ips
		}
		for _, address := range ifaddress {
			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() && !ipnet.IP.IsMulticast() {
				if ip4 := ipnet.IP.To4(); ip4 != nil {
					if ip4[0] == 192 || ip4[0] == 10 || ip4[0] == 172 {
						ipv4s[ipnet.IP.String()] = utils.PRIVATEIPV4
					} else {
						ipv4s[ipnet.IP.String()] = utils.PUBLICIPV4 
					}
				}
			}
		}
	}
        for k,v := range ipv4s {
           ips[v] = append(ips[v],k)
        }
	return ips
}


