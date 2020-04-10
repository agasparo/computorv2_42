package types

import (
	"fmt"
)

type AllT interface {

	Value()
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

func (r *Rationel) Value() {

	fmt.Println(r.Number)
	return
}

func (i *Imaginaire) Value() {

	fmt.Printf("%f + i%f\n", i.A, i.B)
	return
}

func (m *Matrice) Value() {

	fmt.Println(m.Mat)
	return
}

func (f *Fonction) Value() {

	fmt.Println(f.Func)
	return
}