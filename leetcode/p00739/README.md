[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/daily-temperatures/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 739. Daily Temperatures

Given an array of integers `temperatures` represents the daily temperatures, return an array `answer` such that `answer[i]` is the number of days you have to wait after the $i^{th}$ day to get a warmer temperature. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

### Example 1:

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

### Example 2:

```
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
```

### Example 3:

```
Input: temperatures = [30,60,90]
Output: [1,1,0]
```

### Constraints:

- 1 <= temperatures.length <= $10^5$
- `30 <= temperatures[i] <= 100`


### Related Topics

- Array
- Stack
- Monotonic Stack
  
---

# 解題方向

`Monotonic Stack` 觀念題，每次在新增一筆溫度資料進 `Stack` 之前都要先檢查過去天氣是否有符合條件可以移除的資料

- i.e. 找到當前位置 `temperatures[i]` 的下一個比它大的資料位置

Monotonic Stack 原則:

 - 永遠保持 `Stack` 內的元素不是單向遞增就是單向遞減
 - 換句話說，再加入新的元素前需要檢查 `Stack` 內的資料會不會破壞原則，如果會就必須事先移除