package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	numbers := createRandomSlice()

	fmt.Println(numbers)

	for _, value := range numbers {
		if value%2 == 0 {
			fmt.Println("Even")
		} else {
			fmt.Println("Odd")
		}
	}
}

func createRandomSlice() []int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	var randomSlice []int
	n := 0

	for n < 10 {
		randomSlice = append(randomSlice, r.Intn(9))
		n++
	}
	return randomSlice
}
