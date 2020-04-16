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
	if Err(0, error.Syntaxe(tmp), true, "1") {
		return 0, 0, ""
	}

	str[0] = strings.ToLower((strings.Trim(str[0], " ")))
	str[1] = strings.Trim(str[1], " ")
	str_ret := str[0]
	err_pars := 0

	if str[1] == "?" { // cas particulier ppur check les variables
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[0]), " ", ""), &err_pars)
		if Err(err_pars, error.In(data, 0, ""), true, "1") {
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, false, "")
		x, y := maths_imaginaires.CalcVar(par, Vars)
		Vars.Table["?"] = &types.Imaginaire{ x, y }
		str_ret = "?"
		t = 0
	} else if parser.IsFunc(str[0], 0) == 1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.Checkfuncx(str[0], str[1]), error.Checkfuncpa(str[0]), error.In(data, 1, str[0])) {
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, true, str[0])
		res := maths_functions.Init(par, str[0], Vars)
		Vars.Table[str[0]] = &types.Fonction{ res }
		t = 0
	} else if strings.Index(str[1], "i") != -1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.In(data, 0, ""), error.Checkvars(str[0]), "1") {
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		par := parentheses.Parse(data, Vars, false, "")
		x, y := maths_imaginaires.CalcVar(par, Vars)
		Vars.Table[str[0]] = &types.Imaginaire{ x, y }
		t = 0
	} else {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.In(data, 0, ""), error.Checkvars(str[0]), "1") {
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

func Err(err_parse int, e string, a bool, b string) (bool) {

	if e != "1" {
		error.SetError(e)
		return (true)
	}
	if err_parse == 1 {
		error.SetError("You have a mistake with your sign")
		return (true)
	}
	if !a {
		error.SetError("Your var must be just with alpha caracteres and not i")
		return (true)
	}
	if b != "1" {
		error.SetError(b)
		return (true)
	}
	return (false)
}