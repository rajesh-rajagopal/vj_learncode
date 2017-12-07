package main

import (
    "fmt"
    "net"
)

func localAddresses() {
    ifaces, err := net.Interfaces()
    if err != nil {
        fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
        return
    }
    fmt.Println("interfaces : ",ifaces)
    for _, i := range ifaces {
        addrs, err := i.Addrs()
        if err != nil {
            fmt.Print(fmt.Errorf("localAddresses: %+v\n", err.Error()))
            continue
        }
        fmt.Println("address : ",addrs)
        for _, a := range addrs {
            fmt.Println("a   :",a)
            switch v := a.(type) {
            case *net.IPAddr:
                fmt.Printf("i.Name - %v : %s (%s)\n", i.Name, v, v.IP.DefaultMask())
            }

        }
    }
}

func main() {
    localAddresses()
}
