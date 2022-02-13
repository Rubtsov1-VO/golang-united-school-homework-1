package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	hello := "Hello"
	world := ":world_map:"
	rendered := emoji.Sprint(hello,  world)
	rendered = rendered + "!"
	return rendered

}
