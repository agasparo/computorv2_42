package types

import (
	"fmt"
)

type AllT interface {

	Value() (string)
}

/**************************************************************/

type Rationel struct {

	Number float64
}

type Imaginaire struct {
	
	A float64
	B float64
}

type Matrice struct {
	
	Mat string
}

type Fonction struct {
	
	Func string
}

type EquaSol struct {

	Deg int
	Delta float64
	Sol string
}

type Variable struct {

	Table map[string]AllT
}

/**************************************************************/

func (r *Rationel) Value() (string) {

	if isFloatInt(r.Number) {
		return (fmt.Sprintf("%d", int64(r.Number)))
	}
	return (fmt.Sprintf("%f", r.Number))
}

func (i *Imaginaire) Value() (string) {

	if isFloatInt(i.A) && isFloatInt(i.B) {
		return (fmt.Sprintf("%d + %di", int64(i.A), int64(i.B)))
	}
	if isFloatInt(i.A) {
		return (fmt.Sprintf("%d + %fi", int64(i.A), i.B))
	}
	if isFloatInt(i.B) {
		return (fmt.Sprintf("%f + %di", i.A, int64(i.B)))
	}
	return (fmt.Sprintf("%f + %fi", i.A, i.B))
}

func (m *Matrice) Value() (string) {

	return (m.Mat)
}

func (f *Fonction) Value() (string) {

	return (f.Func)
}

func (e *EquaSol) Value() (string) {

	str := fmt.Sprintf("Equation degree : %d\n", e.Deg)
	if e.Deg > 1 {
		str += fmt.Sprintf("∆ = %f\n", e.Delta)
	}
	str += fmt.Sprintf("Solution(s) : %s\n", e.Sol)
	return (str)
}

func isFloatInt(floatValue float64) (bool) {

	if floatValue >= 9223372036854775807 || floatValue <= -9223372036854775808 {
		return false
	}
    return floatValue == float64(int64(floatValue))
}