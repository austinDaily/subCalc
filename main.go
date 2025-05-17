package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		printUsage()
		os.Exit(1)
	}

	cidr := os.Args[1]
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		fmt.Fprint(os.Stderr, "Error: invalid CIDR '%s': %v\n", cidr, err)
		os.Exit(1)
	}

	printSubnetInfo(ip, ipNet)
}

func printUsage() {
	fmt.Println("Subnet Calculator CLI")
	fmt.Println("--------------------")
	fmt.Println("Usage: subnetcalc <CIDR>")
	fmt.Println("Example: subnetcalc 192.168.1.0/24")
}

func printSubnetInfo(ip net.IP, ipNet *net.IPNet) {
	mask := ipNet.Mask
	ones, bits := mask.Size()
	broadcast := calculateBroadcast(ip, mask)
	numHosts := calculateUsableHosts(ones, bits)

	fmt.Println("Input CIDR:        ", ipNet.String())
	fmt.Println("Network Address:   ", ipNet.IP.String())
	fmt.Println("Subnet Mask:       ", net.IP(mask).String())
	fmt.Println("Prefix Length:     ", fmt.Sprintf("/%d", ones))
	fmt.Println("Broadcast Address: ", broadcast.String())
	fmt.Println("Usable Hosts:      ", numHosts)
}

func calculateBroadcast(ip net.IP, mask net.IPMask) net.IP {
	ip = ip.To4() // ensure it's in IPv4 form
	if ip == nil {
		return nil // fallback for IPv6 or invalid IP
	}

	broadcast := make(net.IP, len(ip))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^mask[i]
	}
	return broadcast
}

func calculateUsableHosts(ones, bits int) int {
	hostBits := bits - ones
	if hostBits <= 1 {
		return 0 // /31 and /32 have 0 usable hosts
	}
	return (1 << hostBits) - 2
}
