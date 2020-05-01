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
	input.Input = strings.Split(ReplaceWhiteSpace(text), " ")
	input.Length = len(input.Input)
}

func ReplaceWhiteSpace(text string) (string) {

	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\r\n", "")
	text = strings.ReplaceAll(text, "\f", "")
	text = strings.ReplaceAll(text, "\v", "")
	return (text)
}