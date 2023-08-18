[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/path-sum-iii/description/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 437. Path Sum III

自己看一下

### Related Topics

- Tree
- Depth-First Search
- Binary Tree
  
---

# 解題方向

### Brute-force

核心思想就是 tree traversal，用任何一種遍歷方式都可以，重點在於除了基本的 path sum 之外，每個節點也要當作是 root 重新出發一次  
詳細直接看程式碼吧

**時間複雜度**
- 基本遍歷一次: O(n)
- subtree 開始的遍歷: O(n) or O(logn), 取決於是 single side tree 還是 balanced tree
- 總複雜度: O(nlogn) or O(n^2)

**空間複雜度**
- 暫存開銷: O(1)
- recursion 開銷: O(n) or O(logn), 理由同上

### Memorization

可以參考這篇 [leetcode solution](https://leetcode.com/problems/path-sum-iii/solutions/141424/python-step-by-step-walk-through-easy-to-understand-two-solutions-comparison/)