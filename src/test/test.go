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
	//Vars := Cat{ "Variables" ,  Vars(), 0 }
	//Vars.Length = len(Vars.tab)
	//Run(Vars.tab, Vars)
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
	return (Tvars)
}

func Vars() (map[int]Testes) {

	Tvars := make(map[int]Testes)
	Tvars[0] = Testes{ "nom variable -> i", "i = 2 * 3", "cant have var with name 'i'" }
	return (Tvars)
}