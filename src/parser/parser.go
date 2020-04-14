package parser

import (
	"strconv"
	"strings"
	"fmt"
	"unicode"
	"maths_functions"
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

func Checkfunc(data map[int]string, Vars *types.Variable) (map[int]string) {

	for i := 0; i < len(data); i++ {

		if IsFunc(data[i]) == 1 {
			name, value := maths_functions.GetDataFunc(data[i], Vars)
			fmt.Println(name)
			fmt.Println(value)
		}
	}
	return (data)
}

func IsFunc(str string) (int) {

	p1 := strings.Index(str, "(")
	p2 := strings.Index(str, ")")

	if !IsLetter(str[0:p1]) {
		return (0)
	}

	if !IsLetter(str[p1 + 1:p2]) {
		return (0)
	}

	if p1 != -1 && p2 != -1 && p1 < p2 {

		return (1)
	}
	return (0)
}

func IsLetter(s string) bool {

    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
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
		index := GetCararc(str, "+-/*%")
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
	return (data)
}