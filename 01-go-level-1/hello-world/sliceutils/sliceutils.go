package sliceutils

func reverseSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func removeDublicates(s []int) []int {
	keys := make(map[int]bool)
	newSlice := []int{}
	for _, el := range s {
		if _, val := keys[el]; !val {
			newSlice = append(newSlice, el)
			continue
		}
		keys[el] = true
	}
	return newSlice
}
