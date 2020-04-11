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

type Variable struct {

	Table map[string]AllT
}

/**************************************************************/

func (r *Rationel) Value() (string) {

	fmt.Println(r.Number)
	return "ok"
}

func (i *Imaginaire) Value() (string) {

	return (fmt.Sprintf("%f+%fi", i.A, i.B))
}

func (m *Matrice) Value() (string) {

	return (m.Mat)
}

func (f *Fonction) Value() (string) {

	return (f.Func)
}