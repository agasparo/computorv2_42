package error

import (
	"github.com/fatih/color"
	"strings"
	"parser"
	"maths_functions"
	"types"
	"fmt"
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

func In(tab map[int]string, t int, f string, Dat types.Variable) (string) {

	a := 0
	is_i := 0

	if tab[0] == "-" || tab[0] == "+" {
		a++
	}
	for i := a; i < len(tab); i += 2 {

		if strings.Index(tab[i], "i") != -1 && tab[i] != "i" && t == 0 && !IsUsu(tab, Dat) && !IsPower(tab[i]) {
			if strings.Count(tab[i], "i") > 1 {
				
				if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) {
					return ("'" + tab[i] + "' isn't defined 1")
				}
			}
			tab[i] = strings.ReplaceAll(tab[i], "i", "")
			is_i = 1
		}

		if !parser.IsNumeric(tab[i]) && t == 0 && tab[i] != "i" && !IsPower(tab[i]) {

			if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) {
				return ("'" + tab[i] + "' isn't defined 2")
			}
		}
		if t == 1 {
			x := maths_functions.Getx(f)
			tes := strings.Split(strings.ReplaceAll(tab[i], " ", ""), x)
			fmt.Println(tes)
			if !checktab(tes) && !Is_defined(strings.Join(tes, ""), Dat) {
				return ("'" + tab[i] + "' isn't defined 3")
			}
		}
		if is_i == 1 {
			tab[i] += "i"
		}
		is_i = 0
	}
	return ("1")
}

func IsPower(str string) (bool) {

	if strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1 {

		return (true)
	}
	return (false)
}

func Is_defined(str string, vars types.Variable) (bool) {

	if _, ok := vars.Table[strings.ToLower(str)]; ok {

		return (true)
    }
    return (false)
}

func IsUsu(data map[int]string, vars types.Variable) (bool) {
	
	for i := 0; i < len(data); i++ {

		p1 := strings.Index(data[i], "(")
		if p1 == -1 {
			return (false)
		}
		nstr := data[i][0:p1] + "(x)"
		if _, ok := vars.Table[strings.ToLower(nstr)]; ok {

			return (true)
    	}
	}
	return (false)
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

	if tes[0] != "" && !parser.IsNumeric(tes[0]) && !IsPower(tes[0]) {
		return (false)
	}
	if len(tes) >= 2 && tes[1] != "" && !parser.IsNumeric(tes[1]) && !IsPower(tes[1]) {
		return (false)
	}
	if tes[0] != "" && (len(tes) >= 2 && tes[1] != "") {
		return (false)
	}
	return (true)
}