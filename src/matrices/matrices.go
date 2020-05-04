package matrices

import (
	"parentheses"
	"strings"
	//"fmt"
)

func Parse(str string) (string) {

	index_d := parentheses.IndexString(str, "[")
	fin_c := strings.Index(str, "]")
	
	if fin_c != -1 && fin_c < index_d {
		return ("you have a problem with your matrices syntaxe")
	}
	return (str)
}