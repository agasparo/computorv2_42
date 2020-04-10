package error

import (
	"github.com/fatih/color"
)

func SetError(str string) {

	color.Red(str)
}