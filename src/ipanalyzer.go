package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 2 && len(os.Args) != 3 {
		log.Fatalf("Usage: %s cidr_block [ip_address]", os.Args[0])
	}
	cidr := os.Args[1]

	if len(os.Args) == 2 {
		analyzeCIDR(cidr)
	} else if len(os.Args) == 3 {
		ip := os.Args[2]
		checkIPInCIDR(cidr, ip)
	}
}

func analyzeCIDR(cidr string) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("Invalid CIDR: %v", err)
	}

	networkIP := ip.Mask(ipNet.Mask)

	broadcastIP := make(net.IP, len(networkIP))
	for i := 0; i < len(networkIP); i++ {
		broadcastIP[i] = networkIP[i] | ^ipNet.Mask[i]
	}

	mask := net.IP(ipNet.Mask)

	ones, bits := ipNet.Mask.Size()
	numHosts := 0
	if bits-ones > 1 {
		numHosts = (1 << uint(bits-ones)) - 2
	}

	fmt.Printf("Analysinf network: %s\n\n", cidr)
	fmt.Printf("Network address: %s\n", networkIP)
	fmt.Printf("Broadcast address: %s\n", broadcastIP)
	fmt.Printf("Subnet mask: %s\n", mask)
	fmt.Printf("Number of usable hosts: %d\n", numHosts)
}

func checkIPInCIDR(cidr string, ipStr string) {
	_, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Fatalf("Invalid CIDR: %v", err)
	}
	ip := net.ParseIP(ipStr)
	if ip == nil {
		log.Fatalf("Invalid IP address: %v", ipStr)
	}
	fmt.Println(ipNet.Contains(ip))
}
