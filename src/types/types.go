package types

import (
	"fmt"
	"strconv"
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
	return (strconv.FormatFloat(r.Number, 'g', -1, 64))
}

func (i *Imaginaire) Value() (string) {

	if isFloatInt(i.A) && isFloatInt(i.B) {
		return (fmt.Sprintf("%d + %di", int64(i.A), int64(i.B)))
	}
	if isFloatInt(i.A) {
		return (fmt.Sprintf("%d + %si", int64(i.A), strconv.FormatFloat(i.B, 'g', -1, 64)))
	}
	if isFloatInt(i.B) {
		return (fmt.Sprintf("%s + %di", strconv.FormatFloat(i.A, 'g', -1, 64), int64(i.B)))
	}
	return (fmt.Sprintf("%s + %si", strconv.FormatFloat(i.A, 'g', -1, 64), strconv.FormatFloat(i.B, 'g', -1, 64)))
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
		str += fmt.Sprintf("∆ = %s\n", strconv.FormatFloat(e.Delta, 'g', -1, 64))
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