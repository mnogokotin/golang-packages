package utils

import (
	"fmt"
	"github.com/mnogokotin/golang-packages/logger"
	"os"
)

func Dnd(vs ...any) {
	log := logger.New("local")
	output := ""
	for i, v := range vs {
		if i != 0 {
			output += "\n"
		}
		output += fmt.Sprint(v)
	}
	log.Debug(output)
	os.Exit(1)
}
