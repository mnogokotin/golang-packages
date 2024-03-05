package d

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func Dnd(args ...any) {
	for _, arg := range args {
		farg := fmt.Sprintf("%+v", arg)
		fmt.Println(color.GreenString(farg))
	}
	os.Exit(1)
}
