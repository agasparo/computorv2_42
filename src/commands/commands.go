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
	return (0)
}

func GetAllVars(tab map[string]types.AllT) {

	fmt.Println("List of all vars : ")

	w := new(tabwriter.Writer)
	var indexs, elems string

	for index, element := range tab {

		indexs += index + "\t"
		elems += element.Value() + "\t"
	}

	w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(w, indexs)
    fmt.Fprintln(w, elems)
    fmt.Fprintln(w)
    w.Flush()
}