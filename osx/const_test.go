package osx

import (
	"reflect"
	"testing"
)

func TestArchConsts(t *testing.T) {
	arch32bit := reflect.TypeOf(0).Size() == 4

	if arch32bit != Arch32bit {
		t.Errorf("Arch32bit is incorrect!")
	}
	if arch32bit == Arch64bit {
		t.Errorf("Arch64bit is incorrect!")
	}
}
