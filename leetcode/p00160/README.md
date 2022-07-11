[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/intersection-of-two-linked-lists/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 160. Intersection of Two Linked Lists

Given the heads of two singly linked-lists `headA` and `headB`, return the node at which the two lists intersect. If the two linked lists have no intersection at all, return `null`.

For example, the following two linked lists begin to intersect at node `c1`:

![image](./image/160_statement.png)

The test cases are generated such that there are no cycles anywhere in the entire linked structure.

**Note** that the linked lists must **retain their original structure** after the function returns.

**Custom Judge:**

The inputs to the **judge** are given as follows (your program is **not** given these inputs):

- `intersectVal` - The value of the node where the intersection occurs. This is `0` if there is no intersected node.
- `listA` - The first linked list.
- `listB` - The second linked list.
- `skipA` - The number of nodes to skip ahead in `listA` (starting from the head) to get to the intersected node.
- `skipB` - The number of nodes to skip ahead in `listB` (starting from the head) to get to the intersected node.

The judge will then create the linked structure based on these inputs and pass the two heads, `headA` and `headB` to your program. If you correctly return the intersected node, then your solution will be **accepted**.

### Example 1:

![image](./image/160_example_1_1.png)

```
Input: intersectVal = 8, listA = [4,1,8,4,5], listB = [5,6,1,8,4,5], skipA = 2, skipB = 3
Output: Intersected at '8'

Explanation: 
 - The intersected node's value is 8 (note that this must not be 0 if the two lists intersect).
 - From the head of A, it reads as [4,1,8,4,5].
 - From the head of B, it reads as [5,6,1,8,4,5].
 - There are 2 nodes before the intersected node in A; There are 3 nodes before the intersected node in B.
```

### Example 2:

![image](./image/160_example_2.png)

```
Input: intersectVal = 2, listA = [1,9,1,2,4], listB = [3,2,4], skipA = 3, skipB = 1
Output: Intersected at '2'

Explanation:
 - The intersected node's value is 2 (note that this must not be 0 if the two lists intersect).
 - From the head of A, it reads as [1,9,1,2,4].
 - From the head of B, it reads as [3,2,4]. 
 - There are 3 nodes before the intersected node in A; There are 1 node before the intersected node in B.
```

### Example 3:

![image](./image/160_example_3.png)

```
Input: intersectVal = 0, listA = [2,6,4], listB = [1,5], skipA = 3, skipB = 2
Output: No intersection

Explanation:
 - From the head of A, it reads as [2,6,4]. 
 - From the head of B, it reads as [1,5]. 
 - Since the two lists do not intersect, intersectVal must be 0, while skipA and skipB can be arbitrary values.

Explanation:
 - The two lists do not intersect, so return null.
```

### Constraints:

- The number of nodes of `listA` is in the `m`.
- The number of nodes of `listB` is in the `n`.
- 1 <= m, n <= 3 * $10^4$
- 1 <= Node.val <= $10^5$
- `0 <= skipA < m`
- `0 <= skipB < n`
- `intersectVal` is `0` if `listA` and `listB` do not intersect.
- `intersectVal == listA[skipA] == listB[skipB]` if `listA` and `listB` intersect.

#### Follow up: 

Could you write a solution that runs in `O(m + n)` time and use only `O(1)` memory?

### Related Topics

- Hash Table
- Linked List
- Two Pointers
  
---

# 解題方向

### Solved using length alignment concept

題目表明是 Singly Linked List，所以想要從尾巴往前找是不可能的了，但可以把想法轉換成：**將底部固定，從頭開始逐一比對**。

我們想像一個場景：如果兩條分岔路的後半段是重疊的，我們讓 `A` 與 `B` 在各自**與終點距離相同**的不同入口**同時出發**，只要速度始終保持一致，他們必定會在交會處相遇，如下圖所示：

```
下面是兩個不同的 Linked Lists，以 A, B 分別表示

A :       head -> a1 -> a2 -> a3
                                \
                                |-> c1 -> c2 -> c3 -> none
                                /
B : head -> b1 -> b2 -> b3 -> b4

如果我們直接讓 pointer_A 與 pointer_B 從各自的 head 開始出發，那可想而知肯定是 pointer_A 先抵達終點
但如果我們調整一下，讓 pointer_B 從 b2 處開始出發，那兩個 pointers 移動三次後將會在 c1 處交會

概念很簡單，但還是讓我畫一下怎麼畫圖XD

                  pa0   pa1   pa2
A :       head -> a1 -> a2 -> a3    pb3
                                \   pa3
                                |-> c1 -> c2 -> c3 -> none
                                /
B : head -> b1 -> b2 -> b3 -> b4
                  pb0   pb1   pb2
```

概念講完，實作流程如下：

1. 先分別對 `headA` 與 `headB` 實際走訪一次取得兩個 Linked Lists 的實際長度 `lenA`, `lenB`
2. 比較 `lenA` 與 `lenB` 的大小，將長度較大一邊的 `pointer` 移動到與較短一邊距離相等的位置
3. 兩邊同時往下走訪，直到相遇或抵達終點 (i.e., 彼此不重疊)

### Solved using cycle detection concept

延續上面的概念，我們把握一個重點：**距離相等**。

換句話說，只要 `pointer A` 與 `pointer B` 兩個指針行走的距離相等，它們如果不是在交會處會合，不然就是一起抵達終點對吧？

因此我們只要思考如何讓兩個指針行走的距離變成一樣就好，即：

- `Pointer A` 的總行走距離等於 `A` + `B`
- `Pointer B` 的總行走距離等於 `B` + `A`

小畫家來了：

```
當 Linked Lists 重疊時的情境 : pointer_a11 = pointer_b11 = node_c1

                  pb8   pb9   pb10  pb11
                  pa0   pa1   pa2   pa11
A :       head -> a1 -> a2 -> a3    pb4   pb5   pb6   pb7
                                \   pa3   pa4   pa5   pa6
                                |-> c1 -> c2 -> c3 -> none
                                /
B : head -> b1 -> b2 -> b3 -> b4
            pb0   pb1   pb2   pb3
            pa7   pa8   pa9   pa10

當 Linked List 不重疊時的情境 : pointer_a8 = pointer_b8 = none

            pb5   pb6   pb7   pb8
            pa0   pa1   pa2   pa3
A : head -> a1 -> a2 -> a3 -> none

B : head -> b1 -> b2 -> b3 -> b4 -> none
            pb0   pb1   pb2   pb3   pb4
            pa4   pa5   pa6   pa7   pa8

```