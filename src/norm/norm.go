package norm

import (
	"strings"
	"maths_imaginaires"
	"types"
	"parser"
	"matrices"
)

func Normalize(vars *types.Variable) (bool) {

	for name, element := range vars.Table {

		if strings.Index(name, "mat") != -1 {
			nb := ReplaceVals(element.Value(), vars, name)
			if nb == -1 {
				return (false)
			}

		}
	}
	return (true)
}

func ReplaceVals(str string, vars *types.Variable, name string) (int) {

	pos := 0
	e := strings.Split(str, ";")
	for i := 0; i < len(e); i++ {
		as := strings.Split(e[i], ",")
		for a := 0; a < len(as); a++ {
			as[a] = strings.ReplaceAll(as[a], "[", "")
			as[a] = strings.ReplaceAll(as[a], "]", "")
			data := parser.GetAllIma(as[a], &pos)
			if pos == 1 {
				return -1
			}
			x, y, err := maths_imaginaires.CalcVar(data, vars)
			if err != "" {
				return -1
			}
			as[a] = maths_imaginaires.Float2string(maths_imaginaires.TmpComp{ x, y })
		}
		e[i] = strings.Join(as, ",")
	}
	res := matrices.Modifi(strings.Join(e, ";"))
	vars.Table[name] = &res
	return (0)
}