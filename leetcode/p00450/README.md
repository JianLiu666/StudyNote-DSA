[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/delete-node-in-a-bst/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 450. Delete Node in a BST

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the **root node reference** (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

1. Search for a node to remove.
2. If the node is found, delete the node.

### Example 1:

![image](./image/del_node_1.jpeg)

```
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]

Explanation:
 - Given key to delete is 3. So we find the node with value 3 and delete it.
 - One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
 - Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
```

![image](./image/del_node_supp.jpeg)

### Example 2:

```
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]

Explanation:
 - The tree does not contain a node with value = 0.
```

### Example 3:

```
Input: root = [], key = 0
Output: []
```

### Constraints:

- The number of nodes in the tree is in the range [0, $10^4$].
- -$10^5$ <= Node.val <= $10^5$
- Each node has a **unique** value.
- `root` is a valid binary search tree.
- -$10^5$ <= key <= $10^5$

#### Follow up:

Could you solve it with time complexity `O(height of tree)`?

### Related Topics

- Tree
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

delete node 時總共有下面四種情境：

1. `node.left == NULL and node.right == NULL` -> 直接刪除即可
2. `node.left == NULL and node.right != NULL` -> 用 `node.right` 取代
3. `node.left != NULL and node.right == NULL` -> 用 `node.left` 取代
4. `node.left != NULL and node.right != NULL` -> 兩種方案
   a. 用 `left subtree` 的 `max node` 取代
   b. 用 `right subtree` 的 `min node` 取代