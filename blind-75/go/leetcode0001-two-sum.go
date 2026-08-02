// Problem: Two Sum
// https://leetcode.com/problems/two-sum/

package main

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]

		if value, ok := numsMap[complement]; ok {
			return []int{i, value}
		} else {
			numsMap[nums[i]] = i
		}
	}

	return []int{-1, -1}
}
