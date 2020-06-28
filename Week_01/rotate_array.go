package main

// 采用环状替换
func rotate(nums []int, k int) {
	k = k % len(nums)

	count := 0
	for start := 0; count < len(nums); start++ {
		// count是用来衡量最终结束的条件，因为当len(nums)%k=0时会一直循环在一次的替换过程中
		current := start
		prev := nums[start]
		for {
			next := (current + k) % len(nums)
			temp := nums[next]
			nums[next] = prev
			prev = temp
			current = next
			count++

			if current == start {
				// 回到原来位置
				break
			}
		}
	}
}

// 采用反转，这个方法更好
func rotate2(nums []int, k int) {
	k %= len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// 经典的反转写法
func reverse(nums []int, start, end int) {
	for start < end {
		nums[end], nums[start] = nums[start], nums[end]
		start++
		end--
	}
}

func main() {

}
