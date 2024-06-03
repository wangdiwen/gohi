package gohi

import (
	"fmt"
	"log"
	"testing"
)

func TestHi(t *testing.T) {
	s := Hi()
	log.Println(s)
}

func TestInc(t *testing.T) {
	n := Inc(1)
	log.Println(n)
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
