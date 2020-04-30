package resolve

import (
	"fmt"
	"parser"
	"maths_functions"
	"types"
	"strings"
	"strconv"
	"equations"
	"maps"
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

	res, t := getSignEq(U)
	if res == "1" {
		if t == 2 {
			ModMAps(U)
			RempEq(U.Tab, U, Dat)
			return ("|")
		}
		if t == 1 {
			//RempEq(U.Tab[0], U)
			return ("|")
		}
		RempEq(U.Tab, U, Dat)
		return ("|")
	}
	return (res)
}

func ModMAps(U *Unknown) {

	A := Unknown{}

	tmp := make(map[int]string)
	tmp1 := make(map[int]string)

	tmp[0] = U.Part1[0]
	A.Part1 = tmp

	tmp1[0] = U.Part2[0]
	A.Part2 = tmp1

	A.Tab = make(map[int]string)
	A.Tab[0] = U.Tab[0]
	A.Tab[1] = "0"

	A.Deg_max = make(map[int]int)
	A.Deg_max[0] = U.Deg_max[0]
	A.Deg_max[1] = 0

	U = &A
}

func getSignEq(U *Unknown) (string, int) {

	if len(U.Part1) > 2 {
		
		sign := InitSign(U.Part1[1])
		if sign == "%" {
			return "Sorry i can't resolve this equation", 0
		}
		for i := 3; i < len(U.Part1); i += 2 {

			fmt.Println(U.Part1[i])
			if strings.Index(sign, U.Part1[i]) == -1 {
				return "Sorry i can't resolve this equation", 0
			}
		}
		if len(sign) == 1 && (len(U.Part2) > 2 || (U.Part2[0] != "0" && U.Part2[0] != "-0"))  {
			return "Sorry i can't resolve this equation", 0
		}
		if sign == "*" {
			return "1", 1
		}
		if sign == "/" {
			return "1", 2
		}
	}
	for a := 0; a < len(U.Tab); a++ {
		if strings.Index(U.Tab[a], "(") != -1 {
			return "Sorry i can't resolve this equation", 0
		}
	}
	return "1", 3
}

func InitSign(str string) (string) {

	if str == "+" || str == "-" {
		return ("+-")
	}
	return (str)
}

func RempEq(tab map[int]string, U *Unknown, Dat types.Variable) {

	WE := 0
	pos_s := -3
	var sign string

	for i := 0; i < len(tab); i++ {
		tab[i] = strings.ReplaceAll(tab[i], " ", "")

		if (len(U.Part1) - 1) / 2 < i {
			WE = 1
			pos_s = -3
		}

		pos_s += 2
		if WE == 1 && i - 1 >= 0 {
			fmt.Printf("2\n")
			sign = U.Part2[pos_s]
		} else if WE == 0 && i - 1 >= 0 {
			fmt.Printf("1\n")
			sign = U.Part1[pos_s]
		}

		if strings.Index(tab[i], "|") != -1 {
			e := strings.Split(tab[i], "|")
			ck := 0
			GetAllSign(maps.Join(parser.Checkfunc(parser.GetAllIma(e[0], &ck), Dat), ""), e[1], U, WE, sign)
		} else {
			RPuis(tab[i], 0, WE, U, sign)
		}
	}
}

func GetAllSign(str string, x string, U *Unknown, WE int, signdeb string) {

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
			RPuis(getNumber("-" + str[0:i], x), puis, WE, U, signdeb)
		} else {
			RPuis(getNumber(str[0:i], x), puis, WE, U, signdeb)
		}
		str = str[i:len(str)]
		sign = 0
	}
	puis = GetMaxDeg(str, x)
	if puis < 0 {
		puis = 0
	}
	RPuis(getNumber(str, x), puis, WE, U, signdeb)
}


func getNumber(str string, x string) (string) {

	if strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1  {
		index := strings.Index(str, x)
		str = str[0:index]
	}
	str = strings.ReplaceAll(str, "*", "")
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

func RPuis(nb string, puis int, wEqs int, U *Unknown, sign string) {

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
	
	if sign == "-" {
		if puis == 0 {
			Eq.C -= t1
		} else if puis == 1 {
			Eq.B -= t1
		} else if puis == 2 {
			Eq.A -= t1
		}
	} else if sign == "+" || sign == "" {
		if puis == 0 {
			Eq.C += t1
		} else if puis == 1 {
			Eq.B += t1
		} else if puis == 2 {
			Eq.A += t1
		}
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
			if parser.IsExpression(x, r, Dat) {
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