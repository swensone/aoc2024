package elfsort

import "slices"

// ElfSort is a sorting algorithm based around the elven rules
type ElfSort struct {
	data map[string]bool
}

// New creates a new ElfSort
func New() *ElfSort {
	return &ElfSort{
		data: make(map[string]bool),
	}
}

// AddRule adds a rule to the sort
func (s *ElfSort) AddRule(rule string) {
	s.data[rule] = true
}

// Compare returns sort ordering for strings a and b
func (s *ElfSort) Compare(a, b string) int {
	order, found := s.data[a+"|"+b]
	if !found {
		return 0
	}

	if order {
		return -1
	}
	return 1
}

// Sort sorts the data
func (s *ElfSort) Sort(data []string) {
	slices.SortFunc(data, s.Compare)
}

// IsSorted checks if the data is sorted
func (s *ElfSort) IsSorted(data []string) bool {
	return slices.IsSortedFunc(data, s.Compare)
}
