package sliceutils

// ReverseSlice func
func ReverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// RemoveDublicates func
func RemoveDublicates(s []int) []int {
	keys := make(map[int]bool)
	newSlice := []int{}
	for _, el := range s {
		if _, val := keys[el]; !val {
			newSlice = append(newSlice, el)
			keys[el] = true
		}
	}
	return newSlice
}
