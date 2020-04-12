package parentheses

import (
	"fmt"
	"strings"
	"maths_imaginaires"
	"types"
	"maths_functions"
	"maps"
)

type TmpComp struct {

	a float64
	b float64
}

func Parse(tab map[int]string, Vars *types.Variable) (map[int]string) {

	nb_par := countPara(tab)
	if nb_par == 0 {
		return (tab)
	}
	for max := nb_par; max > 0; max-- {
		index_d := getIndexof(tab, "(", max, 0)
		index_c := getIndexof(tab, ")", max, index_d + 1)
		ntab := maths_functions.SliceTab(tab, index_d, index_c + 1)
		n1, n2 := maths_imaginaires.CalcVar(ntab, Vars)
		res := Float2string(TmpComp{ n1, n2 })
		tab[index_d] = res
		tab = maps.MapSliceCount(tab, index_d + 1, (index_d + 1) - index_c)
		tab = maps.MapReindex(tab)
	}
	return (tab)
}

func countPara(tab map[int]string) (int) {

	c := 0
	for i := 0; i < len(tab); i++ {

		if strings.Index(tab[i], "(") != -1 {
			c++
		}
	}
	return (c)
}

func getIndexof(tab map[int]string, x string, y int, z int) (int) {

	c := 0
	for i := z; i < len(tab); i++ {

		if strings.Index(tab[i], x) != -1 {
			c++
			if c == y {
				return (i)
			}
		}
	}
	return (-1)
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