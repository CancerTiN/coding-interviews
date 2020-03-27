package main

import (
	"fmt"
	"math/rand"
)

func swap(p1 *int, p2 *int) {
	*p1, *p2 = *p2, *p1
}

func partition(data []int, length int, start int, end int) (small int) {
	index := rand.Intn(end)
	swap(&data[index], &data[end])
	small = start - 1
	for index = start; index < end; index++ {
		if data[index] < data[end] {
			small++
			if small != index {
				swap(&data[index], &data[small])
			}
		}
	}
	small++
	swap(&data[small], &data[end])
	return
}

func quickSort(data []int, length int, start int, end int) {
	if start == end {
		return
	}
	index := partition(data, length, start, end)
	if index > start {
		quickSort(data, length, start, index-1)
	}
	if index < end {
		quickSort(data, length, index+1, end)
	}
}

func rotatedArrayMin(data []int) (min int) {
	l := 0
	r := len(data) - 1
	m := (l + r) / 2
	for l < r {
		if data[m-1] > data[m] {
			break
		}
		if data[m] < data[l] {
			r = m
		} else {
			l = m
		}
		m = (l + r) / 2
	}
	min = data[m]
	return
}

func main() {
	slice := []int{3, 4, 5, 1, 2}
	fmt.Println(rotatedArrayMin(slice))
}
