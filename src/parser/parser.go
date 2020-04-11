package parser

import (
	"strconv"
	"strings"
	"fmt"
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

	var index, min int = -1, -1

	for i := 0; i < len(c); i++ {

		index = strings.Index(str, string(c[i]))
		if index != -1 && (index < min || min == -1) {
			min = index
		}
	}
	return (min)
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

	return string(str[index]), add
}

func GetAllIma(str string) (map[int]string) {

	data := make(map[int]string)
	var itab, neg int

	if str[0] == '+' {
		str = str[1:len(str)]
	}

	itab = 0
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			
			str = str[1:len(str)]
			neg = 1
		}
		index := GetCararc(str, "+-/*")
		if index == -1 {
			data[itab] = str
			return (data)
		}
		if neg == 1 {
			data[itab] = "-" + str[i:index]
		} else {
			data[itab] = str[i:index]
		}
		itab++
		neg = 0
		i = i + index
		data[itab] = string(str[i])
		itab++
		str = str[i + 1:len(str)]
		i = -1
	}
	fmt.Println(data)
	return (data)
}