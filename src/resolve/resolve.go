package resolve

import (
	"fmt"
	"parser"
	"maths_functions"
	"types"
	"strings"
	"strconv"
)

type Equation struct {

	A float64
	B float64
	C float64
}

type Unknown struct {

	Tab map[int]string
	Deg_max map[int]int
}

func IsEquation(data map[int]string, tab map[int]string, U *Unknown, Dat types.Variable) (bool) {
	
	U.Tab = make(map[int]string)
	U.Deg_max = make(map[int]int)

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
			U.Tab[len(U.Tab)] = x
			U.Deg_max[len(U.Deg_max)] = GetMaxDeg(val, x)
		}
	}
	fmt.Println(U)
	return (true)
}

func GetMaxDeg(str string, x string) (int) {

	var z, i int
	str = strings.ReplaceAll(str, " ", "")
	a := strings.Index(str, x)
	if a == -1 {
		return (-1)
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

	for i = i; str[i] == sign; i += 2 {

		if string(str[i - 1]) == x {
			puis++
		}
	}
	return puis, i
}

func IsSoluble() {

}

func Init(data map[int]string) {

}