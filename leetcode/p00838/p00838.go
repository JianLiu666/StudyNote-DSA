package p00838

import "strings"

// Time Complexity: O(n)
// Space Complexity: O(n)
func pushDominoes(dominoes string) string {
	var res strings.Builder

	// 補上不影響執行結果的邊界資料, 減低後續需要的邊界判斷
	dominoes = "L" + dominoes + "R"
	s := []byte(dominoes)

	l := 0
	for r := 0; r < len(s); r++ {
		// 遇到沒有被施加外力的骨牌先暫時跳過
		if s[r] == '.' {
			continue
		}

		if s[l] == 'L' && s[r] == 'L' {
			// 表示這中間所有的骨牌都會往左倒
			for l < r {
				res.WriteByte('L')
				l++
			}

		} else if s[l] == 'L' && s[r] == 'R' {
			// 表示這中間所有的骨牌都會維持原樣
			for l < r {
				res.WriteByte(s[l])
				l++
			}

		} else if s[l] == 'R' && s[r] == 'L' {
			// 表示這中間所有的骨牌都會往中心倒, 先處理往右倒的部分
			head, tail := l, r
			for head < tail {
				res.WriteByte('R')
				head++
				tail--
			}
			// 如果存在中心點位置, 該位置因為兩邊施力平均所以會維持站立
			if head == tail {
				res.WriteByte('.')
				head++
			}
			// 處理剩下往左倒的部分
			for head < r {
				res.WriteByte('L')
				head++
			}
			l = head

		} else {
			// 表示這中間所有的骨牌都會往右倒
			for l < r {
				res.WriteByte('R')
				l++
			}
		}
	}

	return res.String()[1:]
}
