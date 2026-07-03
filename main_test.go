package main

import (
	"testing"
)

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	result := generateRandomElements(0)
	if len(result) != 0 {
		t.Errorf("ожидается длина 0, но получена %d", len(result))
	}

	result = generateRandomElements(-5)
	if len(result) != 0 {
		t.Errorf("ожидается длина 0 для отрицательного размера, но получена %d", len(result))
	}

	size := 100
	result = generateRandomElements(size)
	if len(result) != size {
		t.Errorf("ожидается длина %d, но получена %d", size, len(result))
	}

	for i, num := range result {
		if num <= 0 {
			t.Errorf("элемент %d имеет значение %d, ожидаются только положительные числа", i, num)
		}
	}
}

func TestMaximum(t *testing.T) {
	result := maximum([]int{})
	if result != 0 {
		t.Errorf("пустой слайс ожидается 0, но получено %d", result)
	}

	result = maximum([]int{42})
	if result != 42 {
		t.Errorf("должно быть 42, но получено %d", result)
	}

	result = maximum([]int{3, 7, 2, 9, 1})
	if result != 9 {
		t.Errorf("должно быть 9, но получено %d", result)
	}

	result = maximum([]int{100, 50, 30, 10})
	if result != 100 {
		t.Errorf("должно быть 100, но получено %d", result)
	}

	result = maximum([]int{10, 30, 50, 100})
	if result != 100 {
		t.Errorf("должно быть 100, но получено %d", result)
	}

	result = maximum([]int{5, 5, 5, 5})
	if result != 5 {
		t.Errorf("должно быть 5, но получено %d", result)
	}
}

func TestMaxChunks(t *testing.T) {
	result := maxChunks([]int{})
	if result != 0 {
		t.Errorf("пустой слайс ожидается 0, но получено %d", result)
	}

	result = maxChunks([]int{42})
	if result != 42 {
		t.Errorf("должно быть 42, но получено %d", result)
	}

	result = maxChunks([]int{3, 7, 2, 9, 1, 5, 8, 6, 4, 10})
	if result != 10 {
		t.Errorf("должно быть 10, но получено %d", result)
	}

	slice := generateRandomElements(100)
	resultMaximum := maximum(slice)
	resultMaxChunks := maxChunks(slice)
	if resultMaximum != resultMaxChunks {
		t.Errorf("функции отличаются: maximum = %d, maxChunks = %d", resultMaximum, resultMaxChunks)
	}

	mediumSlice := generateRandomElements(100_000)
	resultMaximum = maximum(mediumSlice)
	resultMaxChunks = maxChunks(mediumSlice)
	if resultMaximum != resultMaxChunks {
		t.Errorf("функции отличаются для слайса: maximum = %d, maxChunks = %d", resultMaximum, resultMaxChunks)
	}
}
