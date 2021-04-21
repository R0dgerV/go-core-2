package main

import "fmt"
import "go-core-2/01-fib/pkg/fibonache"

func main() {
	numbers := []int{1, 7, 10, 16, 20}

	for _, num := range numbers {
		fmt.Println(fibonache.Number(num))
	}
}
