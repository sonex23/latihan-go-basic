package main

import "fmt"

func main() {
	// your code here
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("Buzz")
		}
		if i%3 == 0 {
			fmt.Printf("Fuzz")
		}
		if i%5 == 0 {
			fmt.Printf("BuzzFuzz")
		}
		fmt.Println("")
	}
}
