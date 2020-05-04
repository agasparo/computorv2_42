package maths_imaginaires

import (
	"strings"
	"strconv"
	"maps"
	"fmt"
	"regexp"
	"types"
	"replace_vars"
	"matrices"
)

type TmpComp struct {

	A float64
	B float64
}

func CalcVar(data map[int]string, vars *types.Variable) (float64, float64, string) {

	data = CalcMulDivi(data, vars, "")
	data = CalcAddSous(data, vars, "")
	if strings.Index(data[0], "by 0") != -1 || strings.Index(data[0], "matrice") != -1 {
		return 0, 0, data[0]
	}
	a, b := ParseOne(data[0], vars)
	return a, b, ""
}

func CalcMulDivi(data map[int]string, vars *types.Variable, inconnue string) (map[int]string) {

	var Calc TmpComp

	for i := 1; i < len(data); i += 2 {

		if data[i] == "*" && data[i - 1] != inconnue && data[i + 1] != inconnue && !IsPowFunc(inconnue, data[i - 1], data[i + 1]) {
			if strings.Index(data[i - 1], "mat") != -1 || strings.Index(data[i + 1], "mat") != -1 {
				if strings.Index(data[i - 1], "mat") != -1 && strings.Index(data[i + 1], "mat") != -1 {
					if !IsCarre(data[i - 1], data[i + 1]) {
						data[0] = "Your matrice must be carre"
						return (data)
					}
					Calc = TmpComp{0, 0}
					Matrices(&Calc, data[i - 1], data[i + 1], "*", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				}
				if strings.Index(data[i - 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i + 1], vars)
					Calc = TmpComp{nb1, nb2}
					Matrices(&Calc, data[i - 1], "", "*", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				}
				if strings.Index(data[i + 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					Matrices(&Calc, data[i + 1], "", "*", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i + 1]
				}
			} else {
				nb1, nb2 := ParseOne(data[i - 1], vars)
				Calc = TmpComp{nb1, nb2}
				nb3, nb4 := ParseOne(data[i + 1], vars)
				Mul(&Calc, nb3, nb4)
				data = maps.MapSlice(data, i)
				data[i - 1] = Float2string(Calc)
			}
			i = -1
		}

		if data[i] == "%" && data[i - 1] != inconnue && data[i + 1] != inconnue && !IsPowFunc(inconnue, data[i - 1], data[i + 1]) {
			nb1, nb2 := ParseOne(data[i - 1], vars)
			Calc = TmpComp{nb1, nb2}
			nb3, nb4 := ParseOne(data[i + 1], vars)
			if nb3 == 0 {
				data[0] = "Can't do modulo by 0"
				return (data)
			}
			Mod(&Calc, nb3, nb4)
			data = maps.MapSlice(data, i)
			data[i - 1] = Float2string(Calc)
			i = -1
		}

		if data[i] == "/" && data[i - 1] != inconnue && data[i + 1] != inconnue && !IsPowFunc(inconnue, data[i - 1], data[i + 1]) {
			
			if strings.Index(data[i - 1], "mat") != -1 || strings.Index(data[i + 1], "mat") != -1 {
				if strings.Index(data[i - 1], "mat") != -1 && strings.Index(data[i + 1], "mat") != -1 {
					if !IsCarre(data[i - 1], data[i + 1]) {
						data[0] = "Your matrice must be carre"
						return (data)
					}
					Calc = TmpComp{0, 0}
					Matrices(&Calc, data[i - 1], data[i + 1], "/", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				}
				if strings.Index(data[i - 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i + 1], vars)
					nb3, _ := ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					if nb3 == 0 {
						data[0] = "Can't do division by 0"
						return (data)
					}
					Matrices(&Calc, data[i - 1], "", "/", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				}
				if strings.Index(data[i + 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					Matrices(&Calc, data[i + 1], "", "/", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i + 1]
				}
			} else {
				nb1, nb2 := ParseOne(data[i - 1], vars)
				Calc = TmpComp{nb1, nb2}
				nb3, nb4 := ParseOne(data[i + 1], vars)
				if nb3 == 0 {
					data[0] = "Can't do division by 0"
					return (data)
				}
				Divi(&Calc, nb3, nb4)
				data = maps.MapSlice(data, i)
				data[i - 1] = Float2string(Calc)
			}
			i = -1
		}
	}
	return (data)
}

func CalcAddSous(data map[int]string, vars *types.Variable, inconnue string) (map[int]string) {

	var Calc TmpComp
	var nb1, nb2, nb3, nb4 float64

	for i := 1; i < len(data); i += 2 {

		if data[i] == "+" && data[i - 1] != inconnue && data[i + 1] != inconnue && data[i + 2] != "*" && data[i - 2] != "*" && data[i + 2] != "/" && data[i - 2] != "/" && !IsPowFunc(inconnue, data[i - 1], data[i + 1]) {
			
			if strings.Index(data[i - 1], "mat") != -1 && strings.Index(data[i + 1], "mat") != -1 {
				Calc = TmpComp{0, 0}
				Matrices(&Calc, data[i - 1], data[i + 1], "+", vars)
				data = maps.MapSlice(data, i)
				data[i - 1] = data[i - 1]
			} else {
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
				Add(&Calc, nb3, nb4)
				data = maps.MapSlice(data, i)
				data[i - 1] = Float2string(Calc)
			}
			i = -1
		}

		if data[i] == "-" && data[i - 1] != inconnue && data[i + 1] != inconnue && data[i + 2] != "*" && data[i - 2] != "*" && data[i + 2] != "/" && data[i - 2] != "/" && !IsPowFunc(inconnue, data[i - 1], data[i + 1]) {
			
			if strings.Index(data[i - 1], "mat") != -1 && strings.Index(data[i + 1], "mat") != -1 {
				Calc = TmpComp{0, 0}
				Matrices(&Calc, data[i - 1], data[i + 1], "-", vars)
				data = maps.MapSlice(data, i)
				data[i - 1] = data[i - 1]
			} else {
				nb_puis := NegPui(data[i - 1], data[i + 1])
				if nb_puis == data[i - 1] {
					nb1, nb2 = ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					nb3, nb4 = ParseOne(data[i + 1], vars)
					Sous(&Calc, nb3, nb4)
				} else {
					nb1, nb2 = ParseOne(nb_puis, vars)
					Calc = TmpComp{1, 0}
					Divi(&Calc, nb1, nb2)
				}
				data = maps.MapSlice(data, i)
				data[i - 1] = Float2string(Calc)
			}
			i = -1
		}
	}
	return (data)
}

func IsPowFunc(inconnue string, str string, str1 string) (bool) {

	if inconnue == "" {
		return (false)
	}

	if strings.Index(str, inconnue) != -1 && (strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1 ) {
		return (true)
	}
	if strings.Index(str1, inconnue) != -1 && (strings.Index(str1, "ˆ") != -1 || strings.Index(str1, "^") != -1 ) {
		return (true)
	}
	return (false)
}

func NegPui(str string, m string) (string) {

	if  str[len(str) - 1] == 134 || string(str[len(str) - 1]) == "^" {
		return (str + m)
	}
	return (str)
}

func Float2string(Calc TmpComp) (string) {

	if Calc.B == 0 {
		return (fmt.Sprintf("%f", Calc.A))
	} else if Calc.A == 0 {
		return (fmt.Sprintf("%fi", Calc.B))
	} else if Calc.B > 0 {
		return (fmt.Sprintf("%f + %fi", Calc.A, Calc.B))
	}
	return (fmt.Sprintf("%f %fi", Calc.A, Calc.B))
}

func ParseOne(str string, vars *types.Variable) (x float64, y float64) {

	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = replace_vars.GetVars(vars, str)
    str = strings.ReplaceAll(str, " ", "")
    str = strings.ReplaceAll(str, "\n", "")
    if str == "i" {
		str = "1i"
	}

    r, _ := regexp.Compile(`(?m)[+-]?([0-9]*[.])?[0-9]+[-+][+-]?([0-9]*[.])?[0-9]+[i]`)

	if strings.Index(str, "ˆ") != -1 || strings.Index(str, "^") != -1 {
		nstr := strings.Split(str, "ˆ")
		if len(nstr) == 1 {
			nstr = strings.Split(str, "^")
		}
		if !IsNumeric(nstr[0]) {
			nstr[0] = replace_vars.GetVars(vars, nstr[0])
		}
		if !IsNumeric(nstr[1]) {
			nstr[1] = replace_vars.GetVars(vars, nstr[1])
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

	var a, c, d, k float64

	Base := TmpComp{}

	for i := len(nstr) - 1; i > 0; i-- {

		nstr[i] = strings.ReplaceAll(nstr[i], " ", "")
		if nstr[i] == "i" {
			nstr[i] = "1i"
		}
		if nstr[i - 1] == "i" {
			nstr[i - 1] = "1i"
		}
		a, k = TransN(nstr[i])
		nstr[i - 1] = strings.ReplaceAll(nstr[i - 1], " ", "")
		if r.MatchString(nstr[i - 1]) {
			c, d = Trans(nstr[i - 1])
		} else {
			c, d = TransN(nstr[i - 1])
		}
		Base.A = c
		Base.B = d
		if a == 0 {
			a = k
		}
		Pow(&Base, int64(a))
		nstr[i - 1] = Float2string(Base)
	}

	return Base.A, Base.B
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

func IsCarre(m string, m1 string) (bool) {

	return (true)
}

/************************************************************************************************/

func Matrices(Finu *TmpComp, mat string, mat1 string, sign string, vars *types.Variable) {

	var r_mat, r_mat1 string

	if mat != "" {
		r_mat = vars.Table[mat].Value()
	}
	if mat1 != "" {
		r_mat1 = vars.Table[mat1].Value()
	}

	if r_mat1 != "" && r_mat != "" {
		
		if sign == "+" || sign == "-" {
			res := matrices.Modifi(Decomp(r_mat, r_mat1, sign, vars))
			vars.Table[mat] = &res
			return
		}
	}

	if r_mat != "" {
		if sign == "*" || sign == "/" {
			res := matrices.Modifi(CalcMatNb(r_mat, sign, Finu, vars))
			vars.Table[mat] = &res
			return
		}
	}

	if r_mat1 != "" {
		if sign == "*" || sign == "/" {
			res := matrices.Modifi(CalcMatNb(r_mat1, sign, Finu, vars))
			vars.Table[mat1] = &res
			return
		}
	}
}

func Decomp(m string, m1 string, sign string, vars *types.Variable) (string) {

	e := strings.Split(m, ";")
	aj := strings.Split(m1, ";")
	for i := 0; i < len(e); i++ {
		ex := strings.Split(e[i], ",")
		ajx := strings.Split(aj[i], ",")
		for z := 0; z < len(ex); z++ {
			nb1, nb2 := ParseOne(ex[z], vars)
			nb3, nb4 := ParseOne(ajx[z], vars)
			Calc := TmpComp{nb1, nb2}
			if sign == "+" {
				Add(&Calc, nb3, nb4)
			} else {
				Sous(&Calc, nb3, nb4)
			}
			ex[z] = Float2string(Calc)
		}
		e[i] = "[" + strings.Join(ex, ",") + "]"
	}
	m = "[" + strings.Join(e, ";") + "]"
	return (m)
}

func CalcMatNb(m string, s string, Finu *TmpComp, vars *types.Variable) (string) {

	e := strings.Split(m, ";")
	for i := 0; i < len(e); i++ {

		ex := strings.Split(e[i], ",")
		for z := 0; z < len(ex); z++ {
			nb1, nb2 := ParseOne(ex[z], vars)
			Calc := TmpComp{nb1, nb2}
			if s == "*" {
				Mul(&Calc, Finu.A, Finu.B)
			} else {
				Divi(&Calc, Finu.A, Finu.B)
			}
			ex[z] = Float2string(Calc)
		}
		e[i] = "[" + strings.Join(ex, ",") + "]"
	}
	m = "[" + strings.Join(e, ";") + "]"
	return (m)
}

func Add(Finu *TmpComp, a float64, b float64) {

	Finu.A = Finu.A + a
	Finu.B = Finu.B + b
}

func Mul(Finu *TmpComp, a float64, b float64) {

	tmp := ((Finu.A * a) - (Finu.B * b))
	Finu.B = ((Finu.A * b) + (a * Finu.B))
	Finu.A = tmp
}

func Divi(Finu *TmpComp, a float64, b float64) {

	tmp := ((Finu.A * a) + (Finu.B * b)) / ((a * a) + (b * b))
	Finu.B = ((Finu.B * a) - (Finu.A * b)) / ((a * a) + (b * b))
	Finu.A = tmp
}

func Sous(Finu *TmpComp, a float64, b float64) {

	Finu.A = Finu.A - a
	Finu.B = Finu.B - b
}

func Mod(Finu *TmpComp, a float64, b float64) {

	Calc := TmpComp{ Finu.A, Finu.B }
	Divi(&Calc, a, b)
	Calc.A = float64(int64(Calc.A))
	Calc.B = float64(int64(Calc.B))
	Mul(&Calc, a, b)
	Sous(Finu, Calc.A, Calc.B)
}

func Pow(n1 *TmpComp, n2 int64) {

	coe := n1.A
	im := n1.B

	if n2 == 0 {
		n1.A = 1
		n1.B = 0
		return
	}

    for i := int64(1); i < n2; i++ {
        
        Mul(n1, coe, im)
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
	
	Calc := TmpComp{n1.A, n1.B}
	Mul(&Calc, coe, im)
	if IsNan(Calc.A) || IsNan(Calc.B) {
		return (true)
	}
    return (false)
}

func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}