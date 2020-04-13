package commands

import (
	"fmt"
	"types"
	"text/tabwriter"
	"os"
)

func IsCommand(str string, Vars types.Variable) (int) {

	if str == "-list" {
		GetAllVars(Vars.Table)
		return (1)
	}
	if str == "-help" {
		Help()
		return (1)
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
}