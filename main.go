package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	hello := emoji.Sprint("Hello, world!")
	fmt.Println(hello)
}
