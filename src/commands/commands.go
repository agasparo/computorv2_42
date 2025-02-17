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
	"convert"
)

func IsCommand(str string, str1 string, str2 string, Vars types.Variable, Histo types.Histo) (int) {

	if str == "list" {
		GetAllVars(Vars.Table)
		return (1)
	}
	if str == "help" {
		Help()
		return (1)
	}
	if str == "graph" {
		Graph(str1, Vars)
		return(1)
	}
	if str == "set" {
		SetVars(str1, str2, Vars)
		return (1)
	}
	if str == "conv" {
		Convert(str1, str2)
		return (1)
	}
	if str == "histo" {
		Histori(Histo.Table)
		return (1)
	}
	return (0)
}

func Convert(str1 string, str2 string) {

	convert.Wicht(str1, str2)
}

func Histori(tab map[int]types.HistoData) {

	fmt.Println("List of your history : ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	for _, element := range tab {

		fmt.Fprintln(w, element.Value())
	}
    fmt.Fprintln(w)
    w.Flush()
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
	fmt.Println("1 : 'list' -> List all vars")
	fmt.Println("2 : 'graph [function]' -> show a courbe of the function")
	fmt.Println("3 : 'set [var]' -> allow to modify var value")
	fmt.Println("4 : 'conv [value-unite] [unite to convert]' -> convert value")
}

func Graph(str string, Vars types.Variable) {

	if val, ok := Vars.Table[strings.ToLower(str)]; ok {

		if parser.IsFunc(str, 0) == 1 {
			if strings.Index(val.Value(), "]") == -1 || strings.Index(val.Value(), "[") == -1 {
				C := courbe.Courbe{}
				courbe.Init(&Vars, str, &C)
				courbe.Trace(C, Vars)
			} else {
				error.SetError("I can't draw grap with matrice")
			}
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
			if a >= -50 && a <= 50 {
				a = float64(b)
				Vars.Table[str] = &types.Rationel{ a }
			} else {
				error.SetError(str + " must be between -50 and 50")
			}
		}
	}

	if str == "Interval_step" {

		if parser.IsNumeric(str1) {

			a, _ := strconv.ParseFloat(str1, 64)
			b := int(a)
			if a - float64(b) >= 0.1 && a - float64(b) <= 2 {
				Vars.Table[str] = &types.Rationel{ a }
			} else {
				error.SetError(str + " must be between 0.1 and 2")
			}
		}
	}
}