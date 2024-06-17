package utils

func IfElementInSlice(slice *[]string, element *string) int {
	for i, v := range *slice {
		if v == *element {
			return i
		}
	}
	return -1
}
