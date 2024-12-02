package cslices

// DeepCopy returns a copy of the slice
func DeepCopy(s []int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return clone
}
