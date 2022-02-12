package main

import (
	"fmt"
	"github.com/kyokomi/emoji/v2"
)

func main() {
	hello := emoji.Sprint("Hello :world_map:")
	fmt.Println(hello)
}
