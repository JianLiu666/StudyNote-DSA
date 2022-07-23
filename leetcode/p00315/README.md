[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/count-of-smaller-numbers-after-self/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 315. Count of Smaller Numbers After Self

You are given an integer array `nums` and you have to return a new `counts` array. The `counts` array has the property where `counts[i]` is the number of smaller elements to the right of `nums[i]`.

### Example 1:

```
Input: nums = [5,2,6,1]
Output: [2,1,1,0]

Explanation:
 - To the right of 5 there are 2 smaller elements (2 and 1).
 - To the right of 2 there is only 1 smaller element (1).
 - To the right of 6 there is 1 smaller element (1).
 - To the right of 1 there is 0 smaller element.
```

### Example 2:

```
Input: nums = [-1]
Output: [0]
```

### Example 3:

```
Input: nums = [-1,-1]
Output: [0,0]
```

### Constraints:

- 1 <= nums.length <= $10^5$
- -$10^4$ <= nums[i] <= $10^4$

### Related Topics

- Array
- Binary Search
- Divide and Conquer
- Binary Indexed Tree
- Segment Tree
- Merge Sort
- Ordered Set
  
---

# 解題方向

### Solved using Merge Sort Algorithm

本題的目標是返回一個新的 `array`，其中的每個 `array[i]` 對應的是從 `nums[i]` 開始往後數到底，所出現所有小於 `nums[i]` 的元素數量。

我們可以藉由 Merge Sort 的排序過程順便記數即可完成，想法如下:

```
討論之前我們先定義後續的表達方式，依題目給定的測資 Example 1 為例：
  - input nums =    [5,2,6,1]
  - output result = [2,1,1,0]

  我們將之表示為: 5(2), 2(1), 6(1), 1(0)，即5的右邊有2個數字比5還小 (2, 1)

Merge Sort : part of top-down phase

  按照原本的 Merge Sort 過程依序拆分

          +---+---+---+---+
          | 5 | 2 | 6 | 1 |
          +---+---+---+---+
          /               \
      +---+---+       +---+---+
      | 5 | 2 |       | 6 | 1 |
      +---+---+       +---+---+
      /       \       /       \
    +---+   +---+   +---+   +---+
    | 5 |   | 2 |   | 6 |   | 1 |
    +---+   +---+   +---+   +---+

Merge Sort : part of bottom-up phase

  計數過程在重新排序時順便考慮，這時有三種情境:

    1. left > right    2. left < right    3. left = right

      +---+---+          +---+---+          +---+---+  
      | ? | ? |          | ? | ? |          | ? | ? |  
      +---+---+          +---+---+          +---+---+  
      /       \    or    /       \    or    /       \  
    +---+   +---+      +---+   +---+      +---+   +---+
    | 5 |   | 2 |      | 2 |   | 5 |      | 5 |   | 5 |
    +---+   +---+      +---+   +---+      +---+   +---+

    依照本題目標，我們只需要再發生第一種狀況時額外處理記數，其他皆按照 Merge Sort 的流程，該怎麼辦就怎麼辦
  
  同時因為 Merge Sort 的原理我們可以把焦點放在每個階段的 left sub-array 與 right sub-array 之間的比較即可，如下所示:

        +---+---+---+---+---+---+
        | ? | ? | ? | ? | ? | ? |
        +---+---+---+---+---+---+
          /                  \
    +---+---+---+       +---+---+---+
    | 7 | 8 | 9 |       | 1 | 2 | 3 |
    +---+---+---+       +---+---+---+
    
    這裡要注意一件事情，如果在每一次的 left > right 發生時，就對 left sub-array 內的所有元素記數一遍的話，就會讓時間複雜度變成 O(n^2)
    
    為了避免這種情況發生，我們可以定義一個參數先幫我們記數 (e.g. accumulator)，直到下面兩種情況發生時，再將暫存結果累加回去
      
      1. 過程中有部分的 left[i] 比 right[i] 還小時，先將目前為止的暫存結果累加給 left[i]
         例如:

                +---+---+---+---+---+---+       accumulator = 1 (因為 4 < 5)
                | 4 | ? | ? | ? | ? | ? |       
                +---+---+---+---+---+---+       此時只先更新 5 的累計結果，即: 5(1), 8(0), 9(0)
                  /                  \
            +---+---+---+       +---+---+---+
            | 5 | 8 | 9 |       | 4 | 6 | 7 |
            +---+---+---+       +---+---+---+
              ^                       ^
          left cursor            right cursor


      2. right sub-array 全空，一口氣對剩下還在 left sub-array 的元素將暫存結果累加回去
         例如:

                +---+---+---+---+---+---+       accumulator = 3 (因為 right sub-array 全都比 left sub-array 還小)
                | 4 | 5 | ? | ? | ? | ? |       
                +---+---+---+---+---+---+       此時的累計結果與上圖相同:
                  /                  \            - 5(1), 8(0), 9(0)
            +---+---+---+       +---+---+---+   
            | 5 | 8 | 9 |       | 4 | 6 | 7 |   現在將累計結果一次寫回 8, 9，即:
            +---+---+---+       +---+---+---+     - 5(1), 8(3), 9(3)
                  ^                           ^
              left cursor                right cursor

  持續重複直到 Merge Sort 結束，即可得到題目所需的累計結果。
```
