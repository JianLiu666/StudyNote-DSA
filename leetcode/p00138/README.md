[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/copy-list-with-random-pointer/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 138. Copy List with Random Pointer

A linked list of length `n` is given such that each node contains an additional random pointer, which could point to any node in the list, or `null`.

Construct a **deep copy** of the list. The deep copy should consist of exactly `n` **brand new** nodes, where each new node has its value set to the value of its corresponding original node. Both the `next` and `random` pointer of the new nodes should point to new nodes in the copied list such that the pointers in the original list and copied list represent the same list state. **None of the pointers in the new list should point to nodes in the original list**.

For example, if there are two nodes `X` and `Y` in the original list, where `X.random --> Y`, then for the corresponding two nodes `x` and `y` in the copied list, `x.random --> y`.

Return the head of the copied linked list.

The linked list is represented in the input/output as a list of `n` nodes. Each node is represented as a pair of `[val, random_index]` where:

- `val`: an integer representing `Node.val`
- `random_index`: the index of the node (range from `0` to `n-1`) that the `random` pointer points to, or `null` if it does not point to any node.

Your code will **only** be given the `head` of the original linked list.

### Example 1:

![image](./image/e1.png)

```
Input: head = [[7,null],[13,0],[11,4],[10,2],[1,0]]
Output: [[7,null],[13,0],[11,4],[10,2],[1,0]]
```

### Example 2:

![image](./image/e2.png)

```
Input: head = [[1,1],[2,1]]
Output: [[1,1],[2,1]]
```

### Example 3:

![image](./image/e3.png)

```
Input: head = [[3,null],[3,0],[3,null]]
Output: [[3,null],[3,0],[3,null]]
```

### Constraints:

- `0 <= n <= 1000`
- -$10^4$ <= Node.val <= $10^4$
- `Node.random` is `null` or is pointing to some node in the linked list.

### Related Topics

- Hash Table
- Linked List
  
---

# 解題方向

**Step.1 :** 先複製出一組與 `head` 相同的 `Linked List`，然後在過程中將兩邊的節點用 `Hash Table` 對應
**Step.2 :** 再從頭尋訪一次 `head` 處理 `random` 對應

fyi, 重看一遍之後發現 discuss 有一個跟鬼一樣的[解法](https://leetcode.com/problems/copy-list-with-random-pointer/discuss/43485/Clear-and-short-python-O(2n)-and-O(n)-solution)