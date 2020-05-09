package maths_functions

import (
	"strings"
	"types"
	"replace_vars"
	"maths_imaginaires"
	"parser"
	"maps"
	"matrices"
	"norm"
)

func Init(tab map[int]string, x string, vars *types.Variable, Dat types.Variable) (string) {

	x = Getx(x)
	tab = ReplaceU(tab, x, Dat)
	if maps.Array_search_count(tab, "(") >= 1 {
		index := maps.Array_search_last(tab, ")")
		res, ne := ReplaceX(tab, index, x, vars, Dat)
		if res != "ok" {
			return (res)
		}
		tab = ne
		return (JoinTab(tab))
	}
	for i := 0; i < len(tab); i++ {

		if tab[i] != x && strings.Index(tab[i], x) != -1 {
			AddMul(tab[i], x, tab, i)
		}

		if tab[i] != x {
			tab[i] = replace_vars.GetVars(vars, tab[i])
		}

		if strings.Index(tab[i], "]") != -1 || strings.Index(tab[i], "[") != -1 {
			tab = matrices.Parse(tab, Dat, vars)
			if strings.Index(tab[0], "You") != -1 {
				matrices.RemoveTmp(Dat)
				return (tab[0])
			}
			if !norm.Normalize(vars) {
				tab[0] = "You have a mistake in your matrice"
				matrices.RemoveTmp(Dat)
				return (tab[0])
			}
		}
	}
	tab = maths_imaginaires.CalcMulDivi(tab, vars, x)
	tab = maths_imaginaires.CalcAddSous(tab, vars, x)
	for i := 0; i < len(tab); i++ {
		tab[i] = ReplaceMat(tab[i], vars)
	}
	return (JoinTab(tab))
}

func ReplaceU(tab map[int]string, x string, Dat types.Variable) (map[int]string) {

	for i := 0; i < len(tab); i++ {

		nstr := tab[i]
		nstr = strings.ReplaceAll(nstr, ")", "")
		nstr = strings.ReplaceAll(nstr, "(", "")
		if nstr != x && !parser.IsNumeric(nstr) && !parser.Is_defined(nstr, Dat) && strings.Index(nstr, x) == -1 && !IsSign(nstr) {
			if strings.Index(nstr, "]") == -1 && strings.Index(nstr, "[") == -1 && strings.Index(nstr, "mat") == -1 && !maths_imaginaires.IsMat(nstr, &Dat) {
				tab[i] = strings.ReplaceAll(tab[i], nstr, x)
			}
		}
	}
	return (tab)
}

func ReplaceX(tab map[int]string, min int, x string, vars *types.Variable, Dat types.Variable) (string, map[int]string) {
	
	for i := len(tab) - 1; i > min; i-- {

		if tab[i] != x && strings.Index(tab[i], x) != -1 {
			AddMul(tab[i], x, tab, i)
		}

		if tab[i] != x {
			tab[i] = replace_vars.GetVars(vars, tab[i])
		}

		if strings.Index(tab[i], "]") != -1 || strings.Index(tab[i], "[") != -1 {
			tab = matrices.Parse(tab, Dat, vars)
			if strings.Index(tab[0], "You") != -1 {
				matrices.RemoveTmp(Dat)
				return tab[0], tab
			}
			if !norm.Normalize(vars) {
				tab[0] = "You have a mistake in your matrice"
				matrices.RemoveTmp(Dat)
				return tab[0], tab
			}
		}
	}
	ntab := maps.Copy(tab)
	ntab = maps.Cut(ntab, min + 2, len(tab))
	ntab = maths_imaginaires.CalcMulDivi(ntab, vars, x)
	ntab = maths_imaginaires.CalcAddSous(ntab, vars, x)
	atab := maps.Copy(tab)
	atab = maps.Cut(atab, 0, min + 2)
	tab = maps.Combine(ntab, atab, 0, len(ntab))
	tab = maps.Clean(tab)
	for i := len(tab) - 1; i > min; i-- {
		tab[i] = ReplaceMat(tab[i], vars)
	}
	return "ok", tab
}

func ReplaceMat(str string, vars *types.Variable) (string) {

	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	if val, ok := vars.Table[strings.ToLower(str)]; ok {
		return (val.Value())
    }
    return (str)
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

func SliceTab(tab map[int]string, a int, b int) (map[int]string) {

	data := make(map[int]string)
	c := 0

	for i := a; i < b; i++ {

		data[c] = tab[i]
		c++
	}
	return (data)
}

func RempTab(tab map[int]string, data map[int]string, a int) (map[int]string) {

	for i := 0; i < len(data); i++ {

		tab[a + i] = data[i]
	}
	return (tab)
}

func AddMul(str string, x string, tab map[int]string, i int) {

	if  strings.Index(str, "ˆ") == -1 && strings.Index(str, "^") == -1 {
		nstr := strings.Split(str, x)
		Slice1 := SliceTab(tab, i + 1, len(tab))
		if nstr[0] != "" || nstr[1] != "" {
			if nstr[0] == "" {
				if nstr[1] == "-" || nstr[1] == "+" {
					nstr[1] += "1"
				}
				tab[i + 0] = x
				tab[i + 1] = "*"
				tab[i + 2] = nstr[1]
			} else {
				if nstr[0] == "-" || nstr[0] == "+" {
					nstr[0] += "1"
				}
				tab[i + 0] = nstr[0]
				tab[i + 1] = "*"
				tab[i + 2] = x
			}
			tab = RempTab(tab, Slice1, i + 3)
		}
	} else{
		nstr := strings.Split(str, x)
		if nstr[0] != "" && nstr[1] != "" {
			Slice1 := SliceTab(tab, i + 1, len(tab))
			if strings.Index(nstr[0], "ˆ") != -1 || strings.Index(nstr[0], "^") != -1 {
				tab[i + 0] = nstr[0] + x
				tab[i + 1] = "*"
				tab[i + 2] = nstr[1]
			}
			if strings.Index(nstr[1], "ˆ") != -1 || strings.Index(nstr[1], "^") != -1 {
				tab[i + 0] = nstr[0]
				tab[i + 1] = "*"
				tab[i + 2] = x + nstr[1]
			}
			tab = RempTab(tab, Slice1, i + 3)
		}
	}
}

func PuiSign(data map[int]string) (map[int]string) {

	for i := 0; i < len(data); i++ {

		if (strings.Index(data[i], "ˆ") != -1 || strings.Index(data[i], "^") != -1) && i - 1 >= 0 && (data[i - 1] == "-" || data[i - 1] == "+") {
			data[i - 2] = data[i - 2] + data[i - 1] + data[i]
			data = maps.MapSlice(data, i - 1)
		}
	}
	return (data)
}

func Calc(fu string, x string, r string, vars *types.Variable) (float64, float64) {

	parser_err := 0
	fu = strings.ReplaceAll(fu, x, r)
	data := parser.GetAllIma(fu, &parser_err)
	data = PuiSign(data)
	data = maths_imaginaires.CalcMulDivi(data, vars, x)
	data = maths_imaginaires.CalcAddSous(data, vars, x)
	return (maths_imaginaires.ParseOne(data[0], vars))
}

func IsSign(str string) (bool) {

	if strings.Index(str, "+") != -1 || strings.Index(str, "-") != -1 || strings.Index(str, "/") != -1 || strings.Index(str, "*") != -1 || strings.Index(str, "%") != -1 {
		return (true)
	}
	return (false)
}