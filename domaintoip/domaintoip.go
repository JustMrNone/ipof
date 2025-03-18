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

	ipv4, ipv6, err := DomainToIp(domain)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if ipv4 != "" {
		fmt.Println("IPv4 Address:", ipv4)
	}
	if ipv6 != "" {
		fmt.Println("IPv6 Address:", ipv6)
	}
}

func DomainToIp(domain string) (string, string, error) {
	ips, err := net.LookupIP(domain)

	if err != nil {
		return "", "", fmt.Errorf("could not get the IPs: %v", err)
	}

	var ipv4, ipv6 string

	for _, ip := range ips {
		if ipv4 == "" && ip.To4() != nil {
			ipv4 = ip.String()
		} else if ipv6 == "" && ip.To4() == nil {
			ipv6 = ip.String()
		}
	}

	if ipv4 == "" && ipv6 == "" {
		return "", "", fmt.Errorf("no IPs found for domain: %s", domain)
	}

	return ipv4, ipv6, nil
}
