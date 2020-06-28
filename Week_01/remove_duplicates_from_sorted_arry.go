package main

// 双指针法
func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	i, j := 0, 1
	for ; j <= len(nums)-1; j++ {
		if nums[j] != nums[j-1] {
			if j-i > 1 {
				// 如果非相邻则复制
				nums[i+1] = nums[j]
			}
			i++
		}
	}

	return i + 1
}

func main() {

}
