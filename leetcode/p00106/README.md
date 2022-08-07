[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 106. Construct Binary Tree from Inorder and Postorder Traversal

Given two integer arrays `inorder` and `postorder` where `inorder` is the inorder traversal of a binary tree and `postorder` is the postorder traversal of the same tree, construct and return the binary tree.

### Example 1:

![image](./image/tree.jpeg)

```
Input: inorder = [9,3,15,20,7], postorder = [9,15,7,20,3]
Output: [3,9,20,null,null,15,7]
```

### Example 2:

```
Input: inorder = [-1], postorder = [-1]
Output: [-1]
```

### Constraints:

- `1 <= inorder.length <= 3000`
- `postorder.length == inorder.length`
- `-3000 <= inorder[i], postorder[i] <= 3000`
- `inorder` and `postorder` consist of **unique** values.
- Each value of `postorder` also appears in inorder.
- `inorder` is **guaranteed** to be the inorder traversal of the tree.
- `postorder` is **guaranteed** to be the postorder traversal of the tree.

### Related Topics

- Array
- Hash Table
- Divide and Conquer
- Tree
- Binary Tree
  
---

# 解題方向

題目說明要我們用 `inorder traversal` 與 `postorder traversal` 的結果還原出一顆 binary tree

- `In-order (LNR)` 的訪問順序 :
  1. recursively traverse **the current node's left subtree**
  2. visit **the current node**
  3. recursively traverse **the current node's right subtree**

- `Post-order (LRN)` 的訪問順序 :
  1. recursively traverse **the current node's left subtree**
  2. recursively traverse **the current node's right subtree**
  3. visit **the current node**

根據這兩種訪問方式的特性我們可以得出：

- `postorder` 的尾巴為 `root node`
- 對應回 `inorder` 可以知道 `root node` 的左邊為 `left substree`，右邊為 `right subtree`
- 以此類推直到還原完畢

```
   left subtree       right subtree
              v       v
inorder   = [ 9,  3, 15, 20,  7]
                  ^
                    root node
                              v
postorder = [ 9, 15,  7, 20,  3]
```

由於我們是從 `postorder` 中得到 `root node value` 後再從 `inorder` 找到對應的 `value`，這個過程每次都要花費 `O(n)` 的時間 (i.e. array search)

可以透過 `Hash Table` 來建立 `inorder` 的 key-value pairs，把查詢時間從 `O(n)` 降到 `O(1)`