package fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	f := Fibonacci()
	fmt.Println("test")
	if f() != 1 {
		t.Fail()
	}
	if f() != 1 {
		t.Fail()
	}
	if f() != 2 {
		t.Fail()
	}
	if f() != 3 {
		t.Fail()
	}
}

func BenchmarkFibonacci(b *testing.B) {
	f := Fibonacci()
	for i := 0; i < b.N; i++ {
		f()
	}

}

func ExampleFibonacci() {
	f := Fibonacci()
	fmt.Println(f())
	// Output:
	// 1
}
