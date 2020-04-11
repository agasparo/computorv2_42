package parser

import (
	"strconv"
	"strings"
)

func Array_search_count(array []string, to_search string) (res int) {

	count := 0

	for i:= 0; i < len(array); i++ {

		if array[i] == to_search {
			count++;
		}
	}
	return (count)
}

func IsNumeric(s string) (bool) {

    _, err := strconv.ParseFloat(s, 64)
    return (err == nil)
}

func GetCararc(str string, c string) (int) {

	var index int

	for i := 0; i < len(c); i++ {

		index = strings.Index(str, string(c[i]))
		if index != -1 {
			return (index) 
		}
	}
	return (-1)
}

func GetSign(str string, index int) (string, int) {

	var add int = 0

	if str[index] == '+' || str[index] == '-' {

		add++
	}

	if str[index + 1] == '+' || str[index + 1] == '-' {

		add++
	} 

	if str[index] == '+' && str[index + 1] == '-' {

		return "-", add
	}

	if str[index] == '-' && str[index + 1] == '+' {

		return "-", add
	}

	if str[index] == '-' && str[index + 1] == '-' {

		return "+", add
	}

	if str[index] == '+' || str[index + 1] == '+' {

		return "+", add
	}

	return "+", add
}