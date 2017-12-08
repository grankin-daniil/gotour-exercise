package main

import (
	"fmt"
	"strconv"
	"strings"
)

type IPAddr [4]byte

func (addr IPAddr) String() string {
	chunks := make([]string, len(addr))
	for _, chunk := range addr {
		chunks = append(chunks, strconv.Itoa(int(chunk)))
	}
	return strings.Join(chunks, ".")
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
