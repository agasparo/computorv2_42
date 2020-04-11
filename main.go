package main

import (
	"input"
	"types"
	"parser"
	"error"
	"show"
	"maths_imaginaires"
	"fmt"
	"strings"
)

func main() {

	Inputs := input.Data{}
	Vars := types.Variable{}

	Vars.Table = make(map[string]types.AllT)
	for i := 1; i == 1; i = 1 {
		input.ReadSTDIN(&Inputs)
		if Inputs.Input[0] == "exit" {
			fmt.Println("bye")
			return
		}
		r, t, v := basic_check(Inputs, &Vars)
		if r == 1 {
			show.ShowVars(t, Vars.Table[v])
		}
	}
}

func basic_check(Inputs input.Data, Vars *types.Variable) (int, int, string) {

	t := -1

	if parser.Array_search_count(Inputs.Input, "=") != 1 {
		error.SetError("You must have just one =")
		return 0, 0, ""
	}

	str := strings.Split(strings.Join(Inputs.Input, " "), "=")
	str[0] = strings.ToLower((strings.Trim(str[0], " ")))
	str[1] = strings.Trim(str[1], " ")

	if parser.IsFunc(str[0]) == 1 {

		Vars.Table[str[0]] = &types.Fonction{ str[1] }
		t = 0

	} else if strings.Index(str[1], "i") != -1 {

		fmt.Println("imaginaires")
		data := parser.GetAllIma(strings.ReplaceAll(str[1], " ", ""))
		x, y := maths_imaginaires.CalcVar(data, Vars)
		Vars.Table[str[0]] = &types.Imaginaire{ x, y }
		t = 0
	} /*else if strings.Index(str[0], "mat") != -1 || strings.Index(str[0], "var") != -1 {
		
		fmt.Println("matrice")
	} else if  parser.IsNumeric(str[1]) {

		//val, _ := strconv.ParseFloat(str[1], 64)
		//Vars.Table[str[0]] = &types.Rationel{ val }
		//t = 0
		fmt.Println("rationel")
	} else if val, ok := Vars.Table[str[1]]; ok {

		fmt.Println("var")
    	Vars.Table[str[0]] = val
    	//t = 1
    }*/

	return 1, t, str[0]
}