package main

import (
	"fmt"
	"sync"
)

/*
Написать программу, которая конкурентно
рассчитает значение квадратов чисел взятых
из массива (2,4,6,8,10) и выведет их квадраты в stdout.
*/

func computeSquares(nums []int) {
	var wg sync.WaitGroup // Используется для ожидания завершения всех горутин перед завершением программы.

	for _, num := range nums {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Println(n * n)
		}(num)
	}

	wg.Wait()
}

func main() {
	sample := []int{2, 4, 6, 8, 10}
	computeSquares(sample)
}
