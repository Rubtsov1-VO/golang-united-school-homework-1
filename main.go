package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	e := emoji.Sprint(":world_map:")

	fmt.Println(e, []rune(e))

	for _, v := range e {
		fmt.Printf("%[1]T %[1]v\n", v)
	}
}
