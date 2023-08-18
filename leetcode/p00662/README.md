[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/maximum-width-of-binary-tree/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 662. Maximum Width of Binary Tree

自己看一下

### Related Topics

- Tree
- Depth-First Search
- Breath-First Search
- Binary Tree
  
---

# 解題方向

注意好 node 的編號即可  
 - left node index: root.node.val * 2
 - right node index: root.node.val * 2 + 1

範例如下:

```
      1
    /   \
   2     3
  / \   / \
 4  5  6   7

```