package main

import (
	"fmt"
)

func main() {
	var data = []int{5, 1, 9, 7, 3, 6, 8}
	QuickSort(data)
	fmt.Println(data)

	var data2 = []int{10, 8, 5, 1, 9, 7, 3}
	QuickSort(data2)
	fmt.Println(data2)
}

func QuickSort(data []int) {
	if len(data) == 0 || len(data) == 1 {
		return
	}

	quickSort(data, 0, len(data)-1)
}

func quickSort(data []int, start, end int) {
	if start == end {
		return
	}

	baseIdx := swapParts(data, start, end)
	if baseIdx > start {
		quickSort(data, start, baseIdx-1)
	}

	if baseIdx < end {
		quickSort(data, baseIdx+1, end)
	}
}

// 已第一个数据为基准数分成两部分，左边都大于基准数，右边都小于基准数
func swapParts(data []int, start, end int) int {
	if end <= start {
		return start
	}

	var baseNum = data[start]
	var lastSmall = start
	var left = start + 1
	var right = end
	for {
		if left > right {
			break
		}

		// 从左到右找到第一个大于基准数的
		if data[left] < baseNum {
			lastSmall = left
			left++
			continue
		}

		// 从右向左找到第一个小于基准数的
		if data[right] > baseNum {
			right--
			continue
		}

		data[left], data[right] = data[right], data[left]
	}

	// 将基准数和最后一个小于自己的数交换
	if lastSmall != 0 {
		data[start], data[lastSmall] = data[lastSmall], data[start]
	}

	return lastSmall
}
