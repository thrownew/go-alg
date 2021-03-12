package sort

// BubbleSortSimple implementation
func BubbleSortSimple(l []int) {
	for i := range l {
		var change bool
		for j := len(l) - 1; j > i; j-- {
			if l[j-1] > l[j] {
				l[j], l[j-1] = l[j-1], l[j]
				change = true
			}
		}
		if !change {
			break
		}
	}
}
