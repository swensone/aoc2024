package cslices

// RemoveAt returns a new slice with the element at index i removed
func RemoveAt(s []int, i int) []int {
	data := DeepCopy(s)
	return append(data[:i], data[i+1:]...)
}
