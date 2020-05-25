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

func TestUAShortening(t *testing.T) {
	cases := []struct {
		name    string
		ua      string
		shortUA string
	}{
		{"empty", "", ""},
		{"decodability", "~m~~~A~", "~~m~~~~~~A~~"},
		{
			"Chrome Generic Win10",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
			"~Z (~W NT 10.0; Win64; x64) ~a537.36 ~G ~c80.0.3987.132 ~s537.36",
		},
		{
			"Firefox Generic Win10",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:73.0) Gecko/20100101 Firefox/73.0",
			"~Z (~W NT 10.0; Win64; x64; rv:73.0) ~g20100101 ~f73.0",
		},
		{
			"Safari 13.0 macOS",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.5 Safari/605.1.15",
			"~Z (~I; Intel Mac OS X 10_15_3) ~a605.1.15 ~G ~v13.0.5 ~s605.1.15",
		},
		{
			"Chrome Generic macOS",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.132 Safari/537.36",
			"~Z (~I; Intel Mac OS X 10_15_3) ~a537.36 ~G ~c80.0.3987.132 ~s537.36",
		},
		{
			"Firefox Generic Linux",
			"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:73.0) Gecko/20100101 Firefox/73.0",
			"~Z (X11; Ubuntu; ~L x86_64; rv:73.0) ~g20100101 ~f73.0",
		},
		{
			"Edge 18.0 Win10",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18362",
			"~Z (~W NT 10.0; Win64; x64) ~a537.36 ~G ~c70.0.3538.102 ~s537.36 ~e18.18362",
		},
		{
			"Opera Generic Win10",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.72",
			"~Z (~W NT 10.0; Win64; x64) ~a537.36 ~G ~c79.0.3945.130 ~s537.36 OPR/66.0.3515.72",
		},
		{
			"Safari Apple iPhone XR",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
			"~Z (~i; CPU ~i OS 12_0 like Mac OS X) ~a605.1.15 ~G ~v12.0 ~m15E148 ~s604.1",
		},
		{
			"Samsung Galaxy S9",
			"Mozilla/5.0 (Linux; Android 8.0.0; SM-G960F Build/R16NW) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/62.0.3202.84 Mobile Safari/537.36",
			"~Z (~L; ~A 8.0.0; SM-G960F Build/R16NW) ~a537.36 ~G ~c62.0.3202.84 ~M ~s537.36",
		},
		{
			"Googlebot 2.1",
			"Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.92 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
			"~Z (~L; ~A 6.0.1; Nexus 5X Build/MMB29P) ~a537.36 ~G ~c80.0.3987.92 ~M ~s537.36 (~C; Googlebot/2.1; +http://www.google.com/bot.html)",
		},
		{
			"AhrefsBot 6.1",
			"Mozilla/5.0 (compatible; AhrefsBot/6.1; +http://ahrefs.com/robot/)",
			"~Z (~C; AhrefsBot/6.1; +http://ahrefs.com/robot/)",
		},
		{
			"Bytespider",
			"Mozilla/5.0 (Linux; Android 5.0) AppleWebKit/537.36 (KHTML, like Gecko) Mobile Safari/537.36 (compatible; Bytespider; https://zhanzhang.toutiao.com/)",
			"~Z (~L; ~A 5.0) ~a537.36 ~G ~M ~s537.36 (~C; Bytespider; https://zhanzhang.toutiao.com/)",
		},
		{
			"Nimbostratus-Bot 1.3.2",
			"Mozilla/5.0 (compatible; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)",
			"~Z (~C; Nimbostratus-Bot/v1.3.2; http://cloudsystemnetworks.com)",
		},
		{
			"YandexBot 3.0",
			"Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)",
			"~Z (~C; YandexBot/3.0; +http://yandex.com/bots)",
		},
		{
			"Applebot 0.1",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_10_1) AppleWebKit/600.2.5 (KHTML, like Gecko) Version/8.0.2 Safari/600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
			"~Z (~I; Intel Mac OS X 10_10_1) ~a600.2.5 ~G ~v8.0.2 ~s600.2.5 (Applebot/0.1; +http://www.apple.com/go/applebot)",
		},
		{
			"SemrushBot 6",
			"compatible; SemrushBot/6~bl; +http://www.semrush.com/bot.html",
			"~C; SemrushBot/6~~bl; +http://www.semrush.com/bot.html",
		},
	}

	for _, c := range cases {
		if got := ShortenUserAgent(c.ua); got != c.shortUA {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.shortUA, got)
		}
		if got := DecodeShortUA(c.shortUA); got != c.ua {
			t.Errorf("[%s] Expected: %s, got: %s", c.name, c.ua, got)
		}
	}
}
