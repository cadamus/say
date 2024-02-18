package say_test

import (
	"fmt"
	"testing"

	"github.com/cadamus/say"
)

func Test_Hello(t *testing.T) {
	expect := "hello, World!"
	result := say.Hello()

	if result != "hello, World!" {
		t.Error("Incorrect result: expected:", expect, ", got:", result)
	}
}

func ExampleHello() {
	fmt.Println(say.Hello())
	// Output: hello, World!
}
