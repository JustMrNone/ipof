package domaintoip

import (
	"net"
	"testing"
)

func TestDomainToIp(t *testing.T) {
	tests := []struct {
		domain    string
		expectIPs []net.IP
		expectErr bool
	}{
		{"google.com", nil, false}, // Allow any IPs for google.com
		{"invalid.domain", nil, true},
		{"localhost", []net.IP{net.ParseIP("::1"), net.ParseIP("127.0.0.1")}, false},
	}

	for _, test := range tests {
		ips, err := DomainToIp(test.domain)
		if (err != nil) != test.expectErr {
			t.Errorf("DomainToIp(%q) error = %v, expectErr %v", test.domain, err, test.expectErr)
		}
		if !test.expectErr && test.expectIPs != nil && !containsAllIPs(ips, test.expectIPs) {
			t.Errorf("DomainToIp(%q) ips = %v, expectIPs %v", test.domain, ips, test.expectIPs)
		}
	}
}

// containsAllIPs checks if all expected IPs exist in the resolved IPs
func containsAllIPs(resolvedIPs, expectedIPs []net.IP) bool {
	ipSet := make(map[string]bool)
	for _, ip := range resolvedIPs {
		ipSet[ip.String()] = true
	}
	for _, ip := range expectedIPs {
		if !ipSet[ip.String()] {
			return false
		}
	}
	return true
}
