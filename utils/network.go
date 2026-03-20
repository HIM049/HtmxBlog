package utils

import (
	"net"
	"net/http"
	"strings"
)

// GetRealIP returns the real client IP, taking proxy headers into account.
func GetRealIP(r *http.Request) string {
	// Check X-Forwarded-For
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	if xForwardedFor != "" {
		ips := strings.Split(xForwardedFor, ",")
		ip := strings.TrimSpace(ips[0])
		if host, _, err := net.SplitHostPort(ip); err == nil {
			return host
		}
		return ip
	}

	// Check X-Real-IP
	xRealIP := r.Header.Get("X-Real-IP")
	if xRealIP != "" {
		if host, _, err := net.SplitHostPort(xRealIP); err == nil {
			return host
		}
		return xRealIP
	}

	// Fallback to RemoteAddr
	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return r.RemoteAddr
	}
	return host
}
