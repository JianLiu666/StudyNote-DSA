[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/unique-binary-search-trees-ii/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 95. Unique Binary Search Trees II

Given an integer `n`, return all the structurally unique **BST**'s (binary search trees), which has exactly `n` nodes of unique values from `1` to `n`. Return the answer in **any order**.

Example 1:

![image](./image/uniquebstn3.jpeg)

```
Input: n = 3
Output: [[1,null,2,null,3],[1,null,3,2],[2,1,3],[3,1,null,null,2],[3,2,null,1]]
```

### Example 2:

```
Input: n = 1
Output: [[1]]
```

### Constraints:

- `1 <= n <= 8`

### Related Topics

- Dynamic Programming
- Backtracking
- Tree
- Binary Search Tree
- Binary Tree
  
---

# 解題方向

### Solved using Recursion concept

最直覺的思考方式就是遍歷一次 `[1, n]`，就像 in-order traversal 一樣，切成左右子樹 `[1, current-1` 與 `[current+1, n]`，並繼續往下展開，直到走到 leaf node 才停止，示意圖如下：

```
     [1,              2,               3]
 ┌----┻----┐        ┌-┻-┐         ┌----┻----┐
 x   [2,       3]  [1] [3]  [1,       2]    x
    ┌-┻-┐    ┌-┻-┐ ┌┻┐ ┌┻┐ ┌-┻-┐    ┌-┻-┐
    x  [3]  [2]  x x x x x x  [2]  [1]  x 
       ┌┻┐  ┌┻┐               ┌┻┐  ┌┻┐
       x x  x x               x x  x x
```

如上圖所示，能夠成功將所有數字都用上，且走到 leaf node 的路有：

```
  1         1        2        3         3
┌-┻-┐     ┌-┻-┐    ┌-┻-┐    ┌-┻-┐     ┌-┻-┐
x   2     x   3    1   3    1   x     2   x
  ┌-┻-┐     ┌-┻-┐         ┌-┻-┐     ┌-┻-┐
  x   3     2   x         x   2     1   x
     ┌┻┐   ┌┻┐               ┌┻┐   ┌┻┐
     x x   x x               x x   x x
```
