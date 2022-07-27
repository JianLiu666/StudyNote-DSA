[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/min-stack/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 155. Min Stack

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

 - `MinStack()` initializes the stack object.
 - `void push(int val)` pushes the element `val` onto the stack.
 - `void pop()` removes the element on the top of the stack.
 - `int top()` gets the top element of the stack.
 - `int getMin()` retrieves the minimum element in the stack.

You must implement a solution with `O(1)` time complexity for each function.

### Example 1:

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
 - MinStack minStack = new MinStack();
 - minStack.push(-2);
 - minStack.push(0);
 - minStack.push(-3);
 - minStack.getMin(); // return -3
 - minStack.pop();
 - minStack.top();    // return 0
 - minStack.getMin(); // return -2
```

### Constraints:

- -$2^{31}$ <= val <= $2^{31}$ - 1
- Methods `pop`, `top` and `getMin` operations will always be called on **non-empty** stacks.
- At most 3 * $10^4$ calls will be made to `push`, `pop`, `top`, and `getMin`.

### Related Topics

- Stack
- Design
  
---

# 解題方向

`Stack` 暖身題，`getMin()` 只要在多使用一個 `stack` 維護即可

- 因為 Last-in-first-out (LIFO)，只要進來的資料比 `stack` 最頂端的資料還大時就可以忽略不管，畢竟他會更早被 pop 掉
