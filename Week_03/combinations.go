package main

func combine(n int, k int) [][]int {
	if n < 1 || k < 1 || n < k {
		return nil
	}
	cur := make([]int, k, k)
	result := [][]int{}
	helper(cur, 0, 1, n, k, &result)

	return result
}

// ci 是在本次函数中，cur中要添加的元素的位置，避免append操作
func helper(cur []int, ci, start, n, k int, result *[][]int) {
	if k == 0 {
		tmp := make([]int, len(cur))
		copy(tmp, cur)
		*result = append(*result, tmp)
		return
	}

	for i := start; i <= n-k+1; i++ {
		cur[ci] = i
		helper(cur, ci+1, i+1, n, k-1, result)

	}
}

func main() {

}
