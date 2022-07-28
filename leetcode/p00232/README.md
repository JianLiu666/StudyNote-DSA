[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/implement-queue-using-stacks/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 232. Implement Queue using Stacks

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

Implement the `MyQueue` class:

- `void push(int x)` Pushes element x to the back of the queue.
- `int pop()` Removes the element from the front of the queue and returns it.
- `int peek()` Returns the element at the front of the queue.
- `boolean empty()` Returns `true` if the queue is empty, `false` otherwise.

**Notes:**

- You must use **only** standard operations of a stack, which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.
- Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
 

### Example 1:

```
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
- MyQueue myQueue = new MyQueue();
- myQueue.push(1); // queue is: [1]
- myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
- myQueue.peek();  // return 1
- myQueue.pop();   // return 1, queue is [2]
- myQueue.empty(); // return false
```

### Constraints:

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `peek`, and `empty`.
- All the calls to `pop` and `peek` are valid.

#### Follow-up: 

Can you implement the queue such that each operation is **amortized** `O(1)` time complexity? In other words, performing `n` operations will take overall `O(n)` time even if one of those operations may take longer.

### Related Topics

- Stack
- Design
- Queue
  
---

# 解題方向

暖身題等級，來討論 amortized O(1) 的問題，我們可用兩個 `stacks` 來實現，第一個 `stack` `s1` 負責維護 input ordering，另一個 `s2` 負責維護取出時維持 FIFO 的規定，原理如下：

```
假設一個 queue 內已經有 1,2,3,4 四筆資料，當我們要塞入第五筆資料 5 時會長這樣

     +---+---+---+---+    +---+
  <- | 1 | 2 | 3 | 4 | <- | 5 |
     +---+---+---+---+    +---+

如果從 stack 的視角來看會像這樣

  +---+
  | 5 |
  +---+
    |
    v
  +---+
  | 4 |
  +---+
  | 3 |
  +---+
  | 2 |
  +---+
  | 1 |
  +---+

當我們現在想取出第一筆資料 1 時，我們可以透過兩個 stacks 這樣維護 queue 的順序

  stack2.push(stack1.pop())

  +---+                                  +---+ 
  | 5 |                                  | 1 | -> got first element
  +---+                                  +---+ 
  | 4 |                                  | 2 | 
  +---+                                  +---+ 
  | 3 |  ->            =>            ->  | 3 | 
  +---+                                  +---+ 
  | 2 |                                  | 4 | 
  +---+                                  +---+ 
  | 1 |                                  | 5 | 
  +---+                                  +---+ 
  stack1     stack2           stack1     stack2
```

如上圖所示，只要在 `stack2` 內沒有資料的時候才需要把 `stack1` 倒進 `stack2`，所以平均而言會是 `O(1)` 的操作