package p00043

import "strings"

// Time Complexity: O(n^2)
// Space Complexity: O(n)
func multiply(num1 string, num2 string) string {
	// 乘數與被乘數只要為0, 就直接回傳結果
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}

	// 兩數相乘的最大值不會超過兩數的位數總和
	// e.g. 999 x 999 = 998,001
	res := make([]byte, len(num1)+len(num2))

	carried := byte(0)
	// 從尾巴反過來做
	for i := len(num2) - 1; i >= 0; i-- {
		j := len(num1) - 1

		// 用來控制 res 上的 cursor
		cur := len(num2) - 1 - i

		// 有可能數字相乘完畢, 但是進位還有前一位數的進位還沒處理完畢
		for j >= 0 || carried != 0 {
			num := byte(0)
			// 處理位數相乘
			if j >= 0 {
				num += ((num1[j] - '0') * (num2[i] - '0'))
			}
			// 處理進位
			if carried != 0 {
				num += carried
			}
			// 過去已經先運算好的資料也要拿回來處理
			num += res[cur]
			// 更新進位與該數位的資料
			carried, num = num/10, num%10
			res[cur] = num
			cur++
			j--
		}
	}

	idx := len(res) - 1
	// 確保最大位數不是0
	for res[idx] == 0 {
		idx--
	}

	var str strings.Builder
	// 反轉 []byte, 確保字串輸出結果為正確顯示方式 (左高右低)
	for idx >= 0 {
		str.WriteByte(res[idx] + '0')
		idx--
	}

	return str.String()
}
