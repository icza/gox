package httpx

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP returns the client's IP address for the given request.
// It first check the X-Forwarded-For header, and if present, returns its first
// element.
//
// Else Request.RemoteAddr is parsed and used.
func ClientIP(r *http.Request) string {
	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if parts := strings.Split(xff, ","); len(parts) > 0 {
			// Intermediate nodes append, so first is the original client
			return strings.TrimSpace(parts[0])
		}
	}

	addr, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return addr
	}

	// Last resort:
	return r.RemoteAddr
}
