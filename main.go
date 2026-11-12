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
	if size <= 0 {
		return nil
	}

	elements := make([]int, size)
	for index := range elements {
		elements[index] = rand.Intn(SIZE) + 1
	}

	return elements
}

// maximum returns the maximum number of elements.
func maximum(data []int) int {
	if len(data) == 0 {
		return 0
	}

	max := data[0]
	for _, element := range data[1:] {
		if element > max {
			max = element
		}
	}

	return max
}

// maxChunks returns the maximum number of elements in a chunks.
func maxChunks(data []int) int {
	if len(data) == 0 {
		return 0
	}

	chunkSize := len(data) / CHUNKS
	chunkMaximums := make([]int, CHUNKS)
	var waitGroup sync.WaitGroup

	for chunkIndex := 0; chunkIndex < CHUNKS; chunkIndex++ {
		start := chunkIndex * chunkSize
		end := start + chunkSize
		if chunkIndex == CHUNKS-1 {
			end = len(data)
		}

		waitGroup.Add(1)
		go func(index int, chunk []int) {
			defer waitGroup.Done()

			if len(chunk) == 0 {
				chunkMaximums[index] = data[0]
				return
			}

			chunkMaximums[index] = maximum(chunk)
		}(chunkIndex, data[start:end])
	}

	waitGroup.Wait()
	return maximum(chunkMaximums)
}

func main() {
	fmt.Printf("Генерируем %d целых чисел\n", SIZE)
	data := generateRandomElements(SIZE)

	fmt.Println("Ищем максимальное значение в один поток")
	start := time.Now()
	max := maximum(data)
	elapsed := time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)

	fmt.Printf("Ищем максимальное значение в %d потоков\n", CHUNKS)
	start = time.Now()
	max = maxChunks(data)
	elapsed = time.Since(start).Microseconds()

	fmt.Printf("Максимальное значение элемента: %d\nВремя поиска: %d мкс\n", max, elapsed)
}
