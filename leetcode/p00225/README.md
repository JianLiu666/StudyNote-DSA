[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/implement-stack-using-queues/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 225. Implement Stack using Queues

Implement a last-in-first-out (LIFO) stack using only two queues. The implemented stack should support all the functions of a normal stack (`push`, `top`, `pop`, and `empty`).

Implement the `MyStack` class:

 - `void push(int x)` Pushes element x to the top of the stack.
 - `int pop()` Removes the element on the top of the stack and returns it.
 - `int top()` Returns the element on the top of the stack.
 - `boolean empty()` Returns `true` if the stack is empty, `false` otherwise.

**Notes:**

- You must use **only** standard operations of a queue, which means that only `push to back`, `peek/pop from front`, `size` and `is empty` operations are valid.
- Depending on your language, the queue may not be supported natively. You may simulate a queue using a list or deque (double-ended queue) as long as you use only a queue's standard operations.

### Example 1:

```
Input
["MyStack", "push", "push", "top", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 2, 2, false]

Explanation
 - MyStack myStack = new MyStack();
 - myStack.push(1);
 - myStack.push(2);
 - myStack.top();   // return 2
 - myStack.pop();   // return 2
 - myStack.empty(); // return False
```

### Constraints:

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `top`, and `empty`.
- All the calls to `pop` and `top` are valid.

#### Follow-up: 

Can you implement the stack using only one queue?

### Related Topics

- Stack
- Design
- Queue
  
---

# 解題方向

使用兩個 `queues` 或一個 `queue` 其實差不多，直接針對 follow-up 的問題來討論：

 - 只要在每次塞入新資料後，都把之前的資料在塞回 `queue` 尾巴即可

示意圖：

```
假設一個 queue 內已經有 1,2,3,4 四筆資料，當我們要塞入第五筆資料 5 後：

  +---+---+---+---+---+      +---+---+---+---+---+                    +---+---+---+---+---+
  | 4 | 3 | 2 | 1 | 5 |  =>  | 4 | 3 | 2 | 1 | 5 |                =>  | 5 | 4 | 3 | 2 | 1 |
  +---+---+---+---+---+      +---+---+---+---+---+ ^   ^   ^   ^      +---+---+---+---+---+
  |---------------|            |   |   |   |       |   |   |   |
     length = 4                └---+---+---+-------┘   |   |   |
                                   └---+---+-----------┘   |   |
                                       └---+---------------┘   |
                                           └-------------------┘
```

因此，每次寫入新資料時的時間複雜度為 `O(n)`，取出資料時為 `O(1)`