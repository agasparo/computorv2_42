package error

import (
	"github.com/fatih/color"
	"strings"
	"parser"
	"maths_functions"
	"types"
	"maps"
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

func In(data map[int]string, t int, f string, Dat types.Variable) (string) {

	a := 0
	is_i := 0
	tab := maps.Copy(data)
	neg := 0
	ajj := -1
	tmp := ""

	if strings.Count(maps.Join(data, ""), ")") != strings.Count(maps.Join(data, ""), "(") {
		return ("You must have the same number of parentheses")
	}

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
		tab[i] = ReplaceTmp(tab[i])
		if (parser.IsFunc(tab[i], 1) == 1 || parser.IsFunc(tab[i], 0) == 1) && strings.Index(tab[i], "(") != -1 && strings.Index(tab[i], ")") == -1 && i + 1 < len(tab) && !parser.IsNumeric(tab[i + 1]) {
			fmt.Println("ici")
			index := maps.Array_search(tab, ")")
			if index == -1 {
				return ("'" + tab[i] + "' isn't defined 1")
			}
			tab[i] = maps.Add(tab, tab[i], i + 1, index + 1)
			tab = maps.MapSliceCount(tab, i + 1, i - index)
			tab = maps.Reindex(tab)
			tab = maps.Clean(tab)
			ajj = index + 1
		} else if parser.IsFunc(tab[i], 1) != 1 && parser.IsFunc(tab[i], 0) != 1 {
			if !ParaCheck(tab[i]) {
				return ("you must have n * (z not or z) * n not '" + tab[i] + "'")
			}
			if strings.Index(tab[i], "ˆ") == -1 && strings.Index(tab[i], "^") == -1 {
				tab[i] = strings.ReplaceAll(tab[i], "(", "")
				tab[i] = strings.ReplaceAll(tab[i], ")", "")
			} else {
				nstr := strings.Split(tab[i], "ˆ")
				if len(nstr) == 1 {
					nstr = strings.Split(tab[i], "^")
				}
				if parser.IsFunc(nstr[1], 1) == 1 || parser.IsFunc(nstr[1], 0) == 1 {
					
					if len(tab) > i + 1 {
						tab[i] = tab[i] + tab[i + 1] + tab[i + 2]
						tab = maps.MapSlice(tab, i + 1)
						tab = maps.Reindex(tab)
						tab = maps.Clean(tab)
						ajj = i + 1
					}
				} else {
					tab[i] = strings.ReplaceAll(tab[i], "(", "")
					tab[i] = strings.ReplaceAll(tab[i], ")", "")
				}
			}
		}

		if strings.Index(tab[i], "ˆ") != -1 || strings.Index(tab[i], "^") != -1 {
			nstr := strings.Split(tab[i], "ˆ")
			if len(nstr) == 1 {
				nstr = strings.Split(tab[i], "^")
			}
			if tab[i + 1] == "-" && nstr[1] == "" {
				tab[i] = tab[i] + tab[i + 1] + tab[i + 2]
				tab = maps.MapSlice(tab, i + 1)
				tab = maps.Reindex(tab)
				tab = maps.Clean(tab)
				ajj = i + 1
			}
		}

		if strings.Index(tab[i], "i") != -1 && tab[i] != "i" && t == 0 && !IsUsu(tab, Dat) && !IsPower(tab[i], Dat, 0) {
			if strings.Count(tab[i], "i") > 1 {
				if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) {
					return ("'" + tab[i] + "' isn't defined 2")
				}
			}
			is_i = 1
			tmp = tab[i]
			tab[i] = strings.ReplaceAll(tab[i], "i", "")
		}

		if !parser.IsNumeric(tab[i]) && t == 0 && tab[i] != "i" && !IsPower(tab[i], Dat, 0) {

			if is_i == 1 {
				tab[i] = tmp
			}
			if !IsUsu(tab, Dat) && !Is_defined(tab[i], Dat) && !ResFunct(tab[i], Dat) {
				tab[i] = strings.ReplaceAll(tab[i], "i", "")
				if strings.Index(tab[i], "(") != -1 {
					x := maths_functions.Getx(tab[i])
					if !Is_defined(x, Dat) {
						return ("'" + tab[i] + "' isn't defined 4")
					}
				} else {
					if strings.Index(tab[i], "[") == -1 && strings.Index(tab[i], "]") == -1 {
						if tab[i] != "*" {
							if tab[i] == "" {
								return ("you have a problem with your parentheses syntaxe")
							}
							if strings.Index(tab[i], ",") == -1 {
								return ("'" + tab[i] + "' isn't defined 5")
							}
						}
					}
				}
			}
		}
		if t == 1 {
			x := maths_functions.Getx(f)
			tes := strings.Split(strings.ReplaceAll(tab[i], " ", ""), x)
			if !checktab(tes, Dat, tab[i], x) && !Is_defined(strings.Join(tes, ""), Dat) && !IsUsu(tab, Dat) {
				if !ResFunct(tab[i], Dat) && parser.IsFunc(tab[i], 0) != 1 && strings.Index(tab[i], "i") == -1 {
					if strings.Index(tab[i], "[") == -1 && strings.Index(tab[i], "]") == -1 {
						if tab[i] != "*" {
							return ("'" + tab[i] + "' isn't defined 3")
						}
						if tab[i] == "*" && strings.Index(tab[i + 1], "]") == -1 {
							return ("'" + tab[i] + "' isn't defined 3.1")
						}
					}
				}
			}
			if IsUsu(tab, Dat) {
				return ("you can't have an usuel function in your function")
			}
		}
		if is_i == 1 {
			tab[i] = tmp
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
		tmp = ""
	}
	return ("1")
}

func ReplaceTmp(str string) (string) {

	r := 0
	for i := 0; str[i] == '('; i++ {
		r++
	}
	return (str[r:len(str)])
}

func ParaCheck(str string) (bool) {

	tes := strings.Split(strings.ReplaceAll(str, " ", ""), "(")
	if len(tes) == 2 {
		if parser.IsNumeric(tes[0]) && parser.IsNumeric(tes[1]) {
			return (false)
		}
		tes1 := strings.Split(strings.ReplaceAll(str, " ", ""), ")")
		if parser.IsNumeric(tes1[0]) && parser.IsNumeric(tes1[1]) {
			return (false)
		}
	}
	return (true)
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

	for index, _ := range Dat.Table {

		p3 := strings.Index(index, "(")		
		if p3 != -1 {
			if str[0:p1 + 1] == index[0:p3 + 1] {
				return (true)
			}
		}
	}
	return (false)
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
			if val, ok := vars.Table[strings.ToLower(nstr)]; ok {
				if strings.Index(val.Value(), "|") != -1 {
					return (true)
				}
    		}
    	}
	}
	return (false)
}

func Checkvars(str string) (bool) {

	e := strings.Split(str, " ")
	if len(e) > 1 {
		return (false)
	}

	str = strings.ReplaceAll(str, " ", "")

	if str == "" {
		return (false)
	}

	if strings.Index(str, "mat") != -1 {
		return (false)
	}

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
    p1 := strings.Index(str, "(")
    for index, _ := range vars.Table {

		p3 := strings.Index(index, "(")		
		if p3 != -1 {
			if str[0:p1 + 1] == index[0:p3 + 1] && str != index {
				return ("A function with name '" + index[0:p3] + "' is already defined")
			} 
		}
	}
	return ("1")
}

func checktab(tes []string, Dat types.Variable, cmp string, x string) (bool) {

	if tes[0] != "" && !parser.IsNumeric(tes[0]) && !IsPower(tes[0], Dat, 1) {
		return (false)
	}
	if len(tes) >= 2 && tes[1] != "" && !parser.IsNumeric(tes[1]) && !IsPower(tes[1], Dat, 1) {
		return (false)
	}
	if tes[0] == "" && tes[1] == "" && x != cmp {
		return (false)
	}
	if tes[0] != "" && (len(tes) >= 2 && tes[1] != "") {
		if strings.Index(tes[1], "ˆ") == -1 && strings.Index(tes[1], "^") == -1 {
			return (false)
		}
	}
	return (true)
}