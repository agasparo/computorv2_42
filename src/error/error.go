package error

import (
	"github.com/fatih/color"
	"strings"
)

func SetError(str string) {

	color.Red(str)
}

func Syntaxe(str string) (string) {

	if strings.Count(str, "=") != 1 {
		return ("You must have just one =")
	}
	return ("1")
}