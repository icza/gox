package colorx

import (
	"fmt"
	"testing"
)

// ExampleParseHexColor shows how to use the ParseHexColor() function.
func ExampleParseHexColor() {
	hexCols := []string{
		"#112233",
		"#123",
		"#000233",
		"#023",
		"#bAC",
		"invalid",
		"#abcd",
		"#-12",
	}
	for _, hc := range hexCols {
		c, err := ParseHexColor(hc)
		fmt.Printf("%-7s = %3v, %v\n", hc, c, err)
	}

	// Output:
	// #112233 = { 17  34  51 255}, <nil>
	// #123    = { 17  34  51 255}, <nil>
	// #000233 = {  0   2  51 255}, <nil>
	// #023    = {  0  34  51 255}, <nil>
	// #bAC    = {187 170 204 255}, <nil>
	// invalid = {  0   0   0 255}, invalid format
	// #abcd   = {  0   0   0 255}, invalid format
	// #-12    = {  0  17  34 255}, invalid format
}

func TestParseHexColor(t *testing.T) {
	_, err := ParseHexColor("")
	if err == nil {
		t.Errorf("empty string should return error")
	}
}
