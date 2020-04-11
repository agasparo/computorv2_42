package maths_imaginaires

import (
	"parser"
	"strings"
	"strconv"
	"fmt"
)

type TmpComp struct {

	a float64
	b float64
}

func GetAll(str string) {

	data := make(map[int]string)
	var index, itab, neg int
	var tmp_str string

	if str[0] == '+' {
		str = str[1:len(str)]
	}

	itab = 0
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			
			str = str[1:len(str)]
			neg = 1
		}

		index = parser.GetCararc(str, "+-/*")
		tmp_str = str[i:index]
		if neg == 1 {
			data[itab] = ("-" + tmp_str)
		} else {
			data[itab] = tmp_str
		}
		sign, add := parser.GetSign(str, index)
		data[itab] += sign
		i = index + add
		index = parser.GetCararc(str[i:len(str)], "+-/*")
		if index == -1 {
			tmp_str = str[i:len(str)]
		} else {
			tmp_str = str[i:index]
		}
		data[itab] += tmp_str
		itab++
		fmt.Println(data)
		return
	}
}

func ParseOne(str string) (x float64, y float64) {

	str_tmp := strings.ReplaceAll(str, "*", "")
	new_str := strings.Split(str_tmp, "+")

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

func mul(Finu *TmpComp, a float64, b float64) { // a finir

	Finu.a = ((Finu.a * a) - (Finu.b * b))
	Finu.b = ((Finu.a * b) + (a * Finu.b))
}

func divi(Finu *TmpComp, a float64, b float64) { // a faire

}

func sous(Finu *TmpComp, a float64, b float64) {

	Finu.a = Finu.a - a
	Finu.b = Finu.b - b
}