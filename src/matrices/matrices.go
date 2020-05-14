package matrices

import (
	"strings"
	"types"
	"strconv"
	"maps"
	"unicode"
)

func Parse(tab map[int]string, Dat types.Variable, vars *types.Variable) (map[int]string) {

	neg := 0
	is_par := ""
	for z := 0; z < len(tab); z++ {
		neg = 0
		if strings.Index(tab[z], "[") != -1 || strings.Index(tab[z], "]") != -1 {
			if Power(tab[z]) {
				tab[0] = "You are not allow to use Power with matrices"
				return (tab)
			}
			if tab[z][0] == '-' {
				neg = 1
				tab[z] = tab[z][1:len(tab[z])]
			}
			if tab[z][0] == '+' {
				tab[z] = tab[z][1:len(tab[z])]
			}
			Matr := types.Matrice{}
			tab[z] = AddMat(tab, z, &is_par)
			if !CountPara(tab[z]) {
				tab[0] = "You have a problem with your matrices syntaxe 2"
				return (tab)
			}
			if tab[z][0] != '[' || tab[z][1] != '[' {
				tab[0] = "You must have [[ at the begining of your matrice"
				return (tab)
			}
			if tab[z][len(tab[z]) - 1] != ']' || tab[z][len(tab[z]) - 2] != ']' {
				tab[0] = "You must have ]] at the end of your matrice"
				return (tab)
			}
			if strings.Count(tab[z], "[") != strings.Count(tab[z], "]") {
				tab[0] = "You must have the same number of '['' & ']'" 
				return (tab)
			}
			if strings.Count(tab[z], "[") > 2 && strings.Index(tab[z], ";") == -1 {
				tab[0] = "You have a problem with your matrices syntaxe 1"
				return (tab)
			}
			table := strings.Split(tab[z], ";")
			Matr.Mat = make(map[int]types.MatRow)
			for i := 0; i < len(table); i++ {

				part := strings.Split(table[i], ",")
				Line := types.MatRow{}
				Line.Row = make(map[int]string)
				for a := 0; a < len(part); a++ {
					part[a] = strings.ReplaceAll(part[a], "[", "")
					part[a] = strings.ReplaceAll(part[a], "]", "")
					part[a] = strings.ReplaceAll(part[a], "ˆ", "^")
					if part[a] == "" || IsSign(part[a]) || part[a][0] == '^' || part[a][len(part[a]) - 1] == '^' {
						tab[0] = "You must have a number in a matrice"
						return (tab)
					}
					if IsFunc(part[a], 0) == 1 || IsFunc(part[a], 1) == 1 {
						tab[0] = "You are not allow to use functions in matrices"
						return (tab)
					}
					if part[a][0] == '(' || part[a][len(part[a]) - 1] == ')' {
						tab[0] = "You are not allow to use expression with parentheses in matrices"
						return (tab)
					}
					part[a] = strings.ReplaceAll(part[a], ")", "")
					part[a] = strings.ReplaceAll(part[a], "(", "")
					Line.Row[len(Line.Row)] = part[a]
				}
				if !CheckLength(Matr.Mat, len(Line.Row)) {
					tab[0] = "You have a problem with your matrices syntaxe 3"
					return (tab)
				}
				Matr.Mat[len(Matr.Mat)] = Line
			}
			name := GenerateName(Dat)
			if neg == 1 {
				Tmp := make(map[int]string)
				if tab[z][len(tab[z]) -1] == ')' {
					Tmp[0] = "*"
					Tmp[1] = tab[z] + ")"
					Tmp[2] = tab[z + 1]
					tab[z] = "(-1"
				} else {
					Tmp[0] = "*"
					Tmp[1] = "-1)"
					tab[z] = "(" + tab[z]
				}
				tab = maps.CombineN(tab, Tmp, z + 1)
			}
			vars.Table[name] = &Matr
			tab[z] = CheckPara(name, tab[z], Matr)
			if strings.Index(is_par, "|") != - 1 {
				pae := strings.Split(is_par, "|")
				tab[z] = pae[0] + tab[z] + pae[1]
			}
		}
		is_par = ""
	}
	tab = maps.Reindex(tab)
	tab = maps.Clean(tab)
	return (tab)
}

func IsSign(str string) (bool) {

	if str == "+" || str == "-" || str == "/" || str == "*" || str == "%" {
		return (true)
	}
	return (false)
}

func CountPara(str string) (bool) {

	if len(str) < 2 {
		return (false)
	}

	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, "(", "")
	str = str[1:len(str) - 1]

	pv := strings.Count(str, ";")
	po := strings.Count(str, "[")
	pf := strings.Count(str, "]")

	if po != (pv + 1) || pf != (pv + 1){
		return (false)
	}
	return (true)
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

func Power(str string) (bool) {

	str = strings.ReplaceAll(str, "ˆ", "^")
	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, "(", "")
	for e := strings.Index(str, "^"); e != -1; e = strings.Index(str, "^") {
		if e - 1 >= 0 && (str[e - 1] == '[' || str[e - 1] == ']') {
			return (true)
		}
		if e + 1 < len(str) && (str[e + 1] == '[' || str[e + 1] == ']') {
			return (true)
		}
		str = str[e + 1:len(str)]
	}
	return (false)
}

func CheckPara(n string, str string, Mat types.Matrice) (string) {

	return (strings.ReplaceAll(str, Mat.Value(), n))
}

func Modifi(m string) (types.Matrice) {

	M := types.Matrice{}

	m = strings.ReplaceAll(m, "[", "")
	m = strings.ReplaceAll(m, "]", "")
	e := strings.Split(m, ";")
	M.Mat = make(map[int]types.MatRow)
	for i := 0; i < len(e); i++ {
		Line := types.MatRow{}
		Line.Row = make(map[int]string)
		ex := strings.Split(e[i], ",")
		for z := 0; z < len(ex); z++ {
			Line.Row[len(Line.Row)] = ex[z]
		}
		M.Mat[len(M.Mat)] = Line
	}
	return (M)
}

func AddMat(tab map[int]string, z int, is_par *string) (string) {
	
	if strings.Count(tab[z], "[") == strings.Count(tab[z], "]") {
		index := strings.Index(tab[z], "[")
		index_fin := IndexString(tab[z], "]")
		tab[z] = ReplacePara(index, index_fin, tab[z], is_par)
		return (tab[z])
	}

	save := -1
	for i := z + 1; i < len(tab); i++ {
		if strings.Index(tab[i], "[") != -1 && save != -1 {
			break
		}

		if strings.Index(tab[i], "]") != -1 {
			if strings.Count(tab[i], "]") >= 2 {
				save = i
			}
		}
	}
	if save == -1 {
		return (tab[z])
	}
	tab[z] = maps.Add(tab, tab[z], z + 1, save + 1)
	tab = maps.MapSliceCount(tab, z + 1, save - z)
	tab = maps.Reindex(tab)
	tab = maps.Clean(tab)
	index := strings.Index(tab[z], "[")
	index_fin := IndexString(tab[z], "]")
	tab[z] = ReplacePara(index, index_fin, tab[z], is_par)
	return (tab[z])
}

func ReplacePara(deb int, fin int, str string, is_par *string) (string) {

	if deb > 0 {
		*is_par += str[0:deb]
	}
	*is_par += "|"
	if fin + 1 < len(str) {
		*is_par += str[fin + 1:len(str)]
	}
	return (str[deb:fin + 1])
}

func IndexString(str string, s string) (int) {

	pos := -1

	for i := 0; i < len(str); i++ {
		
		if string(str[i]) == s {
			pos = i
		}
	}
	return (pos)
}

func CheckLength(tab map[int]types.MatRow, cmp int) (bool) {

	for i := 0; i < len(tab); i++ {

		if len(tab[i].Row) != cmp {
			return (false)
		}
	}
	return (true)
}

func GenerateName(Dat types.Variable) (string) {

	i := 0
	name := "mat" + strconv.Itoa(i)
	for i := 1; IsDefined(name, Dat); i++ {
		name = "mat" + strconv.Itoa(i)
	}
	return (name)
}

func IsDefined(str string, vars types.Variable) (bool) {

	if _, ok := vars.Table[strings.ToLower(str)]; ok {

		return (true)
    }
   	return (false)
}

func RemoveTmp(vars types.Variable) {

	i := 0
	name := "mat" + strconv.Itoa(i)
	for i := 1; IsDefined(name, vars); i++ {
		delete(vars.Table, name)
		name = "mat" + strconv.Itoa(i)
	}
}

func GetnbLine(ma string) (int) {

	c := 0

	m := strings.Split(ma, ";")
	for i := 0; i < len(m); i++ {
		c++
	}
	return (c)
}

func GetnbCol(ma string) (int) {

	c := 0

	m := strings.Split(ma, ";")
	for i := 0; i < len(m); i++ {
		ms := strings.Split(m[i], ",")
		for z := 0; z < len(ms); z++ {
			c++
		}
		return (c)
	}
	return (c)
}