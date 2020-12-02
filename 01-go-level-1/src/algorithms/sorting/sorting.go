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
