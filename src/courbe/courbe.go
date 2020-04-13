package courbe

import (
	"replace_vars"
	"types"
  	"strconv"
  	"maths_functions"
  	"github.com/wcharczuk/go-chart"
  	"os"
)

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

	var tabx []float64
	var taby []float64
	tabx, taby = CalcPoints(&C, tabx, taby, vars)
	Draw(C, tabx, taby)
}

func CalcPoints(C *Courbe, tabx []float64, taby []float64, vars *types.Variable) ([]float64, []float64) {

	c := 0

	for i := C.Interval_i; i < C.Interval_f; i++ {

		a, _ := maths_functions.Calc(C.Funct, maths_functions.Getx(C.Name), strconv.Itoa(i), vars)
		tabx = append(tabx, float64(i))
		taby = append(taby, a)
		c++
	}
	return tabx, taby
}

func Draw(C Courbe, tabx []float64, taby []float64) {
	
	graph := chart.Chart{
	    Series: []chart.Series{
	        chart.ContinuousSeries{
	        	Name:    C.Name + " = " + C.Funct,
	            XValues: tabx,
	            YValues: taby,
	        },
	    },
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph),
	}

	f, _ := os.Create("res_graph/" + C.Name + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
}