[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-sum-circular-subarray/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 918. Maximum Sum Circular Subarray

自己看一下

### Related Topics

- Array
- Divide and Conquer
- Dynamic Programming
- Queue
- Monotonic Queue
  
---

# 解題方向

這題跟 [53. Maximum Subarray](./../p00053/README.md) 很像，只差在多了頭尾循環需要考慮，強烈建議先做做看那題

因此我們需要考慮的情況只多了一個頭尾循環的 case, 如下

```
case1. [x x x o o o o o o o x x x x] 或
case2. [o o o x x x x x x x o o o o] o = 我們想要的 subarray

case1 可以直接用 dp 求出
case2 可以以換個方式思考, 如果我們想要求出外側的最大值, 我們可以轉向找最小的 minarray, 用 sum(arr) - minsubarray 就是外側的總和

最後我們只要比較 maximum subarray 與 sum(arr) - minimum subarray 誰比較大就是答案了

必須注意一件事情是, 如果當 array 裡面的元素都是負值時, sum(arr) 與 minsubarray 找到的結果都會是一樣的, 此時相減後變成是0
因此我們必須直接返回 maximum subarray 的結果
```

剩下的直接看程式碼吧