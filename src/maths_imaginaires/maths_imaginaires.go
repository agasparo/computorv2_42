package maths_imaginaires

import (
	"strings"
	"strconv"
	"maps"
	"fmt"
)

type TmpComp struct {

	a float64
	b float64
}

func CalcVar(data map[int]string) (float64, float64) {

	fmt.Println(data)
	data = CalcMulDivi(data)
	data = CalcAddSous(data)
	return ParseOne(data[0])
}

func CalcMulDivi(data map[int]string) (map[int]string) {

	for i := 1; i < len(data); i += 2 {

		if data[i] == "*" {
			nb1, nb2 := ParseOne(data[i - 1])
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1])
			mul(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "/" {
			nb1, nb2 := ParseOne(data[i - 1])
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1])
			divi(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}
	}
	return (data)
}

func CalcAddSous(data map[int]string) (map[int]string) {

	for i := 1; i < len(data); i += 2 {

		if data[i] == "+" {
			nb1, nb2 := ParseOne(data[i - 1])
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1])
			add(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "-" {
			nb1, nb2 := ParseOne(data[i - 1])
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1])
			sous(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}
	}
	return (data)
}

func Float2string(Calc TmpComp) (string) {

	if Calc.b >= 0 {
		return (fmt.Sprintf("%f+%fi", Calc.a, Calc.b))
	}
	return (fmt.Sprintf("%f%fi", Calc.a, Calc.b))
}

func ParseOne(str string) (x float64, y float64) {

	var neg int = 0

	if str[0] == '-' {
		neg = 1
		str = str[1:len(str)]
	}
	str_tmp := strings.ReplaceAll(str, "*", "")
	str_tmp = strings.ReplaceAll(str_tmp, "-", "+-")
	new_str := strings.Split(str_tmp, "+")

	if neg == 1 {
		new_str[0] = "-" + new_str[0]
	}

	if strings.Index(new_str[0], "i") != -1 {
		y, _ = strconv.ParseFloat(strings.ReplaceAll(new_str[0], "i", ""), 64)
		x, _ = strconv.ParseFloat(new_str[len(new_str) - 1], 64)
	} else {
		x, _ = strconv.ParseFloat(new_str[0], 64)
		y, _ = strconv.ParseFloat(strings.ReplaceAll(new_str[len(new_str) - 1], "i", ""), 64)
	}

	return x, y
}

/************************************************************************************************/

func add(Finu *TmpComp, a float64, b float64) {

	Finu.a = Finu.a + a
	Finu.b = Finu.b + b
}

func mul(Finu *TmpComp, a float64, b float64) {

	tmp := ((Finu.a * a) - (Finu.b * b))
	Finu.b = ((Finu.a * b) + (a * Finu.b))
	Finu.a = tmp
}

func divi(Finu *TmpComp, a float64, b float64) { // a faire

}

func sous(Finu *TmpComp, a float64, b float64) {

	Finu.a = Finu.a - a
	Finu.b = Finu.b - b
}