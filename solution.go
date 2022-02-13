package solution

import (
	"fmt"

    "github.com/kyokomi/emoji/v2"
)

func GetMessage() {
	hello := emoji.Sprint("Hello, :world_map:!")
	fmt.Println(hello)
}
