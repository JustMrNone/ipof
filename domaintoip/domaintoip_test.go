package domaintoip

import (
	"testing"
)

func TestDomainToIp(t *testing.T) {
	tests := []struct {
		domain     string
		expectIPv4 string
		expectIPv6 string
		expectErr  bool
	}{
		{"google.com", "", "", false}, // Replace with actual expected IPs
		{"invalid.domain", "", "", true},
		{"localhost", "127.0.0.1", "::1", false},
	}

	for _, test := range tests {
		ipv4, ipv6, err := DomainToIp(test.domain)
		if (err != nil) != test.expectErr {
			t.Errorf("DomainToIp(%q) error = %v, expectErr %v", test.domain, err, test.expectErr)
		}
		if ipv4 != test.expectIPv4 && test.expectIPv4 != "" {
			t.Errorf("DomainToIp(%q) ipv4 = %v, expectIPv4 %v", test.domain, ipv4, test.expectIPv4)
		}
		if ipv6 != test.expectIPv6 && test.expectIPv6 != "" {
			t.Errorf("DomainToIp(%q) ipv6 = %v, expectIPv6 %v", test.domain, ipv6, test.expectIPv6)
		}
	}
}
