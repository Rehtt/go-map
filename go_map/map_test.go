package go_map

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := New()
	m.Set("test1", 1)
	m.Set("test2", "2")
	m.Set("test3", true)

	fmt.Println(m.Pull2List())
	fmt.Println(m.Pull2Map())

	fmt.Println(m.Get("test2"))

	m.Delete("test2")

	fmt.Println(m.Pull2List())
	fmt.Println(m.Pull2Map())
}
