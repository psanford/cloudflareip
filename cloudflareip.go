package cloudflareip

import (
	"net/netip"
)

//go:generate go run update_ips.go

// IsCloudflareIP returns true if the ip address falls within one of the known
// CLOUDFLARE ip ranges.
func IsCloudflareIP(ip netip.Addr) bool {
	r := Range(ip)
	return r != nil
}

// Range returns the ip range and metadata an address falls within.
// If the IP is not a Cloudflare IP address it returns nil
func Range(ip netip.Addr) *IPRange {
	_, r, ok := cidrTbl.Lookup(ip)
	if ok {
		return &r
	}
	return nil
}

type IPRange struct {
	Prefix netip.Prefix
}
