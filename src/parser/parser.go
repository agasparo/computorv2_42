package parser

import (
	"strconv"
	"strings"
	"unicode"
	"types"
	"fmt"
	"maths_imaginaires"
	"replace_vars"
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

func Calc(fu string, x string, r string, vars *types.Variable) (float64, float64) {

	fu = strings.ReplaceAll(fu, x, r)
	data := GetAllIma(fu)
	data = maths_imaginaires.CalcMulDivi(data, vars, x)
	data = maths_imaginaires.CalcAddSous(data, vars, x)
	return (maths_imaginaires.ParseOne(data[0], vars))
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

	for i := 0; i < len(data); i++ {

		if IsFunc(data[i], 1) == 1 {
			name, value := GetDataFunc(data[i], Vars.Table)
			x := Getx(name)
			p1 := strings.Index(data[i], "(")
			p2 := strings.Index(data[i], ")")
			r := data[i][p1 + 1:p2]
			a, b := Calc(value, x, replace_vars.GetVars(&Vars, r), &Vars)
			data[i] = "(" + Float2string(TmpComp{ a, b }) + ")"
		}
	}
	return (data)
}

func IsFunc(str string, t int) (int) {

	p1 := strings.Index(str, "(")
	p2 := strings.Index(str, ")")

	if p1 < 0 || p2 < 0 {
		return (0)
	}

	if !IsLetter(str[0:p1]) {
		return (0)
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

func GetAllIma(str string) (map[int]string) {

	data := make(map[int]string)
	var itab, neg int

	if str[0] == '+' {
		str = str[1:len(str)]
	}

	itab = 0
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			
			str = str[1:len(str)]
			neg = 1
		}
		index := GetCararc(str, "+-/*%")
		if index == -1 {
			data[itab] = str
			return (data)
		}
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
		str = str[i + 1:len(str)]
		i = -1
	}
	return (data)
}