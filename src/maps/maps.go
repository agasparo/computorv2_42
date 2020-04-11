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