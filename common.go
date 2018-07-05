package main

type Ints []int

func (i Ints) Len() int {
	return len(i)
}

func (i Ints) Swap(a, b int) {
	i[a], i[b] = i[b], i[a]
}

func (i Ints) Less(a, b int) bool {
	return i[a] < i[b]
}
