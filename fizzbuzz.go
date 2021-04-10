package main

import "fmt"

func main() {
	massiv := [100]int{100}
	for x := range massiv {
		fizzbuzz(x)
	}
	// fizzbuzz(15)
}

func fizzbuzz(num int) {
	if num%3 == 0 {
		fmt.Printf("Fizz")
	} else if num%5 == 0 {
		fmt.Printf("Buzz")
	} else if num%1
	fmt.Printf("\n")
}
