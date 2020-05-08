package types

import (
	"fmt"
	"time"
	"strings"
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
	
	Mat map[int]MatRow
}

type MatRow struct {

	Row map[int]string
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

type Histo struct {

	Table map[int]HistoData
}

type HistoData struct {

	When time.Time
	Command string
	Res string
}

/**************************************************************/

func (hd *HistoData) Value() (string) {

	return (fmt.Sprintf("[%s] %s --> %s", hd.When.Format("Mon Jan _2 15:04:05 2006"), hd.Command, strings.ReplaceAll(hd.Res, "\n", "")))
}

func (r *Rationel) Value() (string) {

	tmp := r.Number

	if isFloatInt(tmp) {
		return (fmt.Sprintf("%d", int64(r.Number)))
	}
	return (fmt.Sprintf("%f", r.Number))
}

func (i *Imaginaire) Value() (string) {

	tmp_a := i.A
	tmp_b := i.B

	if i.B == 0 {

		tmpS := Rationel{ i.A }
		return (tmpS.Value())
	}

	if isFloatInt(tmp_a) && isFloatInt(tmp_b) {
		return (fmt.Sprintf("%d + %di", int64(i.A), int64(i.B)))
	}
	if isFloatInt(tmp_a) {
		return (fmt.Sprintf("%d + %fi", int64(i.A), i.B))
	}
	if isFloatInt(tmp_b) {
		return (fmt.Sprintf("%f + %di", int64(i.A), i.B))
	}
	return (fmt.Sprintf("%f + %fi", i.A, i.B))
}

func (m *Matrice) Value() (string) {

	str := "["

	for i := 0; i < len(m.Mat); i++ {

		str += "["
		for z := 0; z < len(m.Mat[i].Row); z++ {
			nb, _ := strconv.ParseFloat(m.Mat[i].Row[z], 64)
			tmpnb := nb
			if isFloatInt(nb) {
				str = str + fmt.Sprintf("%d", int64(tmpnb))
			} else {
				str = str + m.Mat[i].Row[z]
			}
			if z + 1 < len(m.Mat[i].Row) {
				str += ","
			}
		}
		if i + 1 < len(m.Mat) {
			str += "];"
		} else {
			str += "]"
		}
	}
	str += "]"
	return (str)
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