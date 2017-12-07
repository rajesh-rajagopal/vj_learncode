package main

import (
	"crypto/rand"
	"fmt"
)

func main() {
	buf := make([]byte, 6)
	_, err := rand.Read(buf)
	if err != nil {
		fmt.Println("error:", err)
		return
	}
	// Set the local bit
	buf[0] |= 2
	fmt.Printf("Random MAC address: %02x:%02x:%02x:%02x:%02x:%02x\n", buf[0], buf[1], buf[2], buf[3], buf[4], buf[5])
}

