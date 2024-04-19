package sort

func Quick[T any](a []T, lo, hi int, lt func(a, b T) bool) {
	if lo >= hi {
		return
	}

	p := partition(a, lo, hi, lt)
	Quick(a, lo, p, lt)
	Quick(a, p+1, hi, lt)
}

func partition[T any](a []T, lo, hi int, lt func(a, b T) bool) int {
	pivot := a[lo]
	i := lo + 1
	j := hi

	for {
		for i <= j && lt(a[i], pivot) {
			i++
		}
		for j >= i && !(lt(a[j], pivot)) {
			j--
		}
		if i >= j {
			break
		}
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
	a[lo], a[j] = a[j], a[lo]
	return j
}
