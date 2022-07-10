[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/linked-list-cycle/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 141. Linked List Cycle

Given `head`, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the `next` pointer. Internally, `pos` is used to denote the index of the node that tail's `next` pointer is connected to. **Note that** `pos` **is not passed as a parameter**.

Return `true` if there is a cycle in the linked list. Otherwise, return `false`.

### Example 1:

![image](./image/circularlinkedlist_test1.png)

```
Input: head = [3,2,0,-4], pos = 1
Output: true

Explanation:
 - There is a cycle in the linked list, where the tail connects to the 1st node (0-indexed).
```

### Example 2:

![image](./image/circularlinkedlist_test2.png)

```
Input: head = [1,2], pos = 0
Output: true

Explanation:
 - There is a cycle in the linked list, where the tail connects to the 0th node.
```

### Example 3:

![image](./image/circularlinkedlist_test3.png)

```
Input: head = [1], pos = -1
Output: false

Explanation:
 - There is no cycle in the linked list.
```

### Constraints:

- The number of the nodes in the list is in the range [0, $10^4$].
- -$10^5$ <= Node.val <= $10^5$
- `pos` is `-1` or a **valid index** in the linked-list.

#### Follow up:

Can you solve it using `O(1)` (i.e. constant) memory?

### Related Topics

- Hash Table
- Linked List
- Two Pointers
  
---

# 解題方向

Two Pointers (slow and fast pointer)

 - fast pointer 的移動速度是 slow pointer 的兩倍

假如給定的 `Linked List` 有閉環的話，那 slow pointer 與 fast pointer 最終一定會相遇