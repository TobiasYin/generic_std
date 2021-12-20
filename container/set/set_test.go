package set

import "testing"

func TestSet(t *testing.T) {
	set := New[string]()
	set.Put("data")
	set.Put("abc")
	if !set.Exist("data") {
		t.Fatal()
	}
	if !set.Exist("abc") {
		t.Fatal()
	}
	if set.Exist("adcc") {
		t.Fatal()
	}
	if set.Len() != 2 {
		t.Fatal()
	}
	set.Remove("avvv")
	set.Remove("abc")
	if set.Len() != 1 {
		t.Fatal()
	}
	s := set.ToSlice()
	if s[0] != "data" {
		t.Fatal()
	}
}
