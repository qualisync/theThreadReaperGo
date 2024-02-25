package goroutines

import (
	"sync"
)

func ProcessNumberList() []int {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	doubleNumbers := make([]int, len(numbers))
	var lock sync.Mutex

	var wg sync.WaitGroup
	for index, number := range numbers {
		wg.Add(1)
		go doubleNumberToArray(&wg, &lock, &doubleNumbers[index], number)
	}

	for i := 0; i < len(numbers); i++ {
	}
	wg.Wait() // Wait for the goroutines to finish

	return doubleNumbers
}

func doubleNumberToArray(wg *sync.WaitGroup, lock *sync.Mutex, destination *int, number int) {
	defer wg.Done()
	lock.Lock()
	*destination = number * 2
	lock.Unlock()
}
