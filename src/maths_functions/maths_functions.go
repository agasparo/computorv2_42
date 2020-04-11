package maths_functions

import (
	"fmt"
	"strings"
	"types"
	"replace_vars"
	"maths_imaginaires"
)

func Init(tab map[int]string, x string, vars *types.Variable) (string) {

	fmt.Println(tab)
	x = Getx(x)
	for i := 0; i < len(tab); i++ {

		if tab[i] != x {
			tab[i] = replace_vars.GetVars(vars, tab[i])
		}
	}
	fmt.Println(tab)
	fmt.Println(x)
	tab = maths_imaginaires.CalcMulDivi(tab, vars, x)
	tab = maths_imaginaires.CalcAddSous(tab, vars, x)
	fmt.Println(CountSign(tab))
	if CountSign(tab) == 1 {
		return (JoinTab(tab))
	}
	if CheckSign(tab) == 1 {
		fmt.Println("nop")
	}
	return (JoinTab(tab))
}

func Getx(str string) (string) {

	p1 := strings.Index(str, "(")
	str = str[p1 + 1:len(str)]
	str = strings.ReplaceAll(str, ")", "")
	return (str)
}

func JoinTab(tab map[int]string) (string) {

	str := ""

	for i := 0; i < len(tab); i++ {

		str += tab[i] + " "
	}
	return (str)
}

func CountSign(tab map[int]string) (int) {

	c := 0

	for i := 1; i < len(tab); i += 2 {

		if tab[i] != "" {
			c++
		}
	}
	return (c)
}

func CheckSign(tab map[int]string) (int) { // a finir de modifier

	cmp := "a"
	type_a := "+-"
	type_b := "*/%"

	for i := 1; i < len(tab); i += 2 {

		if cmp == "a" {
			cmp = tab[i]
		} else {
			if cmp != tab[i] && tab[i] != "" {
				return (0)
			}
		}
	}
	return (1)
}