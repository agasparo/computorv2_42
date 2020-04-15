package maths_imaginaires

import (
	"strings"
	"strconv"
	"maps"
	"fmt"
	"regexp"
	"types"
	"replace_vars"
)

type TmpComp struct {

	a float64
	b float64
}

func CalcVar(data map[int]string, vars *types.Variable) (float64, float64) {

	data = CalcMulDivi(data, vars, "")
	data = CalcAddSous(data, vars, "")
	return ParseOne(data[0], vars)
}

func CalcMulDivi(data map[int]string, vars *types.Variable, inconnue string) (map[int]string) {

	for i := 1; i < len(data); i += 2 {

		if data[i] == "*" && data[i - 1] != inconnue && data[i + 1] != inconnue {
			nb1, nb2 := ParseOne(data[i - 1], vars)
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1], vars)
			mul(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "%" && data[i - 1] != inconnue && data[i + 1] != inconnue {
			nb1, nb2 := ParseOne(data[i - 1], vars)
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1], vars)
			mod(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "/" && data[i - 1] != inconnue && data[i + 1] != inconnue {
			nb1, nb2 := ParseOne(data[i - 1], vars)
			Calc := TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1], vars)
			divi(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}
	}
	return (data)
}

func CalcAddSous(data map[int]string, vars *types.Variable, inconnue string) (map[int]string) {

	var Calc TmpComp
	var nb1, nb2, nb3, nb4 float64

	for i := 1; i < len(data); i += 2 {

		if data[i] == "+" && data[i - 1] != inconnue && data[i + 1] != inconnue && data[i + 2] != "*" && data[i - 2] != "*" && data[i + 2] != "/" && data[i - 2] != "/" {
			nb_puis := NegPui(data[i - 1], data[i + 1])
			if nb_puis == data[i - 1] {
				nb1, nb2 = ParseOne(data[i - 1], vars)
				Calc = TmpComp{nb1, nb2}
				nb3, nb4 = ParseOne(data[i + 1], vars)
			} else {
				nb1, nb2 = ParseOne(nb_puis, vars)
				Calc = TmpComp{nb1, nb2}
				nb3 = 0
				nb4 = 0
			}
			add(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "-" && data[i - 1] != inconnue && data[i + 1] != inconnue && data[i + 2] != "*" && data[i - 2] != "*" && data[i + 2] != "/" && data[i - 2] != "/" {
			nb_puis := NegPui(data[i - 1], data[i + 1])
			if nb_puis == data[i - 1] {
				nb1, nb2 = ParseOne(data[i - 1], vars)
				Calc = TmpComp{nb1, nb2}
				nb3, nb4 = ParseOne(data[i + 1], vars)
				sous(&Calc, nb3, nb4)
			} else {
				nb1, nb2 = ParseOne(nb_puis, vars)
				Calc = TmpComp{1, 0}
				divi(&Calc, nb1, nb2)
			}
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}
	}
	return (data)
}

func NegPui(str string, m string) (string) {

	if  str[len(str) - 1] == 134 || string(str[len(str) - 1]) == "^" {
		return (str + m)
	}
	return (str)
}

func Float2string(Calc TmpComp) (string) {

	if Calc.b == 0 {
		return (fmt.Sprintf("%f", Calc.a))
	} else if Calc.a == 0 {
		return (fmt.Sprintf("%fi", Calc.b))
	} else if Calc.b > 0 {
		return (fmt.Sprintf("%f + %fi", Calc.a, Calc.b))
	}
	return (fmt.Sprintf("%f %fi", Calc.a, Calc.b))
}

func ParseOne(str string, vars *types.Variable) (x float64, y float64) {

	if str == "i" {
		str = "1i"
	}

	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	str = replace_vars.GetVars(vars, str)
    str = strings.ReplaceAll(str, " ", "")

    r, _ := regexp.Compile(`(?m)[+-]?([0-9]*[.])?[0-9]+[-+][+-]?([0-9]*[.])?[0-9]+[i]`)

	if strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1 {
		nstr := strings.Split(str, "ˆ")
		if len(nstr) == 1 {
			nstr = strings.Split(str, "^")
		}
		a, b := TransPow(nstr)
		str = Float2string(TmpComp{ a, b })
	}

	if r.MatchString(str) {

		return Trans(str)
	}
	return TransN(str)
}

func TransPow(nstr []string) (x float64, y float64) {
	
	r, _ := regexp.Compile(`(?m)[+-]?([0-9]*[.])?[0-9]+[-+][+-]?([0-9]*[.])?[0-9]+[i]`)

	var a, c, d float64

	Base := TmpComp{}

	for i := len(nstr) - 1; i > 0; i-- {

		if nstr[i] == "i" {
			nstr[i] = "1i"
		}
		if nstr[i - 1] == "i" {
			nstr[i - 1] = "1i"
		}
		a, _ = TransN(nstr[i])
		if r.MatchString(nstr[i - 1]) {
			c, d = Trans(nstr[i - 1])
		} else {
			c, d = TransN(nstr[i - 1])
		}
		Base.a = c
		Base.b = d
		Pow(&Base, int64(a))
		nstr[i - 1] = Float2string(Base)
	}

	return Base.a, Base.b
}

func Trans(str string) (x float64, y float64) {

	neg := 0

	if str[0] == '-' {
		neg = 1
		str = str[1:len(str)]
	}
	str = strings.ReplaceAll(str, "-", "+-")
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

func TransN(str string) (x float64, y float64) {

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

func divi(Finu *TmpComp, a float64, b float64) {

	tmp := ((Finu.a * a) + (Finu.b * b)) / ((a * a) + (b * b))
	Finu.b = ((Finu.b * a) - (Finu.a * b)) / ((a * a) + (b * b))
	Finu.a = tmp
}

func sous(Finu *TmpComp, a float64, b float64) {

	Finu.a = Finu.a - a
	Finu.b = Finu.b - b
}

func mod(Finu *TmpComp, a float64, b float64) {

	Calc := TmpComp{ Finu.a, Finu.b }
	divi(&Calc, a, b)
	Calc.a = float64(int64(Calc.a))
	Calc.b = float64(int64(Calc.b))
	mul(&Calc, a, b)
	sous(Finu, Calc.a, Calc.b)

}

func Pow(n1 *TmpComp, n2 int64) {

	coe := n1.a
	im := n1.b

	if n2 == 0 {
		n1.a = 1
		n1.b = 0
		return
	}

    for i := int64(1); i < n2; i++ {
        
        mul(n1, coe, im)
        if Isinf(n1, coe, im) {
        	return
        }
    }
    Isinf(n1, coe, im)
}

func IsNan(f float64) (bool) {

	return f != f
}

func Isinf(n1 *TmpComp, coe float64, im float64) (bool) {
	
	Calc := TmpComp{n1.a, n1.b}
	mul(&Calc, coe, im)
	if IsNan(Calc.a) || IsNan(Calc.b) {
		return (true)
	}
    return (false)
}