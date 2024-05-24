package passport

import (
	"net"
	"net/mail"
	"net/url"
)

// contains 是否包含
func contains[T comparable](arr []T, target T) bool {
	for i := range arr {
		if arr[i] == target {
			return true
		}
	}
	return false
}

// isZero 零值判断
func isZero[T comparable](v T) bool {
	var zero T
	return v == zero
}

func isIPv4(v string) bool {
	ip := net.ParseIP(v)
	if ip == nil {
		return false
	}
	return ip.To4() != nil
}

func isIPv6(v string) bool {
	ip := net.ParseIP(v)
	if ip == nil {
		return false
	}
	return ip.To4() == nil && len(ip) == net.IPv6len
}

func isURL(s string) bool {
	r, err := url.Parse(s)
	if err != nil {
		return false
	}
	return r.Scheme != "" && r.Host != ""
}

func isEmail(s string) bool {
	addr, err := mail.ParseAddress(s)
	if err != nil {
		return false
	}
	return addr.Name == ""
}
