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
	var index, itab, neg int
	var tmp_str string

	if str[0] == '+' {
		str = str[1:len(str)]
	}

	itab = 0
	for i := 0; i < len(str); i++ {

		if str[i] == '-' {
			
			str = str[1:len(str)]
			neg = 1
		}
		index = GetCararc(str, "+-/*")
		if index == -1 {
			if strings.Index(str, "i") != -1 {
				str += "+0"
			} else {
				str += "+0i"
			}
			index = GetCararc(str, "+-/*")
		}
		if str[index] == '*' && str[index + 1] == 'i' {
			index += GetCararc(str[index + 1:len(str)], "+-/*") + 1
		}
		if str[index] == '/' {
			n1 := str[i:index]
			tmp := GetCararc(str[index + 1:len(str)], "+-/*")
			n2 := str[index + 1:tmp + index + 1]
			fmt.Printf("n1 : %s, n2: %s, tmp :%d\n", n1, n2, tmp)
			is_i := 0
			if strings.Index(n1, "i") != -1 || strings.Index(n2, "i") != -1 {
				n1 = strings.ReplaceAll(n1, "i", "")
				n2 = strings.ReplaceAll(n2, "i", "")
				is_i = 1
			}
			x1, _ := strconv.ParseFloat(n1, 64)
			y1, _ := strconv.ParseFloat(n2, 64)
			n3 := x1 / y1
			var n4 string
			if is_i == 1 {
				n4 = fmt.Sprintf("%fi", n3)
			} else {
				n4 = fmt.Sprintf("%f", n3)
			}
			str = n4 + str[index + tmp + i + 1:]
			index = GetCararc(str, "+-/*")
			i = -1
		} else {
			tmp_str = str[i:i + index]
			if neg == 1 {
				data[itab] = ("-" + tmp_str)
			} else {
				data[itab] = tmp_str
			}
			sign, add := GetSign(str, index)
			data[itab] += sign
			i = index + add
			index = GetCararc(str[i:len(str)], "+-/*")
			if str[index + i] == '*' && (str[index + 1 + i] == 'i' || str[index - 1 + i] == 'i') {
				add := GetCararc(str[index + 1 + i:len(str)], "+-/*")
				if add > -1 {
					index += add + 1
				} else {
					index = add
				}
			}
			if index == -1 {
				tmp_str = str[i:len(str)]
				data[itab] += tmp_str
				i = len(str)
			} else {
				tmp_str = str[i:i + index]
				data[itab] += tmp_str
				itab++
				data[itab] = string(str[i + index])
				itab++
				str = str[i + index + 1:len(str)]
				i = -1
			}
			neg = 0
		}
	}
	return (data)
}