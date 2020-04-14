package parentheses

import (
	"fmt"
	"strings"
	"maths_imaginaires"
	"types"
	"maths_functions"
	"maps"
	"parser"
	"strconv"
)

type TmpComp struct {

	a float64
	b float64
}

func Parse(tab map[int]string, Vars *types.Variable, is_f bool, f_name string) (map[int]string) {

	nb_par := countPara(tab, "(")
	if nb_par == 0 {
		return (tab)
	}
	for max := nb_par; max > 0; max-- {
		index_d := getIndexof(tab, "(", max, 0)
		index_c := getIndexfin(tab, ")", index_d + 1)
		if index_c == -1 {
			index_c = index_d
		}
		ntab := maths_functions.SliceTab(tab, index_d, index_c + 1)
		if is_f {
			if maps.Array_search_count(ntab, maths_functions.Getx(f_name)) >= 1 {
				return (tab)
			}
		}
		gn, powers, pl := PowerC(ntab[0], ntab[len(ntab) - 1])
		if powers != "" {
			if pl == 0 {
				ntab[0] = gn
			} else if pl == 1 {
				ntab[len(ntab) - 1] = gn
			} else {
				nf := strings.Split(gn, "|")
				ntab[0] = nf[0]
				ntab[len(ntab) - 1] = nf[1]
			}
		}
		add, pos, repete := check(ntab)
		n1, n2 := maths_imaginaires.CalcVar(ntab, Vars)
		res := Float2string(TmpComp{ n1, n2 })
		if powers != "" {
			po := parser.GetAllIma(strings.ReplaceAll(add_check(res, powers, pl, "1"), " ", ""))
			a, b := maths_imaginaires.CalcVar(po, Vars)
			res = Float2string(TmpComp{ a, b })
		}
		tab[index_d] = add_check(res, add, pos, repete)
		tab = maps.MapSliceCount(tab, index_d + 1, index_c - index_d)
	}
	return (tab)
}

func PowerC(str string, str1 string) (string, string, int) {

	if str[0] != '(' && str1[len(str1) - 1] != ')' {
		index_d := indexString(str, "(")
		index_f := strings.Index(str1, ")")
		return (str[index_d:len(str)] + "|" + str1[0:index_f + 1]), (str[0:index_d] + "|" + str1[index_f + 1:len(str1)]), 3
	}
	if str[0] != '(' {
		index := indexString(str, "(")
		return str[index:len(str)], str[0:index], 0
	}
	if str1[len(str1) - 1] != ')' {
		index := strings.Index(str1, ")")
		return str1[0:index + 1], str1[index + 1:len(str1)], 1
	}
	return str, "", 0
}

func add_check(str string, add string, pos int, r string) (string) {

	if add == "" {
		return (str)
	}
	if pos == 3 {
		return (strings.ReplaceAll(add, "|", str))
	}
	if pos == 2 {
		nt := strings.Split(r, "|")
		np := strings.Split(add, "|")
		a, _ := strconv.Atoi(nt[0])
		b, _ := strconv.Atoi(nt[1])
		return (strings.Repeat(np[0], a) + str + strings.Repeat(np[1], b))
	}
	if pos == 0 {
		a, _ := strconv.Atoi(r)
		return (strings.Repeat(add, a) + str)
	}
	if pos == 1 {
		a, _ := strconv.Atoi(r)
		return (str + strings.Repeat(add, a))
	}
	return (str)
}

func check(tab map[int]string) (string, int, string) {

	c_d := countPara(tab, "(")
	c_f := countPara(tab, ")")

	if c_f > 1 && c_d > 1 {
		return "(|)", 2, fmt.Sprintf("%d|%d", c_d - 1, c_f - 1)
	}
	if c_d > 1 {
		return "(", 0, fmt.Sprintf("%d", c_d - 1)
	}
	if c_f > 1 {
		return ")", 1, fmt.Sprintf("%d", c_f - 1)
	}
	return "", 0,  fmt.Sprintf("%d", 0)
}

func indexString(str string, s string) (int) {

	pos := -1

	for i := 0; i < len(str); i++ {
		
		if string(str[i]) == s {
			pos = i
		}
	}
	return (pos)
}

func countPara(tab map[int]string, s string) (int) {

	c := 0
	for i := 0; i < len(tab); i++ {

		if strings.Index(tab[i], s) != -1 {
			c += strings.Count(tab[i], s)
		}
	}
	return (c)
}

func getIndexof(tab map[int]string, x string, y int, z int) (int) {

	c := 0

	for i := z; i < len(tab); i++ {

		if strings.Index(tab[i], x) != -1 {
			c += strings.Count(tab[i], x)
			if c == y {
				return (i)
			}
		}
	}
	return (-1)
}

func getIndexfin(tab map[int]string, x string, z int) (int) {

	for i := z; i < len(tab); i++ {

		if strings.Index(tab[i], x) != -1 {
			return (i)
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