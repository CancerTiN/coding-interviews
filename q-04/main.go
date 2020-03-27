package main

import (
	"fmt"
)

func replaceBlankAlpha(s *string) {
	int32Slice := []int32(*s)
	length := len(int32Slice)
	for i := 0; i < length; i++ {
		if int32Slice[i] == ' ' {
			int32Slice = append(int32Slice[:i], append([]int32("%20"), int32Slice[i+1:]...)...)
			length = len(int32Slice)
		}
	}
	*s = string(int32Slice)
}

func main() {
	s := "We are happy."
	fmt.Println(s)
	replaceBlankAlpha(&s)
	fmt.Println(s)
}
