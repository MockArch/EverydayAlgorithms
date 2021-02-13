package listops

// IntList is a list of ints
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Length returns the number of items in the list
func (l IntList) Length() int {
	i := 0
	for range l {
		i++
	}
	return i
}

// Foldl folds each item into the accumulator from the left
func (l IntList) Foldl(fn binFunc, initial int) int {
	n := initial
	for _, i := range l {
		n = fn(n, i)
	}
	return n
}

// Foldr folds each item into the accumulator from the right
func (l IntList) Foldr(fn binFunc, initial int) int {
	n := initial
	for i := l.Length() - 1; i >= 0; i-- {
		n = fn(l[i], n)
	}
	return n
}

// Append another list to the list
func (l IntList) Append(appendList IntList) IntList {
	c := l.Length()
	list := make(IntList, c+appendList.Length())
	for i, n := range l {
		list[i] = n
	}
	for i, n := range appendList {
		list[c+i] = n
	}
	return list
}

// Filter returns a subset of the list as a new list filtered by a pred Function
func (l IntList) Filter(fn predFunc) IntList {
	list := IntList{}
	for _, n := range l {
		if fn(n) {
			list = list.Append([]int{n})
		}
	}
	return list
}

// Map applies a function to each element in the list
func (l IntList) Map(fn unaryFunc) IntList {
	list := make(IntList, l.Length())
	for i, n := range l {
		list[i] = fn(n)
	}
	return list
}

// Reverse the order of the list
func (l IntList) Reverse() IntList {
	c := l.Length()
	list := make(IntList, c)
	for i := range l {
		list[c-i-1] = l[i]
	}
	return list
}

// Concat many lists together
func (l IntList) Concat(lists []IntList) IntList {
	list := l
	for _, nextList := range lists {
		list = list.Append(nextList)
	}
	return list
}
