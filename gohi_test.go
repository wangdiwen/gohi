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

func ExampleHi() {
	s := Hi()
	fmt.Println(s)
	// Output: Hello World!
}
