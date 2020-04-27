package maps

import (
	"strings"
)

func MapSlice(data map[int]string, index int) (map[int]string) {

	for i := index; i < len(data); i++ {

		data[i] = data[i + 2]
		delete(data, i + 2);
	}
	return (data)
}

func MapSliceCount(data map[int]string, index int, add int) (map[int]string) {

	if add == 0 {
		return (data)
	}

	for i := index; i < len(data); i++ {

		data[i] = data[i + add]
		delete(data, i + add);
	}
	return (data)
}

func Array_search_count(array map[int]string, to_search string) (res int) {

	count := 0

	for i := 0; i < len(array); i++ {

		if strings.Index(array[i], to_search) != -1 {
			count++;
		}
	}
	return (count)
}

func CombineN(data map[int]string, in map[int]string, i int) (map[int]string) {

	tab := make(map[int]string)
	tab = Combine(data, tab, 0, i)

	add_in := strings.Count(data[i], "(")
	add_back := strings.Count(data[i], ")")
	in[0] = strings.Repeat("(", add_in) + in[0]
	in[len(in) - 1] = in[len(in) - 1] + strings.Repeat(")", add_back)
	
	tab = Combine(in, tab, 0, len(in))
	tab = Combine(data, tab, i + 1, len(data))
	return (tab)
}

func Combine(data map[int]string, tab map[int]string, min int, max int) (map[int]string) {

	for a := min; a < max; a++ {

		tab[len(tab)] = data[a]
	}
	return (tab)
}

func Join(data map[int]string) (str string) {

	for i := 0; i < len(data); i++ {
		str += data[i]
	}
	return (str)
}