package main

import "fmt"

func find(matrix [][]int, rows int, columns int, number int) (found bool) {
	r, c := 0, columns-1
	for r < rows && c >= 0 {
		if matrix[r][c] == number {
			found = true
			break
		} else if matrix[r][c] > number {
			c--
		} else {
			r++
		}
	}
	return
}

func main() {
	matrix := [][]int{{1, 2, 8, 9}, {2, 4, 9, 12}, {4, 7, 10, 13}, {6, 8, 11, 15}}
	fmt.Println(find(matrix, 4, 4, 7))
}
