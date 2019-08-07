package hello

import (
	"fmt"
	"testing"
)

func TestGoodmorning(t *testing.T) {
	if Goodmorning() != "Hello, Good Morning" {
		t.Error("Don't say Good Morning")
	}
}

func ExampleGoodmorning() {
	fmt.Println(Goodmorning())
	// Output:
	// Hello, Good Morning
}
