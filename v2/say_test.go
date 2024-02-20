package say_test

import (
	"fmt"

	"github.com/cadamus/say/v2"
)

func ExampleHello() {
	fmt.Println(say.Hello(""))
	// Output: hello, World!
}

func ExampleHello_withName() {
	fmt.Println(say.Hello("Golang"))
	// Output: hello, Golang!
}

func ExampleGoodbye() {
	fmt.Println(say.Goodbye())
	// Output: I am just going outside and I may be some time
}
