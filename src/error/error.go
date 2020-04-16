package error

import (
	"github.com/fatih/color"
	"strings"
	"parser"
	"fmt"
	"maths_functions"
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

func In(tab map[int]string, t int, f string) (string) {

	a := 0

	fmt.Println(tab)
	if tab[0] == "-" || tab[0] == "+" {
		a++
	}
	for i := a; i < len(tab); i += 2 {
		if !parser.IsNumeric(tab[i]) && t == 0 {
			return ("'" + tab[i] + "' isn't a number")
		}
		if t == 1 {
			x := maths_functions.Getx(f)
			tes := strings.Split(strings.ReplaceAll(tab[i], " ", ""), x)
			fmt.Println(tes)
			if !checktab(tes) {
				return ("'" + tab[i] + "' isn't a number")
			}
		}
	}
	return ("1")
}

func checktab(tes []string) (bool) {

	if tes[0] != "" && !parser.IsNumeric(tes[0]) {
		return (false)
	}
	if tes[1] != "" && !parser.IsNumeric(tes[1]) {
		return (false)
	}
	return (true)
}