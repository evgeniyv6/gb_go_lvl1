package sorting

import (
	"math/rand"
)

func swap(ar []int, i int, j int) {
	ar[i], ar[j] = ar[j], ar[i]
}

// QuickSort algorithm for int
func QuickSort(nums []int, begin int, end int) {
	if begin > end {
		return
	}
	i, j, pivot := begin, end, nums[rand.Int()%len(nums)]
	for i <= j {
		for nums[i] < pivot {
			i++
		}
		for nums[j] > pivot {
			j--
		}
		if i <= j {
			swap(nums, i, j)
			i++
			j--
		}
	}
	QuickSort(nums, begin, j)
	QuickSort(nums, i, end)
}

// BubleSort algo
func BubleSort(s []int) []int {
	sTmp := append([]int(nil), s...)
	n := len(sTmp)
	count := 1
	for count < n {
		for i := 0; i < n-count; i++ {
			if sTmp[i] > sTmp[i+1] {
				sTmp[i], sTmp[i+1] = sTmp[i+1], sTmp[i]
			}
		}
		count++
	}
	return sTmp
}

// InsertationSort algo
func InsertationSort(s []int) []int {
	sTmp := append([]int(nil), s...)
	n := len(sTmp)
	for i := 1; i < n; i++ {
		el := sTmp[i]
		j := i - 1
		for j >= 0 && el < sTmp[j] {
			sTmp[j+1] = sTmp[j]
			j--
		}
		sTmp[j+1] = el
	}
	return sTmp
}
