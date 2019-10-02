package foo

import (
	"fmt"

	superlog "github.com/tarotaro0/exp-mm-repo/conf"
)

// HelloFoo say "Hello, Foo"
func HelloFoo() {
	fmt.Println("Hello, Foo v1.1.0")
	superlog.SuperLog("Foo")
}
