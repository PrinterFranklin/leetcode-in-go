package leetcode0001

func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, v := range nums {
		if x, ok := hashTable[target-v]; ok {
			return []int{x, i}
		}
		hashTable[v] = i
	}
	return nil
}
