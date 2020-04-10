package maths_imaginaires

import (
	"strings"
	"strconv"
)

func GetAll() {
	// verifier dans chauqe qu'il ny ai pas de lettrre a par i	
}

func ParseOne(str string) (x float64, y float64) {

	str_tmp := strings.ReplaceAll(str, "*", "")
	str_tmp = strings.ReplaceAll(str_tmp, " ", "")
	new_str := strings.Split(str_tmp, "+")

	if strings.Index(new_str[0], "i") != -1 {
		y, _ = strconv.ParseFloat(strings.ReplaceAll(new_str[0], "i", ""), 64)
		x, _ = strconv.ParseFloat(new_str[len(new_str) - 1], 64)
	} else {
		x, _ = strconv.ParseFloat(new_str[0], 64)
		y, _ = strconv.ParseFloat(strings.ReplaceAll(new_str[len(new_str) - 1], "i", ""), 64)
	}

	return x, y
}