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
	for i := 0; i < len(tab); i++ {

		tab[i] = replace_vars.GetVars(vars, tab[i])
	}
	fmt.Println(tab)
	x = Getx(x)
	fmt.Println(x)
	tab = maths_imaginaires.CalcMulDivi(tab, vars, x)
	tab = maths_imaginaires.CalcAddSous(tab, vars, x)
	fmt.Println(tab)
	return ("ok")
}

func Getx(str string) (string) {

	p1 := strings.Index(str, "(")
	str = str[p1 + 1:len(str)]
	str = strings.ReplaceAll(str, ")", "")
	return (str)
}