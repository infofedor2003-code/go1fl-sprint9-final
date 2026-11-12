package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
	}{
		{name: "пустой слайс", size: 0},
		{name: "отрицательный размер", size: -1},
		{name: "несколько элементов", size: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			elements := generateRandomElements(test.size)
			wantLength := test.size
			if wantLength < 0 {
				wantLength = 0
			}

			if len(elements) != wantLength {
				t.Fatalf("длина слайса = %d, ожидалось %d", len(elements), wantLength)
			}

			for _, element := range elements {
				if element <= 0 {
					t.Fatalf("сгенерировано неположительное число: %d", element)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "один элемент", data: []int{42}, want: 42},
		{name: "максимум в начале", data: []int{10, 5, 3}, want: 10},
		{name: "максимум в середине", data: []int{3, 10, 5}, want: 10},
		{name: "максимум в конце", data: []int{3, 5, 10}, want: 10},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maximum(test.data); got != test.want {
				t.Errorf("maximum(%v) = %d, ожидалось %d", test.data, got, test.want)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "пустой слайс", data: []int{}, want: 0},
		{name: "меньше восьми элементов", data: []int{4, 9, 2}, want: 9},
		{name: "остаток в последней части", data: []int{1, 2, 3, 4, 5, 6, 7, 8, 100}, want: 100},
		{name: "максимум в середине", data: []int{1, 2, 3, 4, 50, 6, 7, 8, 9, 10}, want: 50},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := maxChunks(test.data); got != test.want {
				t.Errorf("maxChunks(%v) = %d, ожидалось %d", test.data, got, test.want)
			}
		})
	}
}
