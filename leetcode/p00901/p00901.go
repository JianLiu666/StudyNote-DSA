package p00901

type StockSpanner struct {
	data [][2]int // 0: span, 1: price
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	count := 1

	// 每次要彈出一個數字時, 都要把他的 span 也記下來, 表示除了這個數字之外, 還有多少個比他更小的數字
	// 因為時間是順序遞增的, 如何證明這樣做合法?
	// span:     1   1   1
	//       [ 100, 80, 60 ] <- 這時 70 要進來了, 就把 60 彈掉且記下來, 變成
	// span:     1   1         2
	//       [ 100, 80, '60', 70 ]
	//                    ^ 這個數字已經不存在了
	// 所以即使下一次又再遇到 60, 有 70 擋在前面所以新的 60 span 還是 1
	// 但如果有 80 出現的話如下:
	// span:     1   1         2   1
	//       [ 100, 80, '60', 70, 60 ] <- 80
	//
	// 這個 80 就會依序收掉 60(span:1), 70(span:2) <- 因為 70 都能收掉了, 曾經比 70 小的那些數字肯定也能收掉
	for len(this.data) > 0 && this.data[len(this.data)-1][1] <= price {
		count += this.data[len(this.data)-1][0]
		this.data = this.data[:len(this.data)-1]
	}

	// 每次把 price 加上去時都要順便紀錄下, 對於 price 而言, 到底曾經有多少數字比他更小
	this.data = append(this.data, [2]int{count, price})

	return count
}

/**
分析題目:
每次加進去的時候都要排序, 保持單調遞增, 只要下一個要加入時他不是最低的, 就看要 pop 幾次就知道離多遠了

["StockSpanner", "next", "next", "next", "next", "next", "next", "next"]
[[], [100], [80], [60], [70], [60], [75], [85]]

["StockSpanner","next","next","next","next","next"]
[[],[29],[91],[62],[76],[51]]

 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
*/
