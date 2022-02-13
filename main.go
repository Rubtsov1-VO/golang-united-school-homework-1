package solution

import (
	"github.com/kyokomi/emoji/v2"
	"fmt"
)

func GetMessage() string {
	hello := "Hello"
	world := ":world_map:"
	rendered := emoji.Sprint(hello, world)
	rendered = rendered + " !"
	return fmt.Println(rendered)

}
