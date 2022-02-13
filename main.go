package main

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func main() {
	e := emoji.Sprint(":world_map:")

	fmt.Println("Hello," ,e , "!")

}
