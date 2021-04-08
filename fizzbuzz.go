package main

import "log"

func main() {
	massiv := [100]int{100}
	for x := range massiv {
		fizzbuzz(x)
	}
}

func fizzbuzz(num int) {
	if num/3 == 0 {
		log.Println("Fizz")
	}
	if num/5 == 0 {
		log.Println("Buzz")
	}
	if num/15 == 0 {
		log.Println("FizzBuzz")
	}
}
