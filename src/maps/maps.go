package maps

import (

)

func MapSlice(data map[int]string, index int) (map[int]string) {

	for i := index; i < len(data); i++ {

		data[i] = data[i + 2]
		delete(data, i + 2);
	}
	return (data)
}

func MapSliceCount(data map[int]string, index int, add int) (map[int]string) {

	for i := index; i < len(data); i++ {

		data[i] = data[i + add]
		delete(data, i + add);
	}
	return (data)
}

func MapReindex(data map[int]string) (map[int]string) {

	tab := make(map[int]string)
	c := 0

	for _, element := range data {
		tab[c] = element
		c++
	}
	return (tab)
}