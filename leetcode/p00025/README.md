[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reverse-nodes-in-k-group/)
![image](https://img.shields.io/badge/Difficulty-Hard-red)

---

# 25. Reverse Nodes in k-Group

Given the `head` of a linked list, reverse the nodes of the list `k` at a time, and return the modified list.

`k` is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of `k` then left-out nodes, in the end, should remain as it is.

You may not alter the values in the list's nodes, only nodes themselves may be changed.

### Example 1:

![image](./image/reverse_ex1.jpeg)

```
Input: head = [1,2,3,4,5], k = 2
Output: [2,1,4,3,5]
```

### Example 2:

![image](./image/reverse_ex2.jpeg)

```
Input: head = [1,2,3,4,5], k = 3
Output: [3,2,1,4,5]
```

### Constraints:

- The number of nodes in the list is `n`.
- `1 <= k <= n <= 5000`
- `0 <= Node.val <= 1000`

#### Follow-up: 

Can you solve the problem in `O(1)` extra memory space?

### Related Topics

- Linked List
- Recursion
  
---

# 解題方向

這題的核心概念還是翻轉 Linked List，只要能夠將 group 的頭跟尾定位出來後，剩下的就只是每個 group 內的翻轉問題，解法就跟 `Easy` 的[這題](./../p00026/README.md)一模一樣

LeetCodeCN 的[這篇圖文](https://leetcode.cn/problems/reverse-nodes-in-k-group/solution/tu-jie-kge-yi-zu-fan-zhuan-lian-biao-by-user7208t/)解釋得很好