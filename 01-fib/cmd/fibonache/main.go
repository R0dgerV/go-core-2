package main

import "fmt"
import "go-core-2/01-fib/pkg/fibonache"

func main() {
	numbers := []int{1, 7, 10, 16, 23, 20, 45}
	max := 20

	for _, num := range numbers {
		if num > max {
			fmt.Println(fmt.Sprintf("Ошибка!!! Максимальное допустимое число: %d, Ваше число: %d", max, num))
		}

		fmt.Println(fmt.Sprintf("Для числа %d, фибоначе = %d", num, fibonache.Number(num)))
	}
}
