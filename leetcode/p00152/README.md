[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-product-subarray/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 152. Maximum Product Subarray

自己看一下

### Related Topics

- Array
- Dynamic Programming
  
---

# 解題方向

參考 [官神影片](https://www.youtube.com/watch?v=LQuZkqx16PU)  
這邊只記錄一下片段：
- 因為這題的元素範圍在 [-10, 10] 之間, 所以可能會出現負數相乘後變成正數的問題，因此在保留最大值的過程中，同時也要留意到目前為止的最小值