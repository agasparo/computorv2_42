package input

import (
	"os"
	"bufio"
	"strings"
)

type Data struct {

	Input 	[]string
	Length	int
}

func ReadSTDIN(input *Data) {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	input.Input = strings.Split(strings.ReplaceAll(text, "\n", ""), " ")
	input.Length = len(input.Input)
}