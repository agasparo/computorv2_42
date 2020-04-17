package test

import (
	"os/exec"
	"fmt"
	"log"
	"bytes"
	"strings"
	"github.com/fatih/color"
)

type Cat struct {

	Name 	string
	tab		map[int]Testes
	Length	int
}

type Testes struct {
	
	Name		string
	Input		string
	Outpout		string
}

func DefineAndRun() {

	Syn := Cat{ "Syntaxe" ,  Syntaxe(), 0 }
	Syn.Length = len(Syn.tab)
	Run(Syn.tab, Syn)

	Vars := Cat{ "Variables" ,  Vars(), 0 }
	Vars.Length = len(Vars.tab)
	Run(Vars.tab, Vars)

	Func := Cat{ "Fonctions", Functions(), 0}
	Func.Length = len(Func.tab)
	Run(Func.tab, Func)

	Func_usu := Cat{ "Fonctions usuelles", Functions_usuelles(), 0}
	Func_usu.Length = len(Func_usu.tab)
	Run(Func_usu.tab, Func_usu)
}

func Run(table map[int]Testes, C Cat) {

	color.Magenta("Categorie : %s : (%d teste(s))\n", C.Name, C.Length)
	fmt.Println("")
	c := 0

	for i := 0; i < len(table); i++ {

		cmd := exec.Command("go", "run", "/Users/arthur/Desktop/42/computorv2/main.go", table[i].Input)
		var outb, errb bytes.Buffer
		cmd.Stdout = &outb
		cmd.Stderr = &errb
		err := cmd.Run()
		if err != nil {
		    log.Fatal(err)
		}
		fmt.Printf("test %d [%s (%s)] : ", i, table[i].Name, table[i].Input)
		if CheckRes(outb.String(), table[i].Outpout) {
			color.Set(color.FgGreen)
			fmt.Println(" [ OK ]")
			color.Unset()
		} else {
			color.Set(color.FgRed)
			fmt.Println(" [ KO ]")
			color.Unset()
			fmt.Println("Return : ")
			fmt.Println(outb.String())
			fmt.Printf("You must return : \n%s\n", table[i].Outpout)
			color.Magenta("Success : %d / %d", c, C.Length)
			fmt.Println("")
			return
		}
		c++
	}
	color.Magenta("Success : %d / %d", c, C.Length)
	fmt.Println("")
}

func CheckRes(Outpout string, attOuput string) (bool) {

	re := strings.Index(Outpout, attOuput)
	if re != -1 {
		return (true)
	}
	return (false)
}

func Syntaxe() (map[int]Testes) {

	Tvars := make(map[int]Testes)
	Tvars[0] = Testes{ "teste du =  : 1", "x= 2 * 3", "6" }
	Tvars[1] = Testes{ "teste du =  : 2", "x =2 * 3", "6" }
	Tvars[2] = Testes{ "teste du =  : 3", "x=2 * 3", "6" }
	Tvars[3] = Testes{ "teste du =  : 4", "x = 2 * 3", "6" }
	Tvars[4] = Testes{ "teste du =  : 5", "x 2 * 3", "You must have just one =" }
	Tvars[5] = Testes{ "teste du =  : 6", "x == 2 * 3", "You must have just one =" }
	Tvars[6] = Testes{ "teste du =  : 7", "x = 2 * 3 = 4", "You must have just one =" }

	Tvars[7] = Testes{ "teste avec des lettres  : 1", "x = 23a", "'23a' isn't a number" }
	Tvars[8] = Testes{ "teste avec des lettres  : 2", "f(x) = 23a", "in your function (or not an other unknown)" }
	Tvars[9] = Testes{ "teste avec des lettres  : 3", "f(x) = a23", "in your function (or not an other unknown)" }
	Tvars[10] = Testes{ "teste avec des lettres  : 4", "f(a) = 23a", "23 * a" }
	Tvars[11] = Testes{ "teste avec des lettres  : 5", "f(a) = a23", "a * 23" }
	Tvars[12] = Testes{ "teste avec des lettres  : 6", "f(a) = 23a23", "'23a23' isn't a number" }
	Tvars[13] = Testes{ "teste avec des lettres  : 7", "c = 23i", "0.000000 + 23.000000i" }
	Tvars[14] = Testes{ "teste avec des lettres  : 8", "c = i23i", "'i23i' isn't a number" }
	Tvars[15] = Testes{ "teste avec des lettres  : 9", "x = a23", "'a23' isn't a number" }
	Tvars[16] = Testes{ "teste avec des lettres  : 9", "x = x23", "'x23' isn't a number" }

	Tvars[17] = Testes{ "teste avec les signes : 1", "x = 2 * 3", "6" }
	Tvars[18] = Testes{ "teste avec les signes : 2", "x = 2 * +3", "6" }
	Tvars[19] = Testes{ "teste avec les signes : 3", "x = +2 * +3", "6" }
	Tvars[20] = Testes{ "teste avec les signes : 4", "x = -2 * 3", "-6" }
	Tvars[21] = Testes{ "teste avec les signes : 5", "x = 2 * -3", "-6" }
	Tvars[22] = Testes{ "teste avec les signes : 6", "x = -2 * -3", "6" }
	Tvars[23] = Testes{ "teste avec les signes : 7", "x = 2 */ 3", "You have a mistake with your sign" }
	Tvars[24] = Testes{ "teste avec les signes : 8", "x = 2 /* 3", "You have a mistake with your sign" }
	Tvars[25] = Testes{ "teste avec les signes : 9", "x = 2 - 3", "-1" }
	Tvars[26] = Testes{ "teste avec les signes : 10", "x = 2 - -3", "5" }
	Tvars[27] = Testes{ "teste avec les signes : 11", "x = 2 - +3", "-1" }
	Tvars[28] = Testes{ "teste avec les signes : 12", "x = -2 - 3", "-5" }
	Tvars[29] = Testes{ "teste avec les signes : 13", "x = -2 - -3", "1" }
	Tvars[30] = Testes{ "teste avec les signes : 14", "x = 2 + 3", "5" }
	Tvars[31] = Testes{ "teste avec les signes : 15", "x = +2 + 3", "5" }
	Tvars[32] = Testes{ "teste avec les signes : 16", "x = +2 + +3", "5" }

	Tvars[33] = Testes{ "teste avec les signes : 17", "x = +-2", "You have a mistake with your sign" }
	Tvars[34] = Testes{ "teste avec les signes : 18", "x = +/2", "You have a mistake with your sign" }
	Tvars[35] = Testes{ "teste avec les signes : 19", "x = +*2", "You have a mistake with your sign" }
	Tvars[36] = Testes{ "teste avec les signes : 20", "x = +%2", "You have a mistake with your sign" }
	Tvars[37] = Testes{ "teste avec les signes : 21", "x = +%2", "You have a mistake with your sign" }
	Tvars[38] = Testes{ "teste avec les signes : 22", "x = -+2", "You have a mistake with your sign" }
	Tvars[39] = Testes{ "teste avec les signes : 23", "x = +*2", "You have a mistake with your sign" }
	Tvars[40] = Testes{ "teste avec les signes : 24", "x = -/2", "You have a mistake with your sign" }
	Tvars[41] = Testes{ "teste avec les signes : 25", "x = -%2", "You have a mistake with your sign" }
	Tvars[42] = Testes{ "teste avec les signes : 26", "x = *%2", "You have a mistake with your sign" }

	Tvars[43] = Testes{ "teste avec i : 1", "x = 2 * i", "0.000000 + 2.000000i" }
	Tvars[44] = Testes{ "teste avec i : 2", "x = 2i", "0.000000 + 2.000000i" }
	Tvars[45] = Testes{ "teste avec i : 3", "x = i2", "0.000000 + 2.000000i" }
	Tvars[46] = Testes{ "teste avec i : 4", "x = ia2", "'a2' isn't a number" }
	Tvars[47] = Testes{ "teste avec i : 5", "x = 2ai", "'2a' isn't a number" }
	Tvars[48] = Testes{ "teste avec i : 6", "x = i2i", "'i2i' isn't a number" }

	return (Tvars)
}

func Vars() (map[int]Testes) {

	Tvars := make(map[int]Testes)
	Tvars[0] = Testes{ "nom variable -> i", "i = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[1] = Testes{ "nom variable -> a+", "a+ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[2] = Testes{ "nom variable -> a-", "a- = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[3] = Testes{ "nom variable -> a*", "a* = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[4] = Testes{ "nom variable -> a/", "a/ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[5] = Testes{ "nom variable -> a%", "a% = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[6] = Testes{ "nom variable -> -", "- = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[7] = Testes{ "nom variable -> +", "+ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[8] = Testes{ "nom variable -> *", "* = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[9] = Testes{ "nom variable -> a[", "a[ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[10] = Testes{ "nom variable -> a]", "a] = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[11] = Testes{ "nom variable -> ()", "() = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[12] = Testes{ "nom variable -> (", "( = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[13] = Testes{ "nom variable -> )", ") = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[14] = Testes{ "nom variable -> [", "[ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[15] = Testes{ "nom variable -> a]", "a] = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[16] = Testes{ "nom variable -> ]", "] = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[17] = Testes{ "nom variable -> []", "[] = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[18] = Testes{ "nom variable -> a[]", "a[] = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[19] = Testes{ "nom variable -> @", "@ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[20] = Testes{ "nom variable -> !", "! = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[21] = Testes{ "nom variable -> ?", "? = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[22] = Testes{ "nom variable -> ˆ", "ˆ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[23] = Testes{ "nom variable -> |", "| = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[24] = Testes{ "nom variable -> \\", "\\ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[25] = Testes{ "nom variable -> .", ". = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[26] = Testes{ "nom variable -> ,", ", = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[27] = Testes{ "nom variable -> ˜", "˜ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[28] = Testes{ "nom variable -> `", "` = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[29] = Testes{ "nom variable -> ^", "^ = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[30] = Testes{ "nom variable -> Inf", "Inf = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[31] = Testes{ "nom variable -> +Inf", "+Inf = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[32] = Testes{ "nom variable -> -Inf", "-Inf = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[33] = Testes{ "nom variable -> NaN", "NaN = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[34] = Testes{ "nom variable -> 0", "0 = 2 * 3", "Your var must be just with alpha caracteres and not i" }
	Tvars[35] = Testes{ "nom variable -> a0", "a0 = 2 * 3", "Your var must be just with alpha caracteres and not i" }

	Tvars[36] = Testes{ "x = g", "x = g", "isn't a number" }
	Tvars[37] = Testes{ "3 = 4", "3 = 4", "Your var must be just with alpha caracteres and not i" }
	return (Tvars)
}

func Functions() (map[int]Testes) {

	Tvars := make(map[int]Testes)
	Tvars[0] = Testes{ "teste avec les fonctions (syntaxe) : 1", "f(x = 3 + 2x", "Your var must be just with alpha caracteres and not i" }
	Tvars[1] = Testes{ "teste avec les fonctions (syntaxe): 2", "fx) = 3 + 2x", "'2x' isn't a number" }
	Tvars[2] = Testes{ "teste avec les fonctions (syntaxe): 3", "f() = 3 + 2x", "You must have an unknown" }
	Tvars[3] = Testes{ "teste avec les fonctions (syntaxe): 4", "f(x) = 3 + 2", "in your function (or not an other unknown)" }
	Tvars[4] = Testes{ "teste avec les fonctions (syntaxe): 5", "f(x) = 3 + 2y", "in your function (or not an other unknown)" }
	Tvars[5] = Testes{ "teste avec les fonctions (syntaxe): 6", "f(x) = 3x + 2y", "isn't a number" }
	Tvars[6] = Testes{ "teste avec les fonctions (syntaxe): 7", "f(x) = 3x + 2", "3 * x + 2" }
	Tvars[7] = Testes{ "teste avec les fonctions (syntaxe): 8", "f(x) = x", "x" }
	Tvars[8] = Testes{ "teste avec les fonctions (syntaxe): 9", "f(x) = ", "in your function (or not an other unknown)" }
	//Tvars[9] = Testes{ "teste avec les fonctions (syntaxe): 10", "f(x) = xx", "jsp" }
	//teste de remplacer les fonctions usuelles
	return (Tvars)
}

func Functions_usuelles() (map[int]Testes) {

	Tvars := make(map[int]Testes)
	Tvars[0] = Testes{ "abs function : 1 abs(1)", "c = abs(1)", "1" }
	Tvars[1] = Testes{ "abs function : 2 abs(-1)", "c = abs(-1)", "1" }
	Tvars[2] = Testes{ "abs function : 3 abs(-0)", "c = abs(-0)", "0" }
	Tvars[3] = Testes{ "abs function : 4 abs(0)", "c = abs(0)", "0" }

	Tvars[4] = Testes{ "racinne carre function : 1 v(0)", "c = v(0)", "0" }
	Tvars[5] = Testes{ "racinne carre function : 2 v(1)", "c = v(1)", "1" }
	Tvars[6] = Testes{ "racinne carre function : 3 v(-1)", "c = v(-1)", "Impossible v(x) : [0; +Inf]" }

	Tvars[7] = Testes{ "inverse function : 1 inv(0)", "c = inv(0)", "Impossible inv(x) : [-Inf; 0[ U ]0; +Inf]" }
	Tvars[8] = Testes{ "inverse function : 2 inv(15)", "c = inv(15)", "0.06666" }
	Tvars[9] = Testes{ "inverse function : 3 inv(-20)", "c = inv(-20)", "-0.05" }

	Tvars[10] = Testes{ "expo function : 1 exp(0)", "c = exp(0)", "1" }
	Tvars[11] = Testes{ "expo function : 2 exp(10)", "c = exp(10)", "22026.4657" }
	Tvars[12] = Testes{ "expo function : 3 exp(-2)", "c = exp(-2)", "0.1353" }
	return (Tvars)
}