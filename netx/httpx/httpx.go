package httpx

import (
	"net"
	"net/http"
	"strings"
)

// ClientIP returns the client's IP address for the given request.
// It first checks the X-Forwarded-For header, and if present, returns its first
// element.
//
// Else Request.RemoteAddr is parsed and used.
//
// For details, see https://stackoverflow.com/a/29236630/1705598
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

var uaShortener = strings.NewReplacer(
	"~", "~~", // This is so original ~ are preserved and can be decoded losslessly.
	"Android", "~A",
	"Chrome/", "~c",
	"compatible", "~C",
	"Edge/", "~e",
	"Firefox/", "~f",
	"Gecko/", "~g",
	"(KHTML, like Gecko)", "~G",
	"iPhone", "~i",
	"Macintosh", "~I",
	"AppleWebKit/", "~a",
	"Linux", "~L",
	"Mobile/", "~m",
	"Mobile", "~M",
	"Safari/", "~s",
	"Version/", "~v",
	"Windows", "~W",
	"Mozilla/5.0 ", "~Z ", // This is to replace the prefix
)

// shortUADecoder is the inverse of uaShortener.
var shortUADecoder = strings.NewReplacer(
	"~~", "~",
	"~A", "Android",
	"~c", "Chrome/",
	"~C", "compatible",
	"~e", "Edge/",
	"~f", "Firefox/",
	"~g", "Gecko/",
	"~G", "(KHTML, like Gecko)",
	"~i", "iPhone",
	"~I", "Macintosh",
	"~a", "AppleWebKit/",
	"~L", "Linux",
	"~m", "Mobile/",
	"~M", "Mobile",
	"~s", "Safari/",
	"~v", "Version/",
	"~W", "Windows",
	"~Z ", "Mozilla/5.0 ",
)

// ShortenUserAgent can be used to simplify, shorten user agent strings.
// It replaces most frequent (and less informative) parts with short sequences.
// The transformation is a bijection, the short form can be decoded back
// into the original user agent string, see DecodeShortUA().
//
// Examples:
//   -system: Chrome Generic Win10
//        ua: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36
//     short: ~Z (~W NT 10.0; Win64; x64) ~a537.36 ~G ~c80.0.3987.132 ~s537.36
//
//   -system: Firefox Generic Linux
//        ua: Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:73.0) Gecko/20100101 Firefox/73.0
//     short: ~Z (X11; Ubuntu; ~L x86_64; rv:73.0) ~g20100101 ~f73.0
//
//   -system: Safari 13.0 macOS
//        ua: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Safari/605.1.15
//     short: ~Z (~I; Intel Mac OS X 10_15_3) ~a605.1.15 ~G ~v13.0.5 ~s605.1.15
//
//   -system: Safari Apple iPhone XR
//        ua: Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1
//     short: ~Z (~i; CPU ~i OS 12_0 like Mac OS X) ~a605.1.15 ~G ~v12.0 ~m15E148 ~s604.1
//
//   -system: Samsung Galaxy S9
//        ua: Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36
//     short: ~Z (~L; ~A 8.0.0; SM-G960F Build/R16NW) ~a537.36 ~G ~c62.0.3202.84 ~M ~s537.36
//
// The goal is not to produce the shortest output, but to provide a reasonably
// short output while maintaining readability.
func ShortenUserAgent(ua string) string {
	return uaShortener.Replace(ua)
}

// DecodeShortUA decodes the given shortened user agent string into its original form.
// The shortUA input should be the output of ShortenUserAgent().
func DecodeShortUA(shortUA string) string {
	return shortUADecoder.Replace(shortUA)
}
