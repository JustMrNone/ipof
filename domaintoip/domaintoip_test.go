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
		{"google.com", nil, false}, // Replace with actual expected IPs if needed
		{"invalid.domain", nil, true},
		{"localhost", []net.IP{net.ParseIP("127.0.0.1"), net.ParseIP("::1")}, false},
	}

	for _, test := range tests {
		ips, err := DomainToIp(test.domain)
		if (err != nil) != test.expectErr {
			t.Errorf("DomainToIp(%q) error = %v, expectErr %v", test.domain, err, test.expectErr)
		}
		if !test.expectErr && !compareIPSlices(ips, test.expectIPs) {
			t.Errorf("DomainToIp(%q) ips = %v, expectIPs %v", test.domain, ips, test.expectIPs)
		}
	}
}

func compareIPSlices(a, b []net.IP) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !a[i].Equal(b[i]) {
			return false
		}
	}
	return true
}
