[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 105. Construct Binary Tree from Preorder and Inorder Traversal

Given two integer arrays `preorder` and `inorder` where `preorder` is the preorder traversal of a binary tree and `inorder` is the inorder traversal of the same tree, construct and return the binary tree.

### Example 1:

![image](./image/tree.jpeg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

### Example 2:

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

### Constraints:

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` and `inorder` consist of **unique** values.
- Each value of `inorder` also appears in `preorder`.
- `preorder` is **guaranteed** to be the preorder traversal of the tree.
- `inorder` is **guaranteed** to be the inorder traversal of the tree.

### Related Topics

- Array
- Hash Table
- Divide and Conquer
- Tree
- Binary Tree
  
---

# 解題方向

題目說明要我們用 `preorder traversal` 與 `inorder traversal` 的結果還原出一顆 binary tree

- `Pre-order (NLR)` 的訪問順序 :
  1. visit **the current node**
  2. recursively traverse **the current node's left subtree**
  3. recursively traverse **the current node's right subtree**

- `In-order (LNR)` 的訪問順序 :
  1. recursively traverse **the current node's left subtree**
  2. visit **the current node**
  3. recursively traverse **the current node's right subtree**

根據這兩種訪問方式的特性我們可以得出：

- `preorder` 的頭為 `root node`
- 對應回 `inorder` 可以知道 `root node` 的左邊為 `left subtree`，右邊為 `right subtree`
- 以此類推直到還原完畢

```
preorder = [  3,  9, 20, 15,  7]
              ^
            root node
                  v
inorder   = [ 9,  3, 15, 20,  7]
              ^       ^
   left subtree       right subtree
```

由於我們是從 `preorder` 中得到 `root node value` 後再從 `inorder` 找到對應的 `value`，這個過程每次都要花費 `O(n)` 的時間 (i.e. array search)

可以透過 `Hash Table` 來建立 `inorder` 的 key-value pairs，把查詢時間從 `O(n)` 降到 `O(1)`