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
  	"unicode"
)

type Courbe struct {

	Funct string
	Interval_i int
	Interval_f int
	Name string
	Step float64
}

func Init(vars *types.Variable, str string, C *Courbe) {


	C.Funct = replace_vars.GetVars(vars, str)
	env_1, _ := strconv.ParseFloat(vars.Table["Interval_i"].Value(), 64)
	env_2, _ := strconv.ParseFloat(vars.Table["Interval_f"].Value(), 64)
	env_3, _ := strconv.ParseFloat(vars.Table["Interval_step"].Value(), 64)
	C.Interval_i = int(env_1)
	C.Interval_f = int(env_2)
	C.Step = env_3
	C.Name = str
}

func Trace(C Courbe, vars types.Variable) {

	All := []chart.Series{}
	All = CalcPoints(&C, vars, All)
	Draw(C, All)
}

func CalcPoints(C *Courbe, vars types.Variable, All []chart.Series) ([]chart.Series) {

	parser_err := 0
	tab := parser.GetAllIma(C.Name, &parser_err)
	tab = parser.Checkfunc(tab, vars)
	str := maths_functions.JoinTab(tab)
	br := 0
	var a, k, g float64
	var doi, nn int
	var tabx []float64
	var taby []float64
	var tmp []float64
	var title string

	for i := float64(C.Interval_i); i <= float64(C.Interval_f); i += C.Step {

		doi = 0
		if strings.Index(C.Funct, "|") != -1 {
			str = strings.ReplaceAll(C.Funct, "usu|", "")
			str = parser.Remp(str, maths_functions.Getx(C.Name), replace_vars.GetVars(&vars, fmt.Sprintf("%f", i)), vars)
			str = usuelles_functions.GetUsuF(str, vars)
			if strings.Index(str, "Impossible") != -1 {
				doi = 1
			}
			nn = 1
			a, _ = strconv.ParseFloat(str, 64)
		} else {
			if !IsInter(str, i) {
				doi = 1
			}
			a, _ = maths_functions.Calc(str, maths_functions.Getx(C.Name), fmt.Sprintf("%f", i), &vars)
		}
		if doi == 0 {
			tabx = append(tabx, i)
			taby = append(taby, a)
		} else {
			if nn == 1 {
				title = C.Name + " | " + GetInterval(i, C, br)
			} else {
				title = C.Name + " = " + C.Funct + " | " + GetInterval(i, C, br)
			}
			All = append(All, chart.ContinuousSeries {
	        	Name:    title,
	            XValues: tabx,
	            YValues: taby,
	        })
	       	tabx = tmp
	       	taby = tmp
	       	br = 1
	       	if i == 0 {
				k = 0 + C.Step	       		
	       	} else {
	       		k = i
	       	}
		}
		g = i
	}
	if k == 0 {
		k = g + C.Step
	}
	if nn == 1 {
		title = C.Name + " | " + GetInterval(k, C, br)
	} else {
		title = C.Name + " = " + C.Funct + " | " + GetInterval(k, C, br)
	}
	All = append(All, chart.ContinuousSeries {
	    Name:    title,
	    XValues: tabx,
	    YValues: taby,
	})
	return (All)
}

func IsInter(str string, a float64) (bool) {

	str = strings.ReplaceAll(str, " ", "")

	for i := 0; i < len(str); i++ {

		if str[i] == '/' && IsLetter(string(str[i + 1])) && a == 0 {
			return (false)
		}
		if str[i] == '%' && IsLetter(string(str[i + 1])) && a == 0 {
			return (false)
		}
	}
	return (true)
}

func IsLetter(s string) bool {

    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func GetInterval(i float64, C *Courbe, br int) (string) {

	a := float64(C.Interval_i)
	b := float64(C.Interval_f)

	if br == 1 {
		return (fmt.Sprintf("] %f ; %d ]", i, C.Interval_f))
	}

	if i > a && i < b {
		return (fmt.Sprintf("[ %d ; %f [", C.Interval_i, i))
	}

	if i > a && i >= b {
		return (fmt.Sprintf("[ %d ; %d ]", C.Interval_i, C.Interval_f))
	}
	return ("")
}

func Draw(C Courbe, All []chart.Series) {
	
	graph := chart.Chart{
	    Series: All,
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