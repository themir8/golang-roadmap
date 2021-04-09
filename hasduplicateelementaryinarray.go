package main

import (
	"fmt"
)

func unique(intSlice []int) []int {
	kalitlar := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := kalitlar[entry]; !value {
			kalitlar[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func main() {
	intSlice := []int{1, 1, 30, 23, 23, 5, 7, 7}
	fmt.Println(intSlice)
	uniqueSlice := unique(intSlice)
	fmt.Println(uniqueSlice)
}
