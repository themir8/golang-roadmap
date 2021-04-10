package main

import "log"

func main() {
	massiv := [100]int{100}
	for x := range massiv {
		fizzbuzz(x)
	}
}

func fizzbuzz(num int) {
	if num % 15 == 0 {
		log.Println("FizzBuzz")
	} else if num % 5 == 0 {
		log.Println("Buzz")
	} else if num % 3 == 0 {
		log.Println("Fizz")
	}
}
