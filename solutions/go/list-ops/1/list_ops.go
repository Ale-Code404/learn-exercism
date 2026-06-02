package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	acc := initial

	for i := 0; i < len(s); i++ {
		acc = fn(acc, s[i])
	}

	return acc
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	acc := initial

	for i := len(s) - 1; i >= 0; i-- {
		acc = fn(s[i], acc)
	}

	return acc
}

func (s IntList) Filter(fn func(int) bool) IntList {
	var filtered IntList

	for _, i := range s {
		if fn(i) {
			filtered = append(filtered, i)
		}
	}

	return filtered
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	var transformed IntList

	for _, i := range s {
		transformed = append(transformed, fn(i))
	}

	return transformed
}

func (s IntList) Reverse() IntList {
	var reversed IntList

	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	return reversed
}

func (s IntList) Append(lst IntList) IntList {
	return append(s, lst...)
}

func (s IntList) Concat(lists []IntList) IntList {
	var flatten IntList

	for _, list := range lists {
		flatten = flatten.Append(list)
	}

	return flatten
}
