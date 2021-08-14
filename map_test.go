package go_map

import (
	"strconv"
	"testing"
)

func BenchmarkMap(b *testing.B) {
	m := New()
	for i := 0; i < 1000000; i++ {
		m.Set(strconv.Itoa(i), i)
	}
	for i := 0; i < 1000000; i++ {
		m.Get(strconv.Itoa(i))
	}
	for i := 0; i < 1000000; i++ {
		m.Delete(strconv.Itoa(i))
	}
}
