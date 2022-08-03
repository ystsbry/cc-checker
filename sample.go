package main

func BubbleSort(l []int) []int {
	for i := 0; i < len(l)-1; i++ {
		for j := 0; j < len(l); j++ {
			if l[i] > l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
	return l
}