package matrices

import (
	"strings"
	//"fmt"
	"types"
	"strconv"
	"maps"
)

func Parse(tab map[int]string, Dat types.Variable, vars *types.Variable) (map[int]string) {


	for z := 0; z < len(tab); z++ {

		if strings.Index(tab[z], "[") != -1 {
			Matr := types.Matrice{}
			tab[z] = AddMat(tab, z)
			if strings.Count(tab[z], "[") != strings.Count(tab[z], "]") {
				tab[0] = "You must have the same number of '['' & ']'" 
				return (tab)
			}
			if strings.Count(tab[z], "[") > 2 && strings.Index(tab[z], ";") == -1 {
				tab[0] = "You have a problem with your matrices syntaxe"
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
					Line.Row[len(Line.Row)] = part[a]
				}
				if !CheckLength(Matr.Mat, len(Line.Row)) {
					tab[0] = "You have a problem with your matrices syntaxe"
					return (tab)
				}
				Matr.Mat[len(Matr.Mat)] = Line
			}
			name := GenerateName(Dat)
			vars.Table[name] = &Matr
			tab[z] = name
		}
	}
	tab = maps.Reindex(tab)
	tab = maps.Clean(tab)
	return (tab)
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

func AddMat(tab map[int]string, z int) (string) {
	
	if strings.Count(tab[z], "[") == strings.Count(tab[z], "]") {
		return (tab[z])
	}

	save := -1
	for i := z + 1; i < len(tab); i++ {

		if strings.Index(tab[i], "[") != -1 && save != -1 {
			break
		}

		if strings.Index(tab[i], "]") != -1 {
			save = i
		}
	}
	if save == -1 {
		return (tab[z])
	}
	tab[z] = maps.Add(tab, tab[z], z + 1, save + 1)
	tab = maps.MapSliceCount(tab, z + 1, save - z)
	tab = maps.Reindex(tab)
	tab = maps.Clean(tab)
	return (tab[0])
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