package stringsx

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
	"unsafe"
)

func checkSame(t *testing.T, s1, s2 string, shouldBeSame bool) {
	t.Helper()

	// Just to be sure:
	if s1 != s2 {
		t.Errorf("Excected match")
	}

	hdr1 := (*reflect.StringHeader)(unsafe.Pointer(&s1))
	hdr2 := (*reflect.StringHeader)(unsafe.Pointer(&s2))

	if shouldBeSame {
		if hdr1.Data != hdr2.Data {
			t.Errorf("Expected to be same")
		}
	} else {
		if hdr1.Data == hdr2.Data {
			t.Errorf("Expected to not be the same")
		}

	}

	runtime.KeepAlive(s1)
	runtime.KeepAlive(s2)
}

func TestPool(t *testing.T) {
	x := int64(827364536372)

	s1 := fmt.Sprint(x)
	s2 := fmt.Sprint(x)

	checkSame(t, s1, s1, true) // This is just to test checkSame
	checkSame(t, s1, s2, false)

	pool := NewPool()
	s1i := pool.Interned(s1)
	checkSame(t, s1, s1i, true)

	s2i := pool.Interned(s2)
	checkSame(t, s2, s2i, false) // This will be false: pool will return s1
	checkSame(t, s1, s2i, true)
}
