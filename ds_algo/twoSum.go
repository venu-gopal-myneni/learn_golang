package main

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for index, val := range nums {
		diff := target - val
		index2, exists := m[diff]
		if exists {
			return []int{index, index2}
		}
		m[val] = index
	}
	return nil

}
