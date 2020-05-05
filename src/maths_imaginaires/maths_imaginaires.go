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
					if data[i + 1] == "*" {
						if !IsOkMul(data[i - 1], data[i + 1], vars) {
							data[0] = "Your matrice is not good for multiplication"
							return (data)
						}
						Calc = TmpComp{0, 0}
						Matrices(&Calc, data[i - 1], data[i + 2], "*", vars)
						data = maps.MapSlice(data, i) // mapsilecount
						data[i - 1] = data[i - 1]
					} else {
						data[0] = "Multiplication with matrices is with ** not just one *"
						return (data)
					}
				} else if strings.Index(data[i - 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i + 1], vars)
					Calc = TmpComp{nb1, nb2}
					Matrices(&Calc, data[i - 1], "", "*", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				} else if strings.Index(data[i + 1], "mat") != -1 {
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
			if strings.Index(data[i - 1], "mat") != -1 || strings.Index(data[i + 1], "mat") != -1 {
				data[0] = "Can't calcul matrice with modulo"
				return (data)
			}
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
					if !SizeMat(data[i - 1], data[i + 1], vars) {
						data[0] = "Your matrice must have the same length"
						return (data)
					}
					if !IsCarre(data[i - 1], data[i + 1], vars) {
						data[0] = "Your matrice must be carre"
						return (data)
					}
					Calc = TmpComp{0, 0}
					if IsNul(data[i + 1], vars) {
						data[0] = "Can't do division by 0"
						return (data)
					}
					if matrices.GetnbLine(vars.Table[data[i + 1]].Value()) > 3 {
						data[0] = "Sorry i can't calcul matrice more than 3 x 3 for division"
						return (data)
					}
					c, ci := MatDet(vars.Table[data[i + 1]].Value(), vars)
					if c == 0 && ci == 0 {
						data[0] = "Can't do division by det = 0 (matrice)"
						return (data)
					}
					Matrices(&Calc, data[i - 1], data[i + 1], "/", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				} else if strings.Index(data[i - 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i + 1], vars)
					nb3, _ := ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					if nb3 == 0 {
						data[0] = "Can't do division by 0"
						return (data)
					}
					Matrices(&Calc, data[i - 1], "", "m/", vars)
					data = maps.MapSlice(data, i)
					data[i - 1] = data[i - 1]
				} else if strings.Index(data[i + 1], "mat") != -1 {
					nb1, nb2 := ParseOne(data[i - 1], vars)
					Calc = TmpComp{nb1, nb2}
					if IsNul(data[i + 1], vars) {
						data[0] = "Can't do division by 0"
						return (data)
					}
					if !IsCarre(data[i + 1], data[i + 1], vars) {
						data[0] = "Your matrice must be carre"
						return (data)
					}
					if matrices.GetnbLine(vars.Table[data[i + 1]].Value()) > 3 {
						data[0] = "Sorry i can't calcul matrice more than 3 x 3 for division"
						return (data)
					}
					c, ci := MatDet(vars.Table[data[i + 1]].Value(), vars)
					if c == 0 && ci == 0 {
						data[0] = "Can't do division by det = 0 (matrice)"
						return (data)
					}
					Matrices(&Calc, data[i + 1], "", "/m", vars)
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
				if !SizeMat(data[i - 1], data[i + 1], vars) {
					data[0] = "Your matrice must have the same length"
					return (data)
				}
				Calc = TmpComp{0, 0}
				Matrices(&Calc, data[i - 1], data[i + 1], "+", vars)
				data = maps.MapSlice(data, i)
				data[i - 1] = data[i - 1]
			} else if strings.Index(data[i - 1], "mat") != -1 || strings.Index(data[i + 1], "mat") != -1 {
				data[0] = "I can't add matrice and number"
				return (data)
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
				if !SizeMat(data[i - 1], data[i + 1], vars) {
					data[0] = "Your matrice must have the same length"
					return (data)
				}
				Calc = TmpComp{0, 0}
				Matrices(&Calc, data[i - 1], data[i + 1], "-", vars)
				data = maps.MapSlice(data, i)
				data[i - 1] = data[i - 1]
			} else if strings.Index(data[i - 1], "mat") != -1 || strings.Index(data[i + 1], "mat") != -1 {
				data[0] = "I can't add matrice and number"
				return (data)
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

func SizeMat(m string, m1 string, vars *types.Variable) (bool) {

	ma := vars.Table[m].Value()
	ma1 := vars.Table[m1].Value()

	ml := matrices.GetnbLine(ma)
	mc := matrices.GetnbCol(ma)
	m1l := matrices.GetnbLine(ma1)
	m1c := matrices.GetnbCol(ma1)
	
	if ml != m1l || mc != m1c {
		return (false)
	}
	return (true)
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

func IsCarre(m string, m1 string, vars *types.Variable) (bool) {

	ma := vars.Table[m].Value()
	ma1 := vars.Table[m1].Value()

	ml := matrices.GetnbLine(ma)
	mc := matrices.GetnbCol(ma)
	m1l := matrices.GetnbLine(ma1)
	m1c := matrices.GetnbCol(ma1)
	
	if ml != mc || m1l != m1c {
		return (false)
	}
	return (true)
}

func IsOkMul(m string, m1 string, vars *types.Variable) (bool) {

	ma := vars.Table[m].Value()
	ma1 := vars.Table[m1].Value()

	ml := matrices.GetnbLine(ma)
	mc := matrices.GetnbCol(ma)
	m1l := matrices.GetnbLine(ma1)
	m1c := matrices.GetnbCol(ma1)

	if ml != m1c || m1l != mc {
		return (false)
	}
	return (true)
}

func IsNul(mat string, vars *types.Variable) (bool) {

	c := 0
	ma := vars.Table[mat].Value()
	m := strings.Split(ma, ";")
	for i := 0; i < len(m); i++ {
		ms := strings.Split(m[i], ",")
		for z := 0; z < len(ms); z++ {
			ms[z] = strings.ReplaceAll(ms[z], "[", "")
			ms[z] = strings.ReplaceAll(ms[z], "]", "")
			if ms[z] == "0" {
				c++
			}
		}
	}
	if c >= (matrices.GetnbCol(ma) * matrices.GetnbCol(ma)) {
		return (true)
	}
	return (false)
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

		if sign == "*" {
			res := matrices.Modifi(MulMa(r_mat, r_mat1, vars))
			vars.Table[mat] = &res
			return
		}

		if sign == "/" {
			det, deti := MatDet(r_mat1, vars)
			res2 := matrices.Modifi(Comatrice(r_mat1, vars))
			vars.Table[mat1] = &res2
			r_mat1 = vars.Table[mat1].Value()
			res3 := matrices.Modifi(Transcomatrice(r_mat1))
			vars.Table[mat1] = &res3
			r_mat1 = vars.Table[mat1].Value()
			Finu.A = 1
			Finu.B = 0
			Divi(Finu, det, deti)
			res1 := matrices.Modifi(CalcMatNb(r_mat1, "*", Finu, vars))
			vars.Table[mat1] = &res1
			r_mat1 = vars.Table[mat1].Value()
			res := matrices.Modifi(MulMa(r_mat, r_mat1, vars))
			vars.Table[mat] = &res
			return
		}
	}

	if r_mat != "" {
		if sign == "*" {
			res := matrices.Modifi(CalcMatNb(r_mat, sign, Finu, vars))
			vars.Table[mat] = &res
			return
		}
		if strings.Index(sign, "/") != -1 && sign[0] == 'm' {
			//ici matrice / nb
		}
	}

	if r_mat1 != "" {
		if sign == "*" {
			res := matrices.Modifi(CalcMatNb(r_mat1, sign, Finu, vars))
			vars.Table[mat1] = &res
			return
		}
		if strings.Index(sign, "/") != -1 && sign[len(sign) - 1] == 'm' {
			//ici nb / matrice
		}
	}
}

func Transcomatrice(m string) (string) {

	ml := matrices.GetnbLine(m)

	if ml == 1 {
		return (m)
	}

	if ml == 2 {

		m = strings.ReplaceAll(m, "[", "")
		m = strings.ReplaceAll(m, "]", "")
		cols := strings.Split(m, ";")
		Row0 := strings.Split(cols[0], ",")
		Row1 := strings.Split(cols[1], ",")

		str := "[[" + Row0[0] + "," + Row1[0] + "];[" + Row0[1] + "," + Row1[1] + "]]"
		return (str)
	}

	m = strings.ReplaceAll(m, "[", "")
	m = strings.ReplaceAll(m, "]", "")
	cols := strings.Split(m, ";")
	Row0 := strings.Split(cols[0], ",")
	Row1 := strings.Split(cols[1], ",")
	Row2 := strings.Split(cols[2], ",")
	nstr := "[[" + Row0[0] + "," + Row1[0] + "," + Row2[0] + "];[" + Row0[1] + "," + Row1[1] + "," + Row2[1] + "];[" + Row0[2] + "," + Row1[2] + "," + Row2[2] + "]]" 
	return (nstr)
}

func Comatrice(m string, vars *types.Variable) (string) {

	ml := matrices.GetnbLine(m)

	if ml == 1 {
		return (m)
	}

	if ml == 2 {

		m = strings.ReplaceAll(m, "[", "")
		m = strings.ReplaceAll(m, "]", "")
		cols := strings.Split(m, ";")
		Row0 := strings.Split(cols[0], ",")
		Row1 := strings.Split(cols[1], ",")

		nb1, nb2 := ParseOne(Row1[0], vars)
		nb3, nb4 := ParseOne(Row0[1], vars)

		TmpA := TmpComp{nb1, nb2}
		TmpB := TmpComp{nb3, nb4}

		Mul(&TmpA, -1, 0)
		Mul(&TmpB, -1, 0)

		str := "[[" + Row1[1] + "," + Float2string(TmpA) + "];[" + Float2string(TmpB) + "," + Row0[0] + "]]"
		return (str)
	}

	m = strings.ReplaceAll(m, "[", "")
	m = strings.ReplaceAll(m, "]", "")
	cols := strings.Split(m, ";")
	Row0 := strings.Split(cols[0], ",")
	Row1 := strings.Split(cols[1], ",")
	Row2 := strings.Split(cols[2], ",")

	SousMat := make(map[int]string)
	SousMat[0] = "[[" + Row1[1] + "," + Row1[2] + "];[" + Row2[1] + "," + Row2[2] + "]]"
	SousMat[1] = "[[" + Row1[0] + "," + Row1[2] + "];[" + Row2[0] + "," + Row2[2] + "]]"
	SousMat[2] = "[[" + Row1[0] + "," + Row1[1] + "];[" + Row2[0] + "," + Row2[1] + "]]"
	SousMat[3] = "[[" + Row0[1] + "," + Row0[2] + "];[" + Row2[1] + "," + Row2[2] + "]]"
	SousMat[4] = "[[" + Row0[0] + "," + Row0[2] + "];[" + Row2[0] + "," + Row2[2] + "]]"
	SousMat[5] = "[[" + Row0[0] + "," + Row0[1] + "];[" + Row2[0] + "," + Row2[1] + "]]"
	SousMat[6] = "[[" + Row0[1] + "," + Row0[2] + "];[" + Row1[1] + "," + Row1[2] + "]]"
	SousMat[7] = "[[" + Row0[0] + "," + Row0[2] + "];[" + Row1[0] + "," + Row1[2] + "]]"
	SousMat[8] = "[[" + Row0[0] + "," + Row0[1] + "];[" + Row1[0] + "," + Row1[1] + "]]"

	nstr := "[["
	for i := 0; i < len(SousMat); i++ {
		det, deti := MatDet(SousMat[i], vars)
		Tmp := TmpComp{ det, deti }
		if i % 3 == 0 && i != 0 {
			nstr += "];"
		} else {
			if i > 0 {
				nstr += ","
			}
		}
		if i % 3 == 0 && i != 0 {
			nstr += "["
		}
		if i % 2 == 0 {
			nstr += Float2string(Tmp)
		} else {
			Mul(&Tmp, -1, 0)
			nstr += Float2string(Tmp)
		}
	}
	nstr += "]"
	return (nstr)
}

func MatDet(m string, vars *types.Variable) (float64, float64) {

	ml := matrices.GetnbLine(m)

	if ml == 1 {
		cols := strings.Split(m, ";")
		Row0 := strings.Split(cols[0], ",")
		return ParseOne(Row0[0], vars)
	}

	if ml == 2 {
		cols := strings.Split(m, ";")
		Row0 := strings.Split(cols[0], ",")
		Row1 := strings.Split(cols[1], ",")
		nb1, nb1i := ParseOne(Row0[0], vars)
		nb2, nb2i := ParseOne(Row1[1], vars)
		nb3, nb3i := ParseOne(Row0[1], vars)
		nb4, nb4i := ParseOne(Row1[0], vars)
		TmpA := TmpComp{nb1, nb1i}
		Mul(&TmpA, nb2, nb2i)
		TmpB := TmpComp{nb3, nb3i}
		Mul(&TmpB, nb4, nb4i)
		Sous(&TmpA, TmpB.A, TmpB.B)
		return TmpA.A, TmpA.B
	}

	p1, p1i := CalcVars(1, 2, 0, m, vars)
	p2, p2i := CalcVars(1, 0, 2, m, vars)
	Tmp := TmpComp{ p1, p1i }
	Sous(&Tmp, p2, p2i)
	return Tmp.A, Tmp.B
}

func CalcVars(j int, k int, l int, m string, vars *types.Variable) (float64, float64) {

	cols := strings.Split(m, ";")
	Row0 := strings.Split(cols[0], ",")
	Row1 := strings.Split(cols[1], ",")
	Row2 := strings.Split(cols[2], ",")

	Tmp := TmpComp{ 0, 0 }

	for i := 3; i > 0; i-- {

		nb1, nb1i := ParseOne(Row0[l], vars)
		nb2, nb2i := ParseOne(Row1[j], vars)
		nb3, nb3i := ParseOne(Row2[k], vars)

		TmpA := TmpComp{ nb1, nb1i }
		Mul(&TmpA, nb2, nb2i)
		Mul(&TmpA, nb3, nb3i)
		Add(&Tmp, TmpA.A, TmpA.B)

		k--
		j--
		l--
		if l < 0 {
			l = 2
		}
		if k < 0 {
			k = 2
		}
		if j < 0 {
			j = 2
		}
	}
	return Tmp.A, Tmp.B
}

func MulMa(m string, m1 string, vars *types.Variable) (string) {

	m0 := strings.Split(m, ";")
	ma1 := strings.Split(m1, ";")
	nstr := ""
	
	for a := 0; a < len(m0); a++ {
		m0a := strings.Split(m0[a], ",")
		for i := 0; i < matrices.GetnbCol(m1); i++ {
			Calc := TmpComp{}
			for z := 0; z < len(m0a); z++ {
				ma2 := strings.Split(ma1[z], ",")
				nb1, nb2 := ParseOne(m0a[z], vars)
				nb3, nb4 := ParseOne(ma2[i], vars)
				Tmp := TmpComp{nb1, nb2}
				Mul(&Tmp, nb3, nb4)
				Add(&Calc, Tmp.A, Tmp.B)
			}
			nstr += Float2string(Calc)
			if i  + 1 != matrices.GetnbLine(m) {
				nstr += ","
			} else {
				nstr += ";"
			}
		}
	}
	nstr = nstr[0:len(nstr) - 2]
	return (nstr)
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