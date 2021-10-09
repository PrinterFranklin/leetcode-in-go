package leetcode0912

import (
	"math/rand"
	"sort"
)

// Ints is int array
type Ints []int

func (n Ints) Len() int {
	return len(n)
}

func (n Ints) Less(i, j int) bool {
	return n[i] < n[j]
}

func (n Ints) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

// Solution 0: Standard Sort (implements Solution 1-4)
// Time: O(NlogN)
// Space: O(1)
// Stable: unknown
func sortArray(nums Ints) []int {
	sort.Sort(nums)
	return nums
}

// Solution 1: Quick Sort (fast)
// Time: O(NlogN)
// Space: O(1)
// Stable: no
func sortArray1(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, lo, hi int) {
	if lo < hi {
		p := randomPartition(nums, lo, hi)
		quickSort(nums, lo, p-1)
		quickSort(nums, p+1, hi)
	}
}

func randomPartition(nums []int, lo, hi int) int {
	// pick a random number from lo to hi
	r := rand.Intn(hi-lo+1) + lo
	nums[r], nums[hi] = nums[hi], nums[r]
	return partition(nums, lo, hi)
}

func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	// i is the edge of two arrays
	i := lo
	for j := lo; j < hi; j++ {
		if nums[j] < pivot {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}
	nums[i], nums[hi] = nums[hi], nums[i]
	return i
}

// Solution 2: Heap Sort (fast)
// Time: O(NlogN)
// Space: O(1)
// Stable: no
func sortArray2(nums []int) []int {
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {

}

func sink(nums []int, i, len int) {

}

// Solution 3: Merge Sort (fast)
// Time: O(NlogN)
// Space: O(N)
// Stable: yes
func sortArray3(nums []int) []int {
	mergeSort(nums, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, l, r int) {
	if l < r {
		m := (l + r) / 2
		mergeSort(nums, l, m)
		mergeSort(nums, m+1, r)
		merge(nums, l, m, r)
	}
}

func merge(nums []int, l, m, r int) {
	// Step 1: create two temp arrays
	n1, n2 := m-l+1, r-m

	left := make([]int, n1)
	right := make([]int, n2)

	// Step 2: write elements to temp arrays
	for i := 0; i < n1; i++ {
		left[i] = nums[l+i]
	}
	for j := 0; j < n2; j++ {
		right[j] = nums[m+1+j]
	}

	// initial indexes
	i, j, k := 0, 0, l
	// Step 3: write back the array elements in order
	for i < n1 && j < n2 {
		if left[i] < right[j] {
			nums[k] = left[i]
			i++
		} else {
			nums[k] = right[j]
			j++
		}
		k++
	}
	// Step 4: write back the remaining elements
	for i < n1 {
		nums[k] = left[i]
		i++
		k++
	}
	for j < n2 {
		nums[k] = right[j]
		j++
		k++
	}
}

// Solution 4: Insertion Sort (slow)
// Time: O(N^2)
// Space: O(1)
// Stable: yes
func sortArray4(nums []int) []int {
	// Loop 1: iterate unsorted array
	for i := 1; i < len(nums); i++ {
		tmp := nums[i]
		j := i - 1
		// Loop 2: compare and move
		for j >= 0 && nums[j] > tmp {
			nums[j+1] = nums[j]
			j--
		}
		nums[j+1] = tmp
	}
	return nums
}

// Solution 5: Bubble Sort (too slow)
// Time: O(N^2)
// Space: O(1)
// Stable: yes
func sortArray5(nums []int) []int {
	// Loop 1: compare n-1 times to find the largest element
	for i := 0; i < len(nums)-1; i++ {
		// Loop 2: reduce array capacity
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

// Solution 6: Selection Sort (too slow)
// Time: O(N^2)
// Space: O(1)
// Stable: no
func sortArray6(nums []int) []int {
	// Loop 1: swap the minimum element to the tail of a sorted array
	for i := 0; i < len(nums); i++ {
		minIdx := i
		// Loop 2: find the minimum element index
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}
	return nums
}

// Solution 7: Shell Sort
// Time: O(N^X)
// Space: O(1)
// Stable: no
func sortArray7(nums []int) []int {
	return nil
}

// Solution 8: Counting Sort
// Time: O(N+k)
// Space: O(N+k)
// Stable: yes
func sortArray8(nums []int) []int {
	return nil
}

// Solution 9: Radix Sort
// Time: O(N*k)
// Space: O(N+k)
// Stable: yes
func sortArray9(nums []int) []int {
	return nil
}

// Solution 10: Bucket Sort
// Time: O(N+k)
// Space: O(N+k)
// Stable: yes
func sortArray10(nums []int) []int {
	return nil
}
