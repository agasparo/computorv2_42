package equations

import (
	"maths_imaginaires"
	"fmt"
	"fractions"
)

type Equation struct {

	A float64
	B float64
	C float64
}

func Resolve(Eq map[int]Equation) (int, float64, string) {

	var a int
	var b float64
	var c string

	fmt.Println(Eq)
	if len(Eq) == 2 {
		a, b, c = ResolveNorm(Eq[0], Eq[1])
	}
	return a, b, c
}

func ResolveNorm(a Equation, b Equation) (int, float64, string) {

	var deg int
	var delta float64
	var sol string

	Feq := Equation{}
	Feq.A = a.A + Inverse(b.A)
	Feq.B = a.B + Inverse(b.B)
	Feq.C = a.C + Inverse(b.C)

	if Feq.A != 0 {
		deg, delta, sol = ResolveEqDelta(Feq)
	} else if Feq.B != 0 {
		deg, sol = ResolveDeg1(Feq)
	} else if Feq.A == 0 && Feq.B == 0 {
		deg, sol = ResolveDeg0(a, b)
	}

	return deg, delta, sol
}

func ResolveDeg0(a Equation, b Equation) (int, string) {

	if a.C == b.C {
		return 0, "All Real numbers are solutions"
	}
	return 0, "No solutions"
}

func ResolveEqDelta(eq Equation) (int, float64, string) {

	return Delta(eq)
}

func Delta(Eq Equation) (int, float64, string) {

	var sol string

	St := maths_imaginaires.TmpComp{ Eq.B, 0 }
	maths_imaginaires.Pow(&St, int64(2))
	other := 4 * Eq.A * Eq.C
	delta := St.A - other

	if delta < 0 {
		sol = "Δ < 0, 2 solutions : x1 = (- b + i√-Δ) / 2a and x2 = (- b - i√-Δ) / 2a\n" + Deltainf(Eq, delta)
		return 2, delta, sol
	}

	if delta == 0 {
		sol = "Δ == 0, One solution : x0 = -b / 2a\n" + DeltaNil(Eq, delta)
		return 2, delta, sol
	}

	sol = "Δ > 0, 2 solutions : x1 = (- b + √Δ) / 2a and x2 = (- b - √Δ) / 2a\n" + DeltaSup(Eq, delta)
	return 2, delta, sol
}

func DeltaNil(Eq Equation, delta float64) (string) {
	
	var str string

	b := Eq.B * -1
	other := Eq.A * 2
	res := b / other
	if isFloatInt(res) {
		str = fmt.Sprintf("x0 = %d", int64(res))
		return (str)
	}
	str = fmt.Sprintf("x0 = %f", res)
	Rational := fractions.Rational{res, 0, 0, "", 3, ""}
	fractions.Trasnform(&Rational)
	if len(Rational.Frac) > 0 {
		str += fmt.Sprintf("ou x0 = %s", Rational.Frac)
	}
	return (str)
}

func Deltainf(Eq Equation, delta float64) (string) {

	deb := (Eq.B * -1)
	fin := (Eq.A * 2)

	if isFloatInt(Sqrt(Inverse(delta))) {

		i1 := Sqrt(Inverse(delta)) / fin
		i2 := (Sqrt(Inverse(delta)) / fin) * -1
		a1 := deb / fin
		a2 := deb / fin

		Rationalx1 := fractions.Rational{a1, 0, 0, "", 3, ""}
		Rationali1 := fractions.Rational{i1, 0, 0, "", 3, ""}
		fractions.Trasnform(&Rationalx1)
		fractions.Trasnform(&Rationali1)

		str1 := fmt.Sprintf("x1 = %f + %fi", a1, i1)
		if len(Rationalx1.Frac) > 0 && len(Rationali1.Frac) > 0 {
			str1 += fmt.Sprintf("ou x1 = %s + %si", Rationalx1.Frac, Rationali1.Frac)
		} else if len(Rationalx1.Frac) > 0 {
			str1 += fmt.Sprintf("ou x1 = %s + %fi", Rationalx1.Frac, i1)
		} else if len(Rationali1.Frac) > 0 {
			str1 += fmt.Sprintf("ou x1 = %f + %si", a1, Rationali1.Frac)
		}

		Rationalx2 := fractions.Rational{a2, 0, 0, "", 3, ""}
		Rationali2 := fractions.Rational{i2, 0, 0, "", 3, ""}
		fractions.Trasnform(&Rationalx2)
		fractions.Trasnform(&Rationali2)

		str2 := fmt.Sprintf("\nx2 = %f + %fi", a2, i2)
		if len(Rationalx2.Frac) > 0 && len(Rationali2.Frac) > 0 {
			str2 += fmt.Sprintf("ou x2 = %s + %si", Rationalx2.Frac, Rationali2.Frac)
		} else if len(Rationalx2.Frac) > 0 {
			str2 += fmt.Sprintf("ou x2 = %s + %fi", Rationalx2.Frac, i2)
		} else if len(Rationali2.Frac) > 0 {
			str2 += fmt.Sprintf("ou x2 = %f + %si", a2, Rationali2.Frac)
		}
		return (str1 + str2)
	}

	str1 := fmt.Sprintf("x1 = (%f + i√%f) / %f\n", deb, Inverse(delta), fin)
	str2 := fmt.Sprintf("x2 = (%f - i√%f) / %f", deb, Inverse(delta), fin)
	return (str1 + str2)
}

func DeltaSup(Eq Equation, delta float64) (string) {

	deb := (Eq.B * -1)
	fin := (Eq.A * 2)
	deb_x1 := deb + Sqrt(delta)
	deb_x2 := deb - Sqrt(delta)
	fin_x1 := deb_x1 / fin
	fin_x2 := deb_x2 / fin

	if isFloatInt(Sqrt(delta)) {
		Rationalx1 := fractions.Rational{fin_x1, 0, 0, "", 3, ""}
		fractions.Trasnform(&Rationalx1)
		str1 := fmt.Sprintf("x1 = %f ", fin_x1)
		if len(Rationalx1.Frac) > 0 {
			str1 += fmt.Sprintf("ou x1 = %s", Rationalx1.Frac)
		}
		Rationalx2 := fractions.Rational{fin_x2, 0, 0, "", 3, ""}
		fractions.Trasnform(&Rationalx2)
		str2 := fmt.Sprintf("\nx2 = %f ", fin_x2)
		if len(Rationalx2.Frac) > 0 {
			str2 += fmt.Sprintf("ou x2 = %s", Rationalx2.Frac)
		}
		return (str1 + str2)
	}

	str1 := fmt.Sprintf("x1 = (%f + √%f) / %f\n", deb, delta, fin)
	str2 := fmt.Sprintf("x2 = (%f - √%f) / %f", deb, delta, fin)
	return (str1 + str2)
}

func ResolveDeg1(eq Equation) (int, string) {

	sol := Inverse(eq.C) / eq.B
	if !isFloatInt(sol) {
		Rational := fractions.Rational{sol, 0, 0, "", 3, ""}
		fractions.Trasnform(&Rational)
		return 1, fmt.Sprintf("x0 = %f ou x0 = %s", sol, Rational.Frac)
	}
	return 1, fmt.Sprintf("x0 = %f", sol)
}

func Inverse(a float64) (float64) {

	return (a * -1)
}

func isFloatInt(floatValue float64) (bool) {

	if floatValue >= 9223372036854775807 || floatValue <= -9223372036854775808 {
		return false
	}
    return floatValue == float64(int64(floatValue))
}

func Sqrt(x float64) float64 {
    
    var z float64 = 1

    for i := 1; i <= 10; i++ {

    	z = (z - (Pow(z, 2) - x) / ( 2 * z))
    }
    return (z)
}

func Pow(x float64, n int) (res float64) {

    number := 1.00;

    for i := 0; i < n; i++ {
        number *= x;
    }

    return (number);
}