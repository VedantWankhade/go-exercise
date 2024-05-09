// https://leetcode.com/problems/contains-duplicate/description/

package main

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	dup := make(map[int]struct{})
	for _, n := range nums {
		if _, ok := dup[n]; ok {
			return true
		}
		dup[n] = struct{}{}
	}
	return false
}
