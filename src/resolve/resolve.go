package resolve

import (
	//"fmt"
	"parser"
	"maths_functions"
	"types"
	"strings"
	"strconv"
	"equations"
)

type Unknown struct {

	Tab map[int]string
	Deg_max map[int]int
	Part1 map[int]string
	Part2 map[int]string
	Eqs map[int]equations.Equation
}

type Resol struct {

	Tab map[int]string
}

func IsSoluble(U Unknown) (bool) {

	for i := 0; i < len(U.Deg_max); i++ {

		if U.Deg_max[i] > 2 || U.Deg_max[i] < 0 {
			return (false)
		}
	}
	return (true)
}

func Init(U *Unknown, Dat types.Variable) (string) {

	RempEq(U.Tab, U)
	return ("|")
}

func RempEq(tab map[int]string, U *Unknown) {

	WE := 0

	for i := 0; i < len(tab); i++ {
		tab[i] = strings.ReplaceAll(tab[i], " ", "")

		if len(U.Part1) - 1 < i {
			WE = 1
		}

		if strings.Index(tab[i], "|") != -1 {
			e := strings.Split(tab[i], "|")
			GetAllSign(e[0], e[1], U, WE)
		} else {
			RPuis(tab[i], 0, WE, U)
		}
	}
}

func GetAllSign(str string, x string, U *Unknown, WE int) {

	var puis int
	var sign int

	if str[0] == '-' || str[0] == '+' {
		if str[0] == '-' {
			sign = 1
		}
		str = str[1:len(str)]
	}
	for i := GetIndex(str); i != -1; i = GetIndex(str) {
		if str[0] == '-' || str[0] == '+' {
			if str[0] == '-' {
				sign = 1
			}
			str = str[1:len(str)]
			i = GetIndex(str)
			if i == -1 {
				break
			}
		}
		puis = GetMaxDeg(str[0:i], x)
		if puis < 0 {
			puis = 0
		}
		if sign == 1 {
			RPuis(getNumber("-" + str[0:i], x), puis, WE, U)
		} else {
			RPuis(getNumber(str[0:i], x), puis, WE, U)
		}
		str = str[i:len(str)]
		sign = 0
	}
	puis = GetMaxDeg(str, x)
	if puis < 0 {
		puis = 0
	}
	RPuis(getNumber(str, x), puis, WE, U)
}


func getNumber(str string, x string) (string) {

	str = strings.ReplaceAll(str, "*", "")
	str = strings.ReplaceAll(str, "^", "")
	str = strings.ReplaceAll(str, x, "")
	if str == "" {
		return ("1")
	}
	return (str)
}

func GetIndex(str string) (int) {

	max := -1
	a := strings.Index(str, "/")
	if a != -1 && (max == -1 || a < max) {
		max = a
	}
	a = strings.Index(str, "%")
	if a != -1 && (max == -1 || a < max) {
		max = a
	}
	a = strings.Index(str, "-")
	if a != -1 && (max == -1 || a < max) {
		max = a
	}
	a = strings.Index(str, "+")
	if a != -1 && (max == -1 || a < max) {
		max = a
	}
	return (max)
}

func RPuis(nb string, puis int, wEqs int, U *Unknown) {

	var Eq equations.Equation

	if len(U.Eqs) == 0 {
		U.Eqs = make(map[int]equations.Equation)
	}

	if val, ok := U.Eqs[wEqs]; ok {
		Eq = val
	} else {
		Eq = equations.Equation{}
	}
	
	t1, _ := strconv.ParseFloat(nb, 64)
	
	if puis == 0 {
		Eq.C += t1
	} else if puis == 1 {
		Eq.B += t1
	} else if puis == 2 {
		Eq.A += t1
	}

	U.Eqs[wEqs] = Eq
}

func IsEquation(U *Unknown, Dat types.Variable, t int) (bool) {
	
	var tab map[int]string

	if t == 0 {
		tab = U.Part1
	} else {
		tab = U.Part2
	}

	if len(U.Tab) == 0 {
		U.Tab = make(map[int]string)
		U.Deg_max = make(map[int]int)
	}
	f := 0

	for i := 0; i < len(tab); i += 2 {

		if parser.IsFunc(tab[i], 0) == 1 {
			x := maths_functions.Getx(tab[i])
			p1 := strings.Index(tab[i], "(")
			p2 := strings.Index(tab[i], ")")
			r := tab[i][p1 + 1:p2]
			if parser.IsExpression(x, r) {
				return (false)
			}
			_, val := parser.GetDataFunc(tab[i], Dat.Table)
			if val == "" {
				return (false)
			}
			U.Tab[len(U.Tab)] = val + "|" + x
			U.Deg_max[len(U.Deg_max)] = GetMaxDeg(val, x)
			f++
		} else {
			U.Tab[len(U.Tab)] = tab[i]
			U.Deg_max[len(U.Deg_max)] = 0
			f++
		}
	}
	if f == 0 {
		return (false)
	}
	return (true)
}

func GetMaxDeg(str string, x string) (int) {

	var z, i int
	str = strings.ReplaceAll(str, " ", "")
	a := strings.Index(str, x)
	if a == -1 {
		return (-3)
	}
	max := 1
	str = strings.ReplaceAll(str, "ˆ", "^")
	for a = a; a != -1; a = strings.Index(str, x) {

		if a + 1 >= len(str) {
			return (max)
		}
		if str[a + 1] == '*' {
			z, i = GetDeg('*', str, a + 1, x)
			if z > max {
				max = z
			}
		} else if string(str[a + 1]) == "^" {
			z, _ = strconv.Atoi(string(str[a + 2]))
			i = a + 3
			if z > max {
				max = z
			}
		} else if a - 1 >= 0 && string(str[a - 1]) == "^" {
			z, _ = strconv.Atoi(string(str[a - 2]))
			i = a
			if z > max {
				max = z
			}
		}
		if i + 1 >= len(str) {
			return (max)
		}
		str = str[i + 1:len(str)]
	}
	return (max)
}

func GetDeg(sign byte, str string, deb int, x string) (int, int) {

	puis := 1
	i := deb

	for i = i; i < len(str) && str[i] == sign; i += 2 {

		if string(str[i + 1]) == x {
			puis++
		}
	}
	return puis, i
}