package commands

import (
	"fmt"
	"types"
	"text/tabwriter"
	"os"
	"courbe"
)

func IsCommand(str string, str1 string, Vars types.Variable) (int) {

	if str == "-list" {
		GetAllVars(Vars.Table)
		return (1)
	}
	if str == "-help" {
		Help()
		return (1)
	}
	if str == "-graph" {
		Grap(str1, Vars)
		return(1)
	}
	return (0)
}

func GetAllVars(tab map[string]types.AllT) {

	fmt.Println("List of all vars : ")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	for index, element := range tab {

		fmt.Fprintln(w, index + "\t" + element.Value() + "\t")
	}
    fmt.Fprintln(w)
    w.Flush()
}

func Help() {

	fmt.Println("List of commands : ")
	fmt.Println("1 : '-list' -> List all vars")
	fmt.Println("2 : '-graph [function]' -> show a courbe of the function")
}

func Grap(str string, Vars types.Variable) {

	C := courbe.Courbe{}
	courbe.Init(&Vars, str, &C)
	fmt.Println(C)
	courbe.Trace(C, &Vars)
}