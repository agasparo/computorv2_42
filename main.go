package main

import (
	"input"
	"types"
	"parser"
	"error"
	"show"
	"commands"
	"maths_imaginaires"
	"maths_functions"
	"fmt"
	"strings"
	"parentheses"
	"usuelles_functions"
	"test"
	"os"
)

func main() {

	args := os.Args[1:]

	if len(args) == 1 && args[0] == "test" {
		test.DefineAndRun()
	} else if len(args) == 1 {
		RunTest(args[0])
	} else {
		Run()
	}
}

func Run() {
	Inputs := input.Data{}
	Vars := types.Variable{}
	arg := ""

	Vars.Table = make(map[string]types.AllT)
	usuelles_functions.Init(&Vars)
	for i := 1; i == 1; i = 1 {
		input.ReadSTDIN(&Inputs)
		if Inputs.Input[0] == "exit" {
			fmt.Println("bye")
			return
		}
		if Inputs.Length == 2 {
    		arg = Inputs.Input[1]
		}
		if commands.IsCommand(Inputs.Input[0], arg, Vars) != 1 {
			r, t, v := basic_check(Inputs, &Vars, Vars)
			if r == 1 {
				show.ShowVars(t, Vars.Table[v])
			}
		}
	}	
}

func RunTest(str string) {

	Inputs := input.Data{ strings.Split(str, " "), 1 }
	Vars := types.Variable{}
	arg := ""

	Vars.Table = make(map[string]types.AllT)
	usuelles_functions.Init(&Vars)
	if Inputs.Input[0] == "exit" {
		fmt.Println("bye")
		return
	}
	if Inputs.Length == 2 {
   		arg = Inputs.Input[1]
	}
	if commands.IsCommand(Inputs.Input[0], arg, Vars) != 1 {
		r, t, v := basic_check(Inputs, &Vars, Vars)
		if r == 1 {
			show.ShowVars(t, Vars.Table[v])
		}
	}
}

func basic_check(Inputs input.Data, Vars *types.Variable, Dat types.Variable) (int, int, string) {

	t := -1

	tmp := strings.Join(Inputs.Input, " ")
	str := strings.Split(tmp, "=")
	e := error.Syntaxe(tmp)

	if e != "1" {
		error.SetError(e)
		return 0, 0, ""
	}

	str[0] = strings.ToLower((strings.Trim(str[0], " ")))
	str[1] = strings.Trim(str[1], " ")
	str_ret := str[0]

	if str[1] == "?" {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[0]), " ", ""))
		e = error.In(data, 0, "")
		if e != "1" {
			error.SetError(e)
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, false, "")
		x, y := maths_imaginaires.CalcVar(par, Vars)
		Vars.Table["?"] = &types.Imaginaire{ x, y }
		str_ret = "?"
		t = 0
	} else if parser.IsFunc(str[0], 0) == 1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""))
		e = error.In(data, 1, str[0])
		if e != "1" {
			error.SetError(e)
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, true, str[0])
		res := maths_functions.Init(par, str[0], Vars)
		Vars.Table[str[0]] = &types.Fonction{ res }
		t = 0
	} else if strings.Index(str[1], "i") != -1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""))
		e = error.In(data, 0, "")
		if e != "1" {
			error.SetError(e)
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, false, "")
		x, y := maths_imaginaires.CalcVar(par, Vars)
		Vars.Table[str[0]] = &types.Imaginaire{ x, y }
		t = 0
	} else {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""))
		e = error.In(data, 0, "")
		if e != "1" {
			error.SetError(e)
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, false, "")
		x, _ := maths_imaginaires.CalcVar(par, Vars)
		Vars.Table[str[0]] = &types.Rationel{ x }
		t = 0
	}
	/*else if strings.Index(str[0], "mat") != -1 || strings.Index(str[0], "var") != -1 {
		
		fmt.Println("matrice")
	}*/

	return 1, t, str_ret
}