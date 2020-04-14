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