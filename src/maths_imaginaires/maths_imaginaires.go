package maths_imaginaires

import (
	"strings"
	"strconv"
	"maps"
	"fmt"
	"regexp"
)

type TmpComp struct {

	a float64
	b float64
}

func CalcVar(data map[int]string) (float64, float64) {

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

	if Calc.b == 0 {
		return (fmt.Sprintf("%f", Calc.a))
	} else if Calc.a == 0 {
		return (fmt.Sprintf("%fi", Calc.b))
	} else if Calc.b > 0 {
		return (fmt.Sprintf("%f+%fi", Calc.a, Calc.b))
	}
	return (fmt.Sprintf("%f%fi", Calc.a, Calc.b))
}

func ParseOne(str string) (x float64, y float64) {

	r, _ := regexp.Compile(`(?m)[+-]?([0-9]*[.])?[0-9]+[-+][+-]?([0-9]*[.])?[0-9]+[i]`)

	if r.MatchString(str) {

		neg := 0

		if str[0] == '-' {
			neg = 1
			str = str[1:len(str)]
		}
		strings.ReplaceAll(str, "-", "+-")
		nstr := strings.Split(str, "+")
		if neg == 1 {
			nstr[0] = "-" + nstr[0]
		}

		if strings.Index(nstr[0], "i") != -1 {
			y, _ = strconv.ParseFloat(strings.ReplaceAll(nstr[0], "i", ""), 64)
			x, _ = strconv.ParseFloat(nstr[1], 64)
		} else {
			x, _ = strconv.ParseFloat(nstr[0], 64)
			y, _ = strconv.ParseFloat(strings.ReplaceAll(nstr[1], "i", ""), 64)
		}
		return x, y
	}

	if strings.Index(str, "i") != -1 {
		y, _ = strconv.ParseFloat(strings.ReplaceAll(str, "i", ""), 64)
		x, _ = strconv.ParseFloat("0.000", 64)
	} else {
		x, _ = strconv.ParseFloat(str, 64)
		y, _ = strconv.ParseFloat("0.000", 64)
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