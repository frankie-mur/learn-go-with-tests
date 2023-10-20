package slices

func Sum(l []int) int {
	var total int
	for _, nums := range l {
		total += nums
	}
	return total
}

func SumAllTails(numsToSum ...[]int) []int {
	var result []int
	for _, nums := range numsToSum {
		if len(nums) == 0 {
			result = append(result, 0)
			continue
		}
		result = append(result, Sum(nums[1:]))
	}
	return result
}
