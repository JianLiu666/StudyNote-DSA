package p00068

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	result := []string{}

	level := []string{}
	sum := 0
	for _, w := range words {
		// 這行預計寫下的字串總長度 + 新字串長度 + 空白至少需要預留的長度
		if sum+len(w)+len(level) <= maxWidth {
			// 只要還可以再繼續寫在同一行, 就一直加進去
			level = append(level, w)
			sum += len(w)

		} else {
			// 開始處理分段
			var buffer strings.Builder
			if len(level) == 1 {
				// 很有可能這行只能塞的下一個字串, 特別處理
				buffer.WriteString(level[0])
				buffer.WriteString(strings.Repeat(" ", maxWidth-len(level[0])))
			} else {
				// 判斷字串數量與每個字串之間需要的空白數量
				gap := (maxWidth - sum) / (len(level) - 1)
				for i, e := range level {
					buffer.WriteString(e)
					// 除了最後一個字串之外, 中間都需要插入空白
					// 如果不能剛好整除的話, 從前面的空白處開始每個都多塞一個進去
					if i != len(level)-1 {
						mod := (maxWidth - sum) % (len(level) - 1)
						if mod != 0 && i < mod {
							buffer.WriteString(strings.Repeat(" ", gap+1))
						} else {
							buffer.WriteString(strings.Repeat(" ", gap))
						}
					}
				}
			}
			// 處理完段落後, 也保留下一個段落開始的第一個字串
			result = append(result, buffer.String())
			level = level[:0]
			level = append(level, w)
			sum = len(w)
		}
	}
	// 處理最後一個段落
	result = append(result, strings.Join(level, " "))
	result[len(result)-1] += strings.Repeat(" ", maxWidth-len(result[len(result)-1]))

	return result
}

/**
分析問題
1. 每一行的字數加總起來必須剛好等於 maxWidth
2. 單字不夠時需要補上 ' ', 如果無法整除時, 從第一個開始多補
    - 空白部分就是 len(words)-1
     1234567890123456
    "This....is....an" -> This is an
*/
