package parser

import (
	"strconv"
	"strings"
	"unicode"
	"types"
	"fmt"
	"replace_vars"
	"maps"
	"usuelles_functions"
)

type TmpComp struct {

	a float64
	b float64
}

func Array_search_count(array []string, to_search string) (res int) {

	count := 0

	for i:= 0; i < len(array); i++ {

		if array[i] == to_search {
			count++;
		}
	}
	return (count)
}


func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}

func GetCararc(str string, c string) (int) {

	var index, min int = -1, -1

	for i := 0; i < len(c); i++ {

		index = strings.Index(str, string(c[i]))
		if index != -1 && (index < min || min == -1) {
			min = index
		}
	}
	return (min)
}

func Getx(str string) (string) {

	p1 := strings.Index(str, "(")
	str = str[p1 + 1:len(str)]
	str = strings.ReplaceAll(str, ")", "")
	return (str)
}

func GetDataFunc(str string, tab map[string]types.AllT) (string, string) {

	p1 := strings.Index(str, "(")

	cmp := str[0:p1]

	for index, element := range tab {

		p2 := strings.Index(index, "(")
		if p2 > 0 && cmp == index[0:p2] {
			return index, element.Value()
		}
	}
	return "", ""
}

func Checkfunc(data map[int]string, Vars types.Variable) (map[int]string) {

	var tmp = make(map[int]string)
	var negFunc = 0

	for i := 0; i < len(data); i++ {

		parser_err := 0
		if strings.Index(data[i], "ˆ") != -1 || strings.Index(data[i], "^") != -1 {

			nData := make(map[int]string)
			nstr := strings.Split(data[i], "ˆ")
			if len(nstr) == 1 {
				nstr = strings.Split(data[i], "^")
			}
			if IsFunc(nstr[1], 1) == 1 && strings.Index(nstr[1], ")") == -1 {
				nstr[1] = nstr[1] + data[i + 1] + data[i + 2]
				data = maps.MapSlice(data, i + 1)
				data = maps.Reindex(data)
				data = maps.Clean(data)
			}
			for i := 0; i < len(nstr); i++ {
				tmp[0] = nstr[i]
				nData[i] = maps.Join(Checkfunc(tmp, Vars), "")
			}
			nData = Checkfunc(nData, Vars)
			data[i] = maps.Join(nData, "ˆ")
		}
		if len(data[i]) > 0 && data[i][0] == '-' {
			negFunc = 1
			data[i] = data[i][1:len(data[i])]
		}
		tmps := ReplaceTmp(data[i])
		if IsFunc(tmps, 1) == 1 {
			data[i] = tmps
			p1 := strings.Index(data[i], "(")
			p2 := strings.Index(data[i], ")")
			add := 1
			for p2 = p2; p2 < 0; p2 = strings.Index(data[i], ")") {
				data[i] += data[i + add]
				add++
			}
			if add > 1 {
				data = maps.MapSlice(data, i + 1)
			}
			if negFunc == 1 {
				ntn := make(map[int]string)
				ntn[0] = "*"
				ntn[1] = "-1)"
				ntn[2] = data[i + 1]
				data[i] = "(" + data[i]
				data = maps.CombineN(data, ntn, i + 1)
				i = -1
			} else {
				name, value := GetDataFunc(data[i], Vars.Table)
				x := Getx(name)
				r := data[i][p1 + 1:p2]
				if IsExpression(r, x, Vars) {
					data[0] = "You must have only one number for unknown not an expression"
					return (data)
				}
				if strings.Index(value, "|") != -1 {
					value = strings.ReplaceAll(value, "usu|", "")
					value = Remp(value, x, r, Vars)
					value = usuelles_functions.GetUsuF(value, Vars)
					if strings.Index(value, "Impossible") != -1 {
						data[0] = value
						return (data)
					}
				}
				nstr := strings.ReplaceAll(value, x, replace_vars.GetVars(&Vars, r))
				tmp := strings.ReplaceAll(data[i], Remp(name, x, r, Vars), "(" + nstr + ")")
				nt := GetAllIma(strings.ReplaceAll(tmp, " ", ""), &parser_err)
				data = maps.CombineN(data, nt, i)
				i = -1
			}
		} else if negFunc == 1 {
			data[i] = "-" + data[i]
		}
		negFunc = 0
	}
	return (data)
}

func ReplaceTmp(str string) (string) {

	r := 0
	for i := 0; i < len(str) && str[i] == '('; i++ {
		r++
	} 
	return (str[r:len(str)])
}

func IsExpression(str string, x string, Vars types.Variable) (bool) {

	if str == x || Is_defined(str, Vars) {
		return (false)
	}

	m := strings.Count(str, "-")
	m += strings.Count(str, "*")
	m += strings.Count(str, "/")
	m += strings.Count(str, "+")
	m += strings.Count(str, "%")
	m += strings.Count(str, "=")

	if m >= 2 {
		return (true)
	}
	if str[0] == '*' || str[0] == '/' || str[0] == '%' {
		return (true)
	}
	if str[0] == '+' || str[0] == '-' {
		str = str[0:len(str)]
	}
	index := GetCararc(str, "+-/*%")
	if index == -1 && !IsNumeric(str) {
		return (true)
	}
	if index == -1 && IsNumeric(str) {
		return (false)
	}
	e := strings.Split(str, string(str[index]))
	c := 0
	for i := 0; i < len(e); i++ {

		if !IsNumeric(e[i]) && e[i] != "" {
			return (true)
		} else if e[i] != "" {
			c++
		}
	}
	if c > 1 {
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

func Remp(str string, x string, r string,  Vars types.Variable) (string) {

	p3 := strings.Index(str, "(")
	p4 := strings.Index(str, ")")

	if p3 < 0 || p4 < 0 {
		return (strings.ReplaceAll(str, x, r))
	}
	as := strings.ReplaceAll(str[p3 + 1:p4], x, r)
	str = str[0:p3 + 1] + as + str[p4:len(str)]
	return (str)
}

func IsFunc(str string, t int) (int) {

	p1 := strings.Index(str, "(")
	p2 := strings.Index(str, ")")

	if p1 < 0 {
		return (0)
	}

	if !IsLetter(str[0:p1]) || p1 == 0 {
		return (0)
	}

	if p2 < 0 {
		return (1)
	}

	if t == 0 && !IsLetter(str[p1 + 1:p2]) {
		return (0)
	}

	if p1 != -1 && p2 != -1 && p1 < p2 {

		return (1)
	}
	return (0)
}

func IsLetter(s string) bool {

    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func Float2string(Calc TmpComp) (string) {

	if Calc.b == 0 {
		return (fmt.Sprintf("%f", Calc.a))
	} else if Calc.a == 0 {
		return (fmt.Sprintf("%fi", Calc.b))
	} else if Calc.b > 0 {
		return (fmt.Sprintf("%f + %fi", Calc.a, Calc.b))
	}
	return (fmt.Sprintf("%f %fi", Calc.a, Calc.b))
}

func GetAllIma(str string, pos *int) (map[int]string) {

	data := make(map[int]string)
	var itab, neg, ajj int
	var anc string

	if len(str) == 0 {
		return (data)
	}
	if strings.Index("+-/*%", string(str[len(str) - 1])) != -1 {
		*pos = 1
		return (data)
	}
	if !DebCheck(str) {
		*pos = 1
		return (data)
	}
	if str[0] == '+' {
		str = str[1:len(str)]
	}

	if str[0] == '*' || str[0] == '/' {
		*pos = 1
		return (data)
	}

	itab = 0
	anc = ""
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			
			str = str[1:len(str)]
			neg = 1
		}
		if str[i] == '+' {
			str = str[1:len(str)]
		}
		if i == len(str) {
			return (data)
		}
		if Impossible(string(str[i]), anc) {
			*pos = 1
			return (data)
		}
		index := GetCararc(str, "+-/*%")
		if index + 1 == len(str) || index == 0 {
			return (data)
		}
		if index >= 0 && str[index] == '*' && str[index + 1] == '*' {
			ajj = 1
		}
		if index == -1 {
			if neg == 1 {
				data[itab] = "-" + str
			} else {
				data[itab] = str
			}
			return (data)
		}
		if index > 0 {
			if neg == 1 {
				data[itab] = "-" + str[i:index]
			} else {
				data[itab] = str[i:index]
			}
			itab++
			neg = 0
			i = i + index
			data[itab] = string(str[i])
			itab++
			if ajj == 1 {
				data[itab] = string(str[i])
				itab++
			}
			str = str[i + 1 + ajj:len(str)]
		}
		anc = data[len(data) - 1]
		i = -1
		ajj = 0
	}
	return (data)
}

func DebCheck(str string) (bool) {

	if len(str) == 1 {

		if strings.Index("+-/*%", string(str[0])) != -1 {
			return (false)
		}
		return (true)
	}

	if strings.Index("+-/*%", string(str[0])) != -1 && strings.Index("+-/*%", string(str[1])) != -1 {
		return (false)
	}
	return (true)
}

func Impossible(str string, cmp string) (bool) {

	if cmp == "" {
		return (false)
	}
	if (cmp == "*" || cmp == "/") && str == "/" {
		return (true)
	}
	if cmp == "/" && str == "*" {
		return (true)
	}
	return (false)
}