package p00416

func canPartition(nums []int) bool {
	maximum := 0
	for _, n := range nums {
		maximum += n
	}

	// 不是偶數沒辦法平均兩半, 就不需要找了
	if maximum%2 != 0 {
		return false
	}

	// 計算組出所有數字的辦法, 預設 0 一定組的出來, 都不拿就好
	dp := make([]bool, maximum+1)
	dp[0] = true

	// step.1 先找背包
	for _, n := range nums {
		// step.2 在找容量
		// 從尾往前找的原因是因為, 如果從頭開始找的話, 在還沒有走完完整的一遍 [0, maximum/2] 之前
		// dp 就已經被更新了, 會造成衝突
		for i := maximum / 2; i >= 0; i-- {
			// 意思是如果我加上 n 之前的這個數字就已經組的出來的話, 那我現在這個位置只要再加上 n 就一樣組的出來
			if i-n >= 0 && dp[i-n] {
				dp[i] = true
			}
		}
	}

	return dp[maximum/2]
}

/**
每一個元素都可以選擇拿或不拿 -> 拿的一群, 不拿的一群 -> 本質上還是背包問題
思考狀態轉移方程:

nums[i]: 拿 不拿
         | x |
nums[i]: 拿 不拿
*/
