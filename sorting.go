package main

import "fmt"
import "math/rand"

func bubblesort(nums []int) []int {
	// base case
	if len(nums) <= 1 {
		return nums
	}

	for i := len(nums) - 1; i > 0; i-- {
		// bubble over all
		for j := 0; j < i; j++ {
			if nums[j] > nums[j + 1] {
				var first = nums[j]
				nums[j] = nums[j + 1]
				nums[j + 1] = first
			}
		}
	}

	return nums
}

func mergesort(nums []int) []int {
	// base case
	if len(nums) <= 1 {
		return nums
	}

	// divide the slice
	var middle = len(nums) / 2
	var left = nums[:middle]
	var right = nums[middle:]

	// sort divided slices
	left = mergesort(left)
	right = mergesort(right)

	// merge the now sorted slices
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	// the slice to be returned
	var merged = make([]int, 0)

	// merge slices
	for len(left) > 0 || len(right) > 0 {
		// case both slices not empty
		if len(left) > 0 && len(right) > 0 {
			if left[0] <= right[0] {
				merged = append(merged, left[0])
				left = left[1:]
			} else {
				merged = append(merged, right[0])
				right = right[1:]
			}
		} else if len(left) > 0 {
			merged = append(merged, left[0])
			left = left[1:]
		} else if len(right) > 0 {
			merged = append(merged, right[0])
			right = right[1:]
		}
	}

	return merged
}

func randomInts(length int) []int {
	var nums = make([]int, length)
	for i := 0; i < length; i++ {
		nums[i] = rand.Intn(length)
	}
	return nums
}

func main() {
	var list = randomInts(20)
	fmt.Println(list)
	fmt.Println(bubblesort(list))
}