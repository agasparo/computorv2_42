package maths_imaginaires

import (
	"strings"
	"strconv"
	//"fmt"
)

type TmpComp struct {

	a float64
	b float64
}

func GetAll(str string) (float64, float64) {

	final_number := TmpComp{}
	var n1, n2 float64
	var to = 4

	for i := 0; i < len(str); i += 5 {

		if str[i] == '-' {
			to++
		}

		to += strings.Index(strings.ReplaceAll(str, " ", ""), "+")

		if len(str) > i + to && strings.Index(str[i:i + to], "*") != -1 {
			to++
		}

		if len(str) > i + to {
			n1, n2 = ParseOne(str[i:i + to])
			//fmt.Printf("n1 : %f n2 : %f\n", n1, n2)
		} else {
			add := ""
			if len (str[i:len(str)]) <= 2 {
				add = "+0i"
			}
			n1, n2 = ParseOne(str[i:len(str)] + add)
			//fmt.Printf("n1 : %f n2 : %f\n", n1, n2)
		}

		if i != 0 {
			switch sign := string(str[i - 1]); sign {
				case "+":
					add(&final_number, n1, n2)
				case "-":
					sous(&final_number, n1, n2)
				case "*":
					mul(&final_number, n1, n2)
				case "/":
					divi(&final_number, n1, n2)
			}
		} else {
			final_number.a = n1
			final_number.b = n2
		}

		if len(str) > i + to && strings.Index(str[i:i + to], "-") != -1 {
			i++
		}
		to = 4
	}

	return final_number.a, final_number.b
}

func ParseOne(str string) (x float64, y float64) {

	str_tmp := strings.ReplaceAll(str, "*", "")
	str_tmp = strings.ReplaceAll(str_tmp, " ", "")
	//fmt.Println(str_tmp)
	if strings.Index(str_tmp, "+") == -1 {
		ntab := strings.Split(str_tmp, "-")
		if len(ntab) == 2 {
			str_tmp = strings.ReplaceAll(str_tmp, "-", "+-")
		} else {
			str_tmp = "-" + ntab[0] + ntab[1] + "+" + ntab[2]
		}
	}
	new_str := strings.Split(str_tmp, "+")

	//fmt.Println(new_str)

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