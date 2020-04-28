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
	"resolve"
	"equations"
	"maps"
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
	arg1 := ""

	Vars.Table = make(map[string]types.AllT)
	usuelles_functions.Init(&Vars)
	for i := 1; i == 1; i = 1 {
		input.ReadSTDIN(&Inputs)
		if Inputs.Input[0] == "exit" {
			fmt.Println("bye")
			return
		}
		if Inputs.Length >= 2 {
    		arg = Inputs.Input[1]
		}
		if Inputs.Length == 3 {
    		arg1 = Inputs.Input[2]
		}
		if commands.IsCommand(Inputs.Input[0], arg, arg1, Vars) != 1 {
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
	arg1 := ""

	Vars.Table = make(map[string]types.AllT)
	usuelles_functions.Init(&Vars)
	if Inputs.Input[0] == "exit" {
		fmt.Println("bye")
		return
	}
	if Inputs.Length == 2 {
   		arg = Inputs.Input[1]
	}
	if Inputs.Length == 3 {
    	arg1 = Inputs.Input[2]
	}
	if commands.IsCommand(Inputs.Input[0], arg, arg1, Vars) != 1 {
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
	Eq_Data := resolve.Unknown{}

	if strings.Index(str[1], "?") != -1 && strings.Count(str[1], "?") == 1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[0]), " ", ""), &err_pars)
		data_r := parser.GetAllIma(strings.ReplaceAll(strings.ReplaceAll(strings.ToLower(str[1]), "?", ""), " ", ""), &err_pars)
		data_r = maps.Reindex(data_r)
		data = maps.Reindex(data)
		Eq_Data.Part1 = data
		Eq_Data.Part2 = data_r
		if Err(err_pars, error.In(data, 0, "", Dat), true, "1") {
			return 0, 0, ""
		}
		if Err(err_pars, error.In(data_r, 0, "", Dat), true, "1") {
			return 0, 0, ""
		}
		data = parser.Checkfunc(data, Dat)
		if strings.Index(data[0], "Impossible") != -1 || strings.Index(data[0], "for unknown not an expression") != -1 {
			error.SetError(data[0])
			return 1, -1, str_ret
		}
		data_r = parser.Checkfunc(data_r, Dat)
		if strings.Index(data_r[0], "Impossible") != -1 || strings.Index(data_r[0], "for unknown not an expression") != -1 {
			error.SetError(data_r[0])
			return 1, -1, str_ret
		}
		if len(data) == 1 {
			data = parser.GetAllIma(strings.ReplaceAll(strings.ToLower(maps.Join(data, "")), " ", ""), &err_pars)
			/*if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
				return 0, 0, ""
			}*/
		}
		if len(data_r) == 1 {
			data_r = parser.GetAllIma(strings.ReplaceAll(strings.ToLower(maps.Join(data_r, "")), " ", ""), &err_pars)
			/*if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
				return 0, 0, ""
			}*/
		}
		if data_r[0] != "" {
			if !resolve.IsEquation(&Eq_Data, Dat, 0) || !resolve.IsEquation(&Eq_Data, Dat, 1) {
				error.SetError("This equation isn't soluble")
				return 1, -1, str_ret
			}
			if !resolve.IsSoluble(Eq_Data) {
				error.SetError("This equation isn't soluble")
				return 1, -1, str_ret
			}
			response := resolve.Init(&Eq_Data, Dat)
			if strings.Index(response, "|") == -1 {
				error.SetError(response)
				return 1, -1, str_ret
			}
			deg, delta, sol := equations.Resolve(Eq_Data.Eqs)
			Vars.Table["?"] = &types.EquaSol{ deg, delta, sol }
			str_ret = "?"
			t = 0
		} else {
			// test is defined
			par := parentheses.Parse(data, Vars, false, "")
			if strings.Index(par[0], "by 0") != -1 {
				error.SetError(par[0])
				return 1, -1, str_ret
			}
			x, y, err := maths_imaginaires.CalcVar(par, Vars)
			if err != "" {
				error.SetError(err)
				return 1, -1, str_ret
			}
			Vars.Table["?"] = &types.Imaginaire{ x, y }
			str_ret = "?"
			t = 0
		}
	} else if parser.IsFunc(str[0], 0) == 1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.Checkfuncx(str[0], str[1], Dat), error.Checkfuncpa(str[0]), error.In(data, 1, str[0], Dat)) {
			return 0, 0, ""
		}
		data = maps.Reindex(data)
		data = parser.Checkfunc(data, Dat)
		if strings.Index(data[0], "Impossible") != -1 || strings.Index(data[0], "for unknown not an expression") != -1 {
			error.SetError(data[0])
			return 1, -1, str_ret
		}
		res := maths_functions.Init(data, str[0], Vars)
		Vars.Table[str[0]] = &types.Fonction{ res }
		t = 0
	} else if strings.Index(str[1], "i") != -1 {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
			return 0, 0, ""
		}
		data = maps.Reindex(data)
		if !Function_var(data, Dat) {
			error.SetError("variable can't be equal to a function")
			return 1, -1, str_ret
		}
		data = parser.Checkfunc(data, Dat)
		if strings.Index(data[0], "Impossible") != -1 || strings.Index(data[0], "for unknown not an expression") != -1 {
			error.SetError(data[0])
			return 1, -1, str_ret
		}
		if len(data) == 1 {
			data = parser.GetAllIma(strings.ReplaceAll(strings.ToLower(maps.Join(data, "")), " ", ""), &err_pars)
			/*if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
				return 0, 0, ""
			}*/
		}
		par := parentheses.Parse(data, Vars, false, "")
		if strings.Index(par[0], "by 0") != -1 {
			error.SetError(par[0])
			return 1, -1, str_ret
		}
		x, y, err := maths_imaginaires.CalcVar(par, Vars)
		if err != "" {
			error.SetError(err)
			return 1, -1, str_ret
		}
		Vars.Table[str[0]] = &types.Imaginaire{ x, y }
		t = 0
	} else {
		data := parser.GetAllIma(strings.ReplaceAll(strings.ToLower(str[1]), " ", ""), &err_pars)
		if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
			return 0, 0, ""
		}
		data = maps.Reindex(data)
		if !Function_var(data, Dat) {
			error.SetError("variable can't be equal to a function")
			return 1, -1, str_ret
		}
		data = parser.Checkfunc(data, Dat)
		if strings.Index(data[0], "Impossible") != -1 || strings.Index(data[0], "for unknown not an expression") != -1 {
			error.SetError(data[0])
			return 1, -1, str_ret
		}
		if len(data) == 1 {
			data = parser.GetAllIma(strings.ReplaceAll(strings.ToLower(maps.Join(data, "")), " ", ""), &err_pars)
			/*if Err(err_pars, error.In(data, 0, "", Dat), error.Checkvars(str[0]), "1") {
				return 0, 0, ""
			}*/
		}
		par := parentheses.Parse(data, Vars, false, "")
		if strings.Index(par[0], "by 0") != -1 {
			error.SetError(par[0])
			return 1, -1, str_ret
		}
		x, y, err := maths_imaginaires.CalcVar(par, Vars)
		if err != "" {
			error.SetError(err)
			return 1, -1, str_ret
		}
		if y > 0 {
			Vars.Table[str[0]] = &types.Imaginaire{ x, y }
		} else {
			Vars.Table[str[0]] = &types.Rationel{ x }
		}
		t = 0
	}
	/*else if strings.Index(str[0], "mat") != -1 || strings.Index(str[0], "var") != -1 {
		
		fmt.Println("matrice")
	}*/

	return 1, t, str_ret
}

func Function_var(data map[int] string, Dat types.Variable) (bool) {

	inter := 0

	for i := 0; i < len(data); i++ {

		if parser.IsFunc(data[i], 1) == 1 || parser.IsFunc(data[i], 0) == 1 {

			inter++
			p1 := strings.Index(data[i], "(")
			p2 := strings.Index(data[i], ")")
			nstr := data[i][p1 + 1:p2]
			if parser.IsNumeric(nstr) || error.Is_defined(nstr, Dat) {
				return (true)
			}
		}
	}
	if inter > 0 {
		return (false)
	}
	return (true)
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