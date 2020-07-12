### 第三周总结

---

&nbsp;&nbsp;&nbsp;&nbsp;本周学习了递归和分治、回溯算法。
* 递归算法在链表中使用较多，因为链表本身就是递归定义的，所以本周对链表的3种遍历以及之前的一些老题目的算法整理了下，在如下的图中：<br>
![链表的递归](https://github.com/danyXu/algorithm011-class01/blob/master/Week_03/1.png)

* 分治思想是大问题分解成彼此独立的小问题，小问题之间没有任何关系，具体讲就是不存在状态依赖，每个小问题的解决方法是一样的，也就是递归的处理就行，比如归并排序、mapReduce处理方法，而回溯算法的基本思想是穷举，遍历所有情况，但需要在关键的地方进行回头，要么是已经处理过了或不符合条件不用处理，如：
  遍历完所有节点，全局维护一个遍历状态；走到最后一层，处理完了就完了
*   ```go
    // 归并排序
    func mergerSort(arr []int, a, b int) {
        if b-a <= 1 {
            return
        }
    
        c := (a + b) / 2
        mergerSort(arr, a, c)
        mergerSort(arr, c, b)
    
        arrLeft := make([]int, c-a)
        arrRight := make([]int, b-c)
        copy(arrLeft, arr[a:c])
        copy(arrRight, arr[c:b])
        i := 0
        j := 0
        for k := a; k < b; k++ {
            if i >= c-a {
                arr[k] = arrRight[j]
                j++
            } else if j >= b-c {
                arr[k] = arrLeft[i]
                i++
            } else if arrLeft[i] < arrRight[j] {
                arr[k] = arrLeft[i]
                i++
            } else {
                arr[k] = arrRight[j]
                j++
            }
        }
    }

    // 回溯的0-1背包问题
    func backpack(nums [][]int, total int) int {
      dp := make([][]int, len(nums))
      //初始化二维数组
      for i := 0; i < len(nums); i++ {
        dp[i] = make([]int, total+1)
      }
      
      //放入第一个物品，填第一行列表
      for i:= nums[0][0]; i < total; i++ {
        dp[0][i] = nums[0][1]
      }
      
      for i := 1; i < len(nums); i++ {
        for j:= nums[i][0]; j < total; j++ {
          dp[i][j] = max(dp[i-1][j], dp[i-1][j-nums[i][0]] + nums[i][1])
        }
      }
      return dp[len(nums) - 1][total]
    }
    
    ```