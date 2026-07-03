package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	SIZE   = 100_000_000
	CHUNKS = 8
)

// generateRandomElements generates random elements.
func generateRandomElements(size int) []int {
	// ваш код здесь
	if size <= 0 {
		return []int{}
	}
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = rand.Intn(1000000) + 1
	}
	return data
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	max := data[0]
	for i := 1; i < len(data); i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	// ваш код здесь
	if len(data) == 0 {
		return 0
	}
	chunkSize := len(data) / CHUNKS
	if chunkSize == 0 {
		return maximum(data)
	}
	maxValues := make([]int, CHUNKS)

	var wg sync.WaitGroup
	for i := 0; i < CHUNKS; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			start := i * chunkSize
			end := start + chunkSize

			if i == CHUNKS-1 {
				end = len(data)
			}

			maxValues[i] = maximum(data[start:end])
		}(i)
	}
	wg.Wait()
	return maximum(maxValues)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	// ваш код здесь
	numbers := generateRandomElements(SIZE)
	fmt.Println("Поиск максимума в одном потоке")

	// ваш код здесь
	start := time.Now()
	max := maximum(numbers)
	elapsed := time.Since(start).Microseconds()
	fmt.Println("Максимум:", max)
	fmt.Println("Время:", elapsed, "мкс")
	fmt.Println()

	fmt.Printf("Поиск максимума в %d потоках\n", CHUNKS)

	// ваш код здесь
	start = time.Now()
	max = maxChunks(numbers)
	elapsed = time.Since(start).Microseconds()

	fmt.Println("Максимум:", max)
	fmt.Println("Время:", elapsed, "мкс")
}
