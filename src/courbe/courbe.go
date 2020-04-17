package courbe

import (
	"replace_vars"
	"types"
  	"strconv"
  	"maths_functions"
  	"github.com/wcharczuk/go-chart"
  	"os"
  	"fmt"
  	"os/exec"
  	"parser"
  	"usuelles_functions"
  	"strings"
)

type Courbe struct {

	Funct string
	Interval_i int
	Interval_f int
	Name string
} 

func Init(vars *types.Variable, str string, C *Courbe) {


	C.Funct = replace_vars.GetVars(vars, str)
	env_1, _ := strconv.ParseFloat(vars.Table["Interval_i"].Value(), 64)
	env_2, _ := strconv.ParseFloat(vars.Table["Interval_f"].Value(), 64)
	C.Interval_i = int(env_1)
	C.Interval_f = int(env_2)
	C.Name = str
}

func Trace(C Courbe, vars types.Variable) {

	var tabx []float64
	var taby []float64
	tabx, taby = CalcPoints(&C, tabx, taby, vars)
	Draw(C, tabx, taby)
}

func CalcPoints(C *Courbe, tabx []float64, taby []float64, vars types.Variable) ([]float64, []float64) { // refaire

	parser_err := 0
	tab := parser.GetAllIma(C.Name, &parser_err)
	tab = parser.Checkfunc(tab, vars)
	str := maths_functions.JoinTab(tab)
	var a float64
	var nn, doi int

	for i := C.Interval_i; i < C.Interval_f; i++ {

		doi = 0
		if strings.Index(C.Funct, "|") != -1 {
			str = strings.ReplaceAll(C.Funct, "usu|", "")
			str = parser.Remp(str, maths_functions.Getx(C.Name), replace_vars.GetVars(&vars, strconv.Itoa(i)), vars)
			str = usuelles_functions.GetUsuF(str, vars)
			if strings.Index(str, "Impossible") != -1 {
				doi = 1
			}
			a, _ = strconv.ParseFloat(str, 64)
			nn = 1
		} else {
			a, _ = maths_functions.Calc(str, maths_functions.Getx(C.Name), strconv.Itoa(i), &vars)
		}
		if doi == 0 {
			tabx = append(tabx, float64(i))
			taby = append(taby, a)
		}
	}
	if nn == 1 {
		C.Funct = ""
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

	f, _ := os.Create(C.Name + ".png")
	defer f.Close()
	graph.Render(chart.PNG, f)
	
	cmd := exec.Command("sh", "catimg.sh", C.Name + ".png")
    stdout, err := cmd.Output()

    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println(string(stdout))
    os.Remove(C.Name + ".png")
}