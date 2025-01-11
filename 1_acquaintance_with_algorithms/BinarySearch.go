package main

import (
	"fmt"
)

// BinarySearch - це структура, яка містить методи для виконання бінарного пошуку
type BinarySearch struct{}

// searchIterative - ітеративний метод бінарного пошуку
func (bs BinarySearch) searchIterative(list []int, item int) int {
	low := 0
	high := len(list) - 1

	for low <= high {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1 // Повертає -1, якщо елемент не знайдено
}

// searchRecursive - рекурсивний метод бінарного пошуку
func (bs BinarySearch) searchRecursive(list []int, low, high, item int) int {
	if high >= low {
		mid := (low + high) / 2
		guess := list[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			return bs.searchRecursive(list, low, mid-1, item)
		}
		return bs.searchRecursive(list, mid+1, high, item)
	}

	return -1 // Повертає -1, якщо елемент не знайдено
}

func main() {
	bs := BinarySearch{}
	my_list := []int{1, 3, 5, 7, 9}

	fmt.Println(bs.searchIterative(my_list, 3))  // => 1
	fmt.Println(bs.searchIterative(my_list, -1)) // => -1

	fmt.Println(bs.searchRecursive(my_list, 0, len(my_list)-1, 3))  // => 1
	fmt.Println(bs.searchRecursive(my_list, 0, len(my_list)-1, -1)) // => -1
}
