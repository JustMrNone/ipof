package domaintoip

import (
	"fmt"
	"net"
	"os"
)

func Ipof() {
	var domain string

	// Check if a keyword argument is provided
	if len(os.Args) > 1 {
		domain = os.Args[1]
	} else {
		fmt.Print("Enter domain name: ")
		fmt.Scan(&domain)
	}

	ipv4, err := DomainToIp(domain)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(ipv4) > 0 {
		fmt.Println("IPv4 Addresses:")
		for _, ip := range ipv4 {
			fmt.Println(" -", ip)
		}
	}
}

func DomainToIp(domain string) ([]net.IP, error) {
	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, fmt.Errorf("could not get the IPs: %v", err)
	}
	return ips, nil
}
