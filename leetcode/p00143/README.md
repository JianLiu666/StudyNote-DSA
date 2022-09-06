[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/reorder-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 143. Reorder List

You are given the head of a singly linked-list. The list can be represented as:

> L~0~ → L~1~ → … → L~n-1~ → L~n~

Reorder the list to be on the following form:

> L~0~ → L~n~ → L~1~ → L~n-1~ → L~2~ → L~n-2~ → …

You may not modify the values in the list's nodes. Only nodes themselves may be changed.

### Example 1:

![image](./image/reorder1linked-list.jpeg)

```
Input: head = [1,2,3,4]
Output: [1,4,2,3]
```

### Example 2:

![image](./image/reorder2-linked-list.jpeg)

```
Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
```

### Constraints:

- The number of nodes in the list is in the range [1, 5 * $10^4$].
- `1 <= Node.val <= 1000`

### Related Topics

- Linked List
- Two Pointers
- Stack
- Recursion
  
---

# 解題方向

### Solved using Stack concept

一開始還在思考 Singly Linked List 要怎麼樣才能把尾巴的資料擺回來前面，直到看到 `Stack` 的標籤後才瞬間恍然大悟，做的題目還是太少了...

有了 `Stack` 後事情就簡單很多，第一次用 Slow Fast Pointers 找出 Linked List 的中間後，把從 Slow Pointer 開始到最後一個節點通通都丟進 `Stack` 內

接著在用一個新的指針從 `head` 開始把資料從 `Stack` 一筆一筆拿出來交叉串起來即可