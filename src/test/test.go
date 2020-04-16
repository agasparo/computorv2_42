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
}

func Run(table map[int]Testes, C Cat) {

	color.Magenta("Categorie : %s : (%d teste(s))\n", C.Name, C.Length)
	fmt.Println("")

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
			return
		}
	}
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
	Tvars[8] = Testes{ "teste avec des lettres  : 2", "f(x) = 23a", "'23a' isn't a number" }
	Tvars[9] = Testes{ "teste avec des lettres  : 3", "f(x) = a23", "'a23' isn't a number" }
	Tvars[10] = Testes{ "teste avec des lettres  : 4", "f(a) = 23a", "23 * a" }
	Tvars[11] = Testes{ "teste avec des lettres  : 5", "f(a) = a23", "a * 23" }
	Tvars[12] = Testes{ "teste avec des lettres  : 6", "f(a) = 23a23", "'23a23' isn't a number" }
	Tvars[13] = Testes{ "teste avec des lettres  : 7", "c = 23i", "23i" }
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
	return (Tvars)
}