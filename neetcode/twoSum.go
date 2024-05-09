// https://leetcode.com/problems/two-sum/description/

package main

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		a, ok := hash[target-nums[i]]
		if ok {
			return []int{i, a}
		}
		hash[nums[i]] = i
	}
	return []int{-1, -1}
}
