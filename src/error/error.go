package error

import (
	"github.com/fatih/color"
	"strings"
	"parser"
	//"fmt"
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
			if !checktab(tes) {
				return ("'" + tab[i] + "' isn't a number")
			}
		}
	}
	return ("1")
}

func Checkvars(str string) (bool) {

	str = strings.ReplaceAll(str, " ", "")

	if str == "i" || str == "ˆ" || strings.Index(str, "inf") != -1 || strings.Index(str, "nan") != -1 {
		return (false)
	}

	for i := 0; i < len(str); i++ {
		if !parser.IsLetter(string(str)) {
			return (false)
		}
	}
	return (true)
}

func Checkfuncpa(str string) (bool) {

	c := strings.Count(str, ")")
	d := strings.Count(str, "(")

	if c + d != 2 {
		return (false)
	}
	return (true)
}

func Checkfuncx(str string, str1 string) (string) {

	x := maths_functions.Getx(str)

	if x == "" {
		return ("You must have an unknown")
	}
	if strings.Count(str1, x) == 0 {
		return ("You must have '" + x + "' in your function (or not an other unknown)")
	}
	return ("1")
}

func checktab(tes []string) (bool) {

	if tes[0] != "" && !parser.IsNumeric(tes[0]) {
		return (false)
	}
	if len(tes) >= 2 && tes[1] != "" && !parser.IsNumeric(tes[1]) {
		return (false)
	}
	if tes[0] != "" && (len(tes) >= 2 && tes[1] != "") {
		return (false)
	}
	return (true)
}