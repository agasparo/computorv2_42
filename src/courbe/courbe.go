package courbe

import (
	"replace_vars"
	"types"
	"os/exec"
  	"fmt"
  	"os"
  	"strings"
  	"strconv"
  	"maths_functions"
)

type Complexe struct {

	a float64
	b float64
}

type Courbe struct {

	Funct string
	X int
	Y int
	Echelle_x int
	Echelle_y int
	Interval_i int
	Interval_f int
	Name string
} 

func Init(vars *types.Variable, str string, C *Courbe) {


	C.Funct = replace_vars.GetVars(vars, str)
	size := strings.Split(strings.ReplaceAll(GetSize(), "\n", ""), " ")
	C.X, _ = strconv.Atoi(size[0])
	C.Y, _ = strconv.Atoi(size[1])
	C.Interval_i = 0
	C.Interval_f = 50
	C.Name = str
}

func Trace(C Courbe, vars *types.Variable) {

	tab := make(map[int]Complexe)
	CalcPoints(&C, tab, vars)
	fmt.Println(tab)
}

func CalcPoints(C *Courbe, tab map[int]Complexe, vars *types.Variable) {

	c := 0

	for i := C.Interval_i; i < C.Interval_f; i++ {

		a, b := maths_functions.Calc(C.Funct, maths_functions.Getx(C.Name), strconv.Itoa(i), vars)
		tab[c] = Complexe{ a, b }
		c++
	}
	C.Echelle_x = int(float64(C.X) / tab[c - 1].a)
	C.Echelle_y = int(float64(C.Y) / tab[c - 1].b)
}

func GetSize() (string) {
	cmd := exec.Command("stty", "size")
  	cmd.Stdin = os.Stdin
  	out, _ := cmd.Output()
  	return (string(out))
}