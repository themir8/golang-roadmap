package main

import "fmt"

func main() {
	fmt.Println(palindrome("mirsaid diasrim"))          // true
	fmt.Println(palindrome("aziza aziza"))              // true
	fmt.Println(palindrome("palindrome palindromemas")) // false
}

func palindrome(input string) bool {
	for i := 0; i < len(input)/2; i++ {
		if input[i] != input[len(input)-i-1] {
			return false
		}
	}
	return true
}
