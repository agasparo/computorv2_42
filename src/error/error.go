package error

import (
	"github.com/fatih/color"
	"strings"
	"parser"
	"maths_functions"
	"types"
	"fmt"
	"maps"
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

func In(data map[int]string, t int, f string, Dat types.Variable) (string) {

	a := 0
	is_i := 0
	tab := data
	neg := 0
	ajj := -1

	if tab[0] == "-" || tab[0] == "+" {
		a++
	}
	for i := a; i < len(tab); i += 2 {

		if tab[i][0] == '-' || tab[i][0] == '+' {
			if tab[i][0] == '-' {
				neg = 1
			}
			tab[i] = tab[i][1:len(tab[i])]
		}

		if strings.Index(tab[i], "(") != -1 && strings.Index(tab[i], ")") == -1 && i + 1 < len(tab) && !parser.IsNumeric(tab[i + 1]) {

			index := maps.Array_search(tab, ")")
			if index == -1 {
				return ("'" + tab[i] + "' isn't defined 0")
			}
			tab[i] = maps.Add(tab, tab[i], i + 1, index + 1)
			tab = maps.MapSliceCount(tab, i + 1, i - index)
			tab = maps.Reindex(tab)
			tab = maps.Clean(tab)
			ajj = index + 1
		}

		if strings.Index(tab[i], "i") != -1 && tab[i] != "i" && t == 0 && !IsUsu(tab, Dat) && !IsPower(tab[i], Dat, 0) {
			if strings.Count(tab[i], "i") > 1 {
				
				if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) {
					return ("'" + tab[i] + "' isn't defined 1")
				}
			}
			tab[i] = strings.ReplaceAll(tab[i], "i", "")
			is_i = 1
		}

		fmt.Println(IsPower(tab[i], Dat, 0))
		if !parser.IsNumeric(tab[i]) && t == 0 && tab[i] != "i" && !IsPower(tab[i], Dat, 0) {

			if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) && !ResFunct(tab[i], Dat) {
				return ("'" + tab[i] + "' isn't defined 2")
			}
		}
		if t == 1 {
			x := maths_functions.Getx(f)
			tes := strings.Split(strings.ReplaceAll(tab[i], " ", ""), x)
			if !checktab(tes, Dat) && !Is_defined(strings.Join(tes, ""), Dat) && !IsUsu(tab, Dat) {
				if !ResFunct(tab[i], Dat) {
					return ("'" + tab[i] + "' isn't defined 3")
				}
			}
			if IsUsu(tab, Dat) {
				return ("you can't have an usuel function in your function")
			}
		}
		if is_i == 1 {
			tab[i] += "i"
		}
		if neg == 1 {
			tab[i] = "-" + tab[i]
		}
		if ajj != -1 {
			i += ajj - 2
		}
		is_i = 0
		neg = 0
		ajj = -1
	}
	return ("1")
}

func ResFunct(str string, Dat types.Variable) (bool) {

	p1 := strings.Index(str, "(")
	if p1 == -1 {
		return (false)
	}
	p2 := strings.Index(str, ")")
	if p2 == -1 {
		return (false)
	}

	if !parser.IsNumeric(str[p1 + 1:p2]) { 
		return (false)
	}
	return (true)
}

func IsPower(str string, Dat types.Variable, t int) (bool) {

	if strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1 {

		nstr := strings.Split(str, "ˆ")
		if len(nstr) == 1 {
			nstr = strings.Split(str, "^")
		}
		if t == 0 {
			
			for i := 0; i < len(nstr); i++ {

				if nstr[i] == "" {
					return (false)
				}

				if parser.IsFunc(nstr[i], 1) == 1 {
					p1 := strings.Index(nstr[i], "(")
					p2 := strings.Index(nstr[i], ")")
					nstr[i] = nstr[i][p1 + 1:p2]
				}

				if !Is_defined(nstr[i], Dat) && !parser.IsNumeric(nstr[i]) {
					return (false)
				}
			}
		}
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
		p2 := strings.Index(data[i], ")")
		if p2 == -1 {
			return (false)
		}
		if !parser.IsNumeric(data[i][p1 + 1:p2]) { 
			nstr := data[i][0:p1] + "(x)"
			if _, ok := vars.Table[strings.ToLower(nstr)]; ok {
				return (true)
    		}
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

func Checkfuncx(str string, str1 string, vars types.Variable) (string) {

	x := maths_functions.Getx(str)

	if x == "" {
		return ("You must have an unknown")
	}
	if strings.Count(str1, x) == 0 {
		return ("You must have '" + x + "' in your function (or not an other unknown)")
	}
	cmp := strings.ToLower(str)
	if val, ok := vars.Table[cmp]; ok {

		if strings.Index(val.Value(), "usu|") != -1 {
			return ("Your fonction can't be with the same name like usual function")
		}
    }
	return ("1")
}

func checktab(tes []string, Dat types.Variable) (bool) {

	if tes[0] != "" && !parser.IsNumeric(tes[0]) && !IsPower(tes[0], Dat, 1) {
		return (false)
	}
	if len(tes) >= 2 && tes[1] != "" && !parser.IsNumeric(tes[1]) && !IsPower(tes[1], Dat, 1) {
		return (false)
	}
	if tes[0] != "" && (len(tes) >= 2 && tes[1] != "") {
		return (false)
	}
	return (true)
}