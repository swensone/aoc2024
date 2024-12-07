package cslices

// RemoveAt returns a new slice with the element at index i removed
func RemoveAt(s []int, i int) []int {
	data := DeepCopy(s)
	return append(data[:i], data[i+1:]...)
}

// RemoveElement returns a new slice with the element S removed
func RemoveElement(s []string, e string) []string {
	res := []string{}
	for _, v := range s {
		if v != e {
			res = append(res, v)
		}
	}
	return res
}