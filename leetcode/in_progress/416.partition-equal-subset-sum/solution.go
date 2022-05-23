package main

func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
		// indexMap[v] = i
	}
	if sum%2 != 0 {
		return false
	}
	left := sum/2 - nums[0]
	return findLeft(1, left, nums)
}

func findLeft(beginIndex int, sum int, nums []int) bool {
	if sum == 0 {
		return true
	}
	if len(nums) < beginIndex+1 {
		return false
	}
	if len(nums) == beginIndex+1 {
		// if nums[beginIndex]==
		return nums[beginIndex] == sum
	}
	s1 := findLeft(beginIndex+1, sum, nums)
	s2 := findLeft(beginIndex+1, sum-nums[beginIndex], nums)
	return s1 || s2
}
