package httpx

import (
	"net/http"
	"testing"
)

func TestClientIP(t *testing.T) {
	cases := []struct {
		name            string
		xff, remoteAddr string
		exp             string
	}{
		{"empty", "", "", ""},
		{"xff", "1.2.3.4", "", "1.2.3.4"},
		{"xff-2", "1.2.3.4,5.6.7.8", "", "1.2.3.4"},
		{"xff-3", "  1.2.3.4  ,  5.6.7.8", "", "1.2.3.4"},
		{"remoteAddr", "", "1.2.3.4:5678", "1.2.3.4"},
		{"remoteAddr-2", "", "1.2.3.4", "1.2.3.4"},
		{"bad-remoteAddr", "", "invalid", "invalid"},
	}

	for _, c := range cases {
		r := &http.Request{
			Header: http.Header{
				"X-Forwarded-For": []string{c.xff},
			},
			RemoteAddr: c.remoteAddr,
		}

		if got := ClientIP(r); got != c.exp {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.exp, got)
		}
	}
}
