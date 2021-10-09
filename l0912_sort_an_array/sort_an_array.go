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

// Solution 1: Quick Sort
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

// Solution 2: Heap Sort
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

// Solution 3: Merge Sort
// Time: O(NlogN)
// Space: O(N)
// Stable: yes
func sortArray3(nums []int) []int {
	return nil
}

// Solution 4: Insertion Sort
// Time: O(N^2)
// Space: O(1)
// Stable: yes
func sortArray4(nums []int) []int {
	return nil
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

// Solution 6: Selection Sort
// Time: O(N^2)
// Space: O(1)
// Stable: no
func sortArray6(nums []int) []int {
	return nil
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
