package say_test

import (
	"fmt"

	"github.com/cadamus/say"
)

func ExampleHello() {
	fmt.Println(say.Hello())
	// Output: hello, World!
}

func ExampleGoodbye() {
	fmt.Println(say.Goodbye())
	// Output: I am just going outside and I may be some time
}
