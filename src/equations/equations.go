package equations

import (

	"fmt"
	"parser"
	"types"
	"strings"
)

type Equation struct {

	A float64
	B float64
	C float64
}

//func checkSign() (bool) {

//}

func ResolveDivi(eq string, max_deg int, vars types.Variable) {

	var data map[int]string
	var pos int

	tmp := strings.Split(eq, "|")
	eq = tmp[0]

	if max_deg == 1 || max_deg == -1 {
		data = parser.GetAllIma(eq, &pos)
		data = parser.Checkfunc(data, vars)
	}

	if max_deg == 2 || max_deg == -2 {
		data = parser.GetAllIma(eq, &pos)
		data = parser.Checkfunc(data, vars)
	}

	if max_deg == 0 {
		data = parser.GetAllIma(eq, &pos)
		data = parser.Checkfunc(data, vars)
	}

	fmt.Println(data)
}

func ResolveMulti() {

}

func ResolveNorm() {

}

func Delta() {

}

func Simple() {

}