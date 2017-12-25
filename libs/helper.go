package libs

import (
	"encoding/json"
	"fmt"
)

func InIntSlice(slice []int, n int) bool {
	for _, v := range slice {
		if v == n {
			return true
		}
	}
	return false
}

func AppendUniqueInt(ints []int, new_ints ...int) []int {
	for _, ni := range new_ints {
		in := false
		for _, ii := range ints {
			if ii == ni {
				in = true
				break
			}
		}
		if !in {
			ints = append(ints, ni)
		}
	}
	return ints
}

func JsonDecode(jsonStr string) interface{} {
	var data interface{}
	err := json.Unmarshal([]byte(jsonStr), &data)
	if err != nil {
		fmt.Println(err)
	}
	return data
}
