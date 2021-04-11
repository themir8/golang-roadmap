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

func dublicate(slice []int) map[int]int {
	key := make(map[int]bool)
	key2 := make(map[int]int)

	for i := 0; i < len(slice); i++ {
		if key[slice[i]] {
			key2[slice[i]] = slice[i]
		} else {
			key[slice[i]] = true
		}
	}

	return key2
}

func main() {
	intSlice := []int{1, 1, 30, 23, 23, 5, 7, 7, 7}
	fmt.Println(intSlice)
	uniqueSlice := dublicate(intSlice)
	fmt.Println(uniqueSlice)
}
