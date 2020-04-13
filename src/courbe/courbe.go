package courbe

import (
	"replace_vars"
	"types"
  	"fmt"
  	"strings"
  	"strconv"
  	"maths_functions"
  	"github.com/wcharczuk/go-chart"
  	"os"
)

type Complexe struct {

	a_y float64
	b_y float64
	a_x float64
	b_x float64
}

type Courbe struct {

	Funct string
	Interval_i int
	Interval_f int
	Name string
} 

func Init(vars *types.Variable, str string, C *Courbe) {


	C.Funct = replace_vars.GetVars(vars, str)
	C.Interval_i = 0
	C.Interval_f = 50
	C.Name = str
}

func Trace(C Courbe, vars *types.Variable) {

	tab := make(map[int]Complexe)
	CalcPoints(&C, tab, vars)
	fmt.Println(C)
	fmt.Println(tab)
	//Draw()
}

func CalcPoints(C *Courbe, tab map[int]Complexe, vars *types.Variable) {

	c := 0

	for i := C.Interval_i; i < C.Interval_f; i++ {

		a, b := maths_functions.Calc(C.Funct, maths_functions.Getx(C.Name), strconv.Itoa(i), vars)
		tab[c] = Complexe{ a, b, float64(i), 0 }
		c++
	}
}

/*func Draw() {
	graph := chart.Chart{
	    Series: []chart.Series{
	        chart.ContinuousSeries{
	            XValues: []float64{1.0, 2.0, 3.0, 4.0},
	            YValues: []float64{1.0, 2.0, 3.0, 4.0},
	        },
	    },
	}

	f, _ := os.Create("output.png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}*/