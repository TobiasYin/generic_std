package atomic_test

import (
	"math/rand"
	"runtime"
	"testing"

	. "github.com/TobiasYin/generic_std/sync/atomic"
)

func TestValue(t *testing.T) {
	var v Value[int]
	if v.Load() != 0 {
		t.Fatal("initial Value is not nil")
	}
	v.Store(42)
	x := v.Load()
	if x != 42 {
		t.Fatalf("wrong value: got %+v, want 42", x)
	}
	v.Store(84)
	x = v.Load()
	if x != 84 {
		t.Fatalf("wrong value: got %+v, want 84", x)
	}
}

func TestValueLarge(t *testing.T) {
	var v Value[string]
	v.Store("foo")
	x := v.Load()
	if x != "foo" {
		t.Fatalf("wrong value: got %+v, want foo", x)
	}
	v.Store("barbaz")
	x = v.Load()
	if x != "barbaz" {
		t.Fatalf("wrong value: got %+v, want barbaz", x)
	}
}

func TestValueConcurrent(t *testing.T) {
	tests := [][]uint64{
		{
			uint64(0), ^uint64(0), uint64(1 + 2<<32), uint64(3 + 4<<32),
		},
	}
	p := 4 * runtime.GOMAXPROCS(0)
	N := int(1e5)
	if testing.Short() {
		p /= 2
		N = 1e3
	}
	for _, test := range tests {
		var v Value[uint64]
		done := make(chan bool, p)
		for i := 0; i < p; i++ {
			go func() {
				r := rand.New(rand.NewSource(rand.Int63()))
				expected := true
			loop:
				for j := 0; j < N; j++ {
					x := test[r.Intn(len(test))]
					v.Store(x)
					x = v.Load()
					for _, x1 := range test {
						if x == x1 {
							continue loop
						}
					}
					t.Logf("loaded unexpected value %+v, want %+v", x, test)
					expected = false
					break
				}
				done <- expected
			}()
		}
		for i := 0; i < p; i++ {
			if !<-done {
				t.FailNow()
			}
		}
	}
}

func BenchmarkValueRead(b *testing.B) {
	var v Value[*int]
	v.Store(new(int))
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			x := v.Load()
			if *x != 0 {
				b.Fatalf("wrong value: got %v, want 0", *x)
			}
		}
	})
}
