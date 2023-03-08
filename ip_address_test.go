package apace

import (
	"net"
	"testing"
)

func TestRandomIPv4(t *testing.T) {
	//Generate ip address
	ipAddress := RandomIPv4()
	// Parse it
	if net.ParseIP(ipAddress) == nil {
		t.Errorf("Expected valid IP address, but got %v", ipAddress)
	}

}

func BenchmarkRandomIPv4(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		RandomIPv4()
	}
}
