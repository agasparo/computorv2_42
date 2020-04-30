package input

import (
	"os"
	"bufio"
	"strings"
	//"fmt"
)

type Data struct {

	Input 	[]string
	Length	int
}

func ReadSTDIN(input *Data) {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input.Input = strings.Split(strings.ReplaceAll(strings.ReplaceAll(text, "\t", ""), "\n", ""), " ")
	input.Length = len(input.Input)
}