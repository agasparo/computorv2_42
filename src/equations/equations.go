package equations

import (
	"maths_imaginaires"
	"fmt"
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
	
	b := Eq.B * -1
	other := Eq.A * 2
	res := b / other
	return (fmt.Sprintf("x0 = %f", res))
}

func Deltainf(Eq Equation, delta float64) (string) {

	deb := (Eq.B * -1)
	fin := (Eq.A * 2)
	Rationalx1 := fmt.Sprintf("x1 = (%f + i√%f) / %f", deb, Inverse(delta), fin)
	Rationalx2 := fmt.Sprintf("x2 = (%f - i√%f) / %f", deb, Inverse(delta), fin)
	return (Rationalx1 + Rationalx2)
}

func DeltaSup(Eq Equation, delta float64) (string) {

	deb := (Eq.B * -1)
	fin := (Eq.A * 2)
	Rationalx1 := fmt.Sprintf("x1 = (%f + √%f) / %f", deb, delta, fin)
	Rationalx2 := fmt.Sprintf("x2 = (%f - √%f) / %f", deb, delta, fin)
	return (Rationalx1 + Rationalx2)
}

func ResolveDeg1(eq Equation) (int, string) {

	sol := eq.C / eq.B
	return 1, fmt.Sprintf("x0 = %f", sol)
}

func Inverse(a float64) (float64) {

	return (a * -1)
}