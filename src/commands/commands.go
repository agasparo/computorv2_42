package commands

import (
	"fmt"
	"types"
	"text/tabwriter"
	"os"
	"courbe"
	"parser"
	"strconv"
	"error"
	"strings"
)

func IsCommand(str string, str1 string, str2 string, Vars types.Variable) (int) {

	if str == "-list" {
		GetAllVars(Vars.Table)
		return (1)
	}
	if str == "-help" {
		Help()
		return (1)
	}
	if str == "-graph" {
		Graph(str1, Vars)
		return(1)
	}
	if str == "-set" {
		SetVars(str1, str2, Vars)
		return (1)
	}
	return (0)
}

func GetAllVars(tab map[string]types.AllT) {

	fmt.Println("List of all vars : ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	for index, element := range tab {

		if index != "?" {
			fmt.Fprintln(w, index + "\t" + element.Value() + "\t")
		}
	}
    fmt.Fprintln(w)
    w.Flush()
}

func Help() {

	fmt.Println("List of commands : ")
	fmt.Println("1 : '-list' -> List all vars")
	fmt.Println("2 : '-graph [function]' -> show a courbe of the function")
}

func Graph(str string, Vars types.Variable) {

	if _, ok := Vars.Table[strings.ToLower(str)]; ok {

		if parser.IsFunc(str, 0) == 1 {
			C := courbe.Courbe{}
			courbe.Init(&Vars, str, &C)
			courbe.Trace(C, Vars)
		} else {
			error.SetError(str + " isn't a function")
		}
    } else {
    	error.SetError(str + " function doesn't exist")
    }
}

func SetVars(str string, str1 string, Vars types.Variable) {

	if str == "Interval_i" || str == "Interval_f" {

		if parser.IsNumeric(str1) {
			
			a, _ := strconv.ParseFloat(str1, 64)
			b := int(a)
			a = float64(b)
			Vars.Table[str] = &types.Rationel{ a }
		}
	}

	if str == "Interval_step" {

		if parser.IsNumeric(str1) {

			a, _ := strconv.ParseFloat(str1, 64)
			b := int(a)
			if a - float64(b) >= 0.1 {
				Vars.Table[str] = &types.Rationel{ a }
			}
		}
	}
}