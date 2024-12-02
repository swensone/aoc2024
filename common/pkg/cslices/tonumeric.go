package cslices

import "strconv"

// ToIntSlice converts a slice of strings to a slice of integers
func ToIntSlice(s []string) ([]int, error) {
	var numbers []int

	for _, n := range s {
		num, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}

// ToFloat64Slice converts a slice of strings to a slice of float64s
func ToFloat64Slice(s []string) ([]float64, error) {
	var numbers []float64

	for _, n := range s {
		num, err := strconv.ParseFloat(n, 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}

	return numbers, nil
}
