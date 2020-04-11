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
	c := 0
	var indexs, elems string

	for index, element := range tab {

		c++
		indexs += index + "\t"
		elems += element.Value() + "\t"
	}

	if c == 0 {
		fmt.Println("no var set")
	} else {
		w := new(tabwriter.Writer)
		w.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.Debug|tabwriter.AlignRight)
		fmt.Fprintln(w, indexs)
    	fmt.Fprintln(w, elems)
    	fmt.Fprintln(w)
    	w.Flush()
	}
}