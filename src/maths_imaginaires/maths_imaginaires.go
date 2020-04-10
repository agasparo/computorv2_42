package maths_imaginaires

import (
	"strings"
	"strconv"
	"fmt"
)

type TmpComp struct {

	a float64
	b float64
}

func GetAll(str string) (string) {

	final_number := TmpComp{}

	for i := 0; i < len(str); i += 5 {

		if len(str) > i + 4 {
			fmt.Println(str[i:i + 4])
			n1, n2 := ParseOne(str[i:i + 4])
			fmt.Printf("n1 : %f n2 : %f\n", n1, n2)
		} else {
			fmt.Println(str[i:len(str)])
			add := ""
			if len (str[i:len(str)]) == 1 {
				add = "+0i"
			}
			n1, n2 := ParseOne(str[i:len(str)] + add)
			fmt.Printf("n1 : %f n2 : %f\n", n1, n2)
		}

		if i != 0 {
			switch sign := string(str[i - 1]); sign {
				case "+":
					add(&TmpComp, n1, n2)
				case "-":
					sous(&TmpComp, n1, n2)
				case "*":
					mul(&TmpComp, n1, n2)
				case "/":
					divi(&TmpComp, n1, n2)
			}
		}
	}

	return ("ok")
}

func ParseOne(str string) (x float64, y float64) {

	str_tmp := strings.ReplaceAll(str, "*", "")
	str_tmp = strings.ReplaceAll(str_tmp, " ", "")
	str_tmp = strings.ReplaceAll(str_tmp, "-", "+-")
	new_str := strings.Split(str_tmp, "+")

	fmt.Println(new_str)

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

	Finu.a = ((Finu.a * a) - (Finu.b * b))
	Finu.b = ((Finu.a * b) + (a * Finu.b))
}

func divi(Finu *TmpComp, a float64, b float64) {

}

func sous(Finu *TmpComp, a float64, b float64) {

	Finu.a = Finu.a - a
	Finu.b = Finu.b - b
}