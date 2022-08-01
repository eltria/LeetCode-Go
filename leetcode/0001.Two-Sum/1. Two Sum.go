package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, x := range nums {
		y := target - x
		if j, ok := m[y]; ok {
			return []int{j, i}
		}
		m[x] = i
	}

	return nil
}
