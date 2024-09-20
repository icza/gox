package colorx

import (
	"errors"
	"image/color"
)

var errInvalidFormat = errors.New("invalid format")

// ParseHexColor parses a web color given by its hex RGB format.
// See https://en.wikipedia.org/wiki/Web_colors for input format.
//
// For details, see https://stackoverflow.com/a/54200713/1705598
func ParseHexColor(s string) (c color.RGBA, err error) {
	if len(s) == 0 || s[0] != '#' {
		return c, errInvalidFormat
	}

	c.A = 0xff

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
