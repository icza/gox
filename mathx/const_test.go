package mathx

import "testing"

func TestMinMaxUint(t *testing.T) {
	var max uint
	max = ^max

	if MaxUint != max {
		t.Errorf("MaxUint is not max!")
	}
	if MinUint != max+1 {
		t.Errorf("MinUint is not min!")
	}
}

func TestMinMaxInt(t *testing.T) {
	maxuint := uint(0)
	maxuint = ^maxuint
	max := int(maxuint >> 1)

	if MaxInt != max {
		t.Errorf("MaxInt is not max!")
	}
	if MinInt != -max-1 {
		t.Errorf("MinInt is not min!")
	}
}
