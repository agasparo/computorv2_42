package courbe

import (
	"replace_vars"
	"types"
	"os/exec"
  	"fmt"
  	"os"
)

type Courbe struct {

	Funct string
	X int
	Y int
	Echelle_x int
	Echelle_y int
} 

func Init(vars *types.Variable, x_max int, y_max int, str string, C *Courbe) {


	C.Funct = replace_vars.GetVars(vars, str)
	C.X = x_max
	C.Y = y_max
	C.Echelle_x = 0
	C.Echelle_y = 0
	fmt.Println(GetSize())
}

func Trace() {

}

func GetSize() {
	cmd := exec.Command("stty", "size")
  	cmd.Stdin = os.Stdin
  	out, err := cmd.Output()
  	return (string(out))
}