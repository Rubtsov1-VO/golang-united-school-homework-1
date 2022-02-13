package solution

import (
	"fmt"

	"github.com/kyokomi/emoji/v2"
)

func solution() {
	hello := emoji.Sprint("Hello, :world_map:!")
	fmt.Println(hello)
}
