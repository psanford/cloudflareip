package cloudflareip

import (
	"net/netip"
	"testing"
)

func TestIsCloudflareIP(t *testing.T) {
	cfIPs := []netip.Addr{
		netip.MustParseAddr("103.21.244.21"),
		netip.MustParseAddr("2405:8100::1"),
	}

	for _, addr := range cfIPs {
		if !IsCloudflareIP(addr) {
			t.Errorf("Expected %s to match cloudflare ip but did not", addr)
		}
	}

	nonCloudflareIPs := []netip.Addr{
		netip.MustParseAddr("127.0.0.12"),
		netip.MustParseAddr("10.48.20.96"),
		netip.MustParseAddr("8.8.8.8"),
		netip.MustParseAddr("2a05:d03a:8000::1"),
	}
	for _, addr := range nonCloudflareIPs {
		if IsCloudflareIP(addr) {
			t.Errorf("%s is not an cloudflare ip address, but it matched", addr)
		}
	}
}
