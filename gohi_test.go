package gohi

import (
	"fmt"
	"testing"
)

func TestHi(t *testing.T) {
	s := Hi()
	fmt.Println(s)
}

func TestInc(t *testing.T) {
	n := Inc(1)
	fmt.Println(n)
}

func ExampleHi() {
	s := Hi()
	fmt.Println(s)
	// Output: Hello World!
}

func ExampleInc() {
	i := 1
	n := Inc(i)
	fmt.Println(n)
	// Output: 2
}
