package main

// 前K个高频元素  使用桶排序：时间复杂度o(n),空间复杂度:o(n)

func topKFrequent(nums []int, k int) []int {
	m := make(map[int]int)

	for _, v := range nums {
		m[v] += 1
	}

	n := len(nums)
	buckets := make([][]int, n)

	for k1, v := range m {
		buckets[v-1] = append(buckets[v-1], k1)
	}

	result := make([]int, 0)

	for j := n - 1; j >= 0; j-- {
		if len(buckets[j]) != 0 {
			result = append(result, buckets[j]...)
			if len(result) == k {
				break
			}
		}
	}

	return result
}

func main() {

}
