package main

import "fmt"

func main() {
	xi := []int{5, 1, 90, 8, -1, 38, -16}
	fmt.Println("Selection Sort:\n", sort(selectionSort, xi))
	fmt.Println("Bubble Sort:\n", sort(bubbleSort, xi))

}

func sort(f func([]int) []int, x []int) []int {
	output := f(x)
	return output
}
func swap(sli []int, x, y int) {
	sli[x], sli[y] = sli[y], sli[x]
}
func selectionSort(input []int) []int {
	s := make([]int, 0, len(input))
	s = append(s, input[:]...)

	for i := range s {
		minIndex := i

		for j, v := range s[i:] {
			if s[minIndex] > v {
				minIndex = j + i
			}
		}

		swap(s, i, minIndex)
	}
	return s
}
func bubbleSort(input []int) []int {
	s := make([]int, 0, len(input))
	s = append(s, input[:]...)

	for i := 0; i < len(s)-1; i++ {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				swap(s, j, j+1)
			}
		}
	}
	return s
}
