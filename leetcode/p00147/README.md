[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/insertion-sort-list/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 147. Insertion Sort List

Given the `head` of a singly linked list, sort the list using **insertion sort**, and return the sorted list's head.

The steps of the **insertion sort** algorithm:

1. Insertion sort iterates, consuming one input element each repetition and growing a sorted output list.
2. At each iteration, insertion sort removes one element from the input data, finds the location it belongs within the sorted list and inserts it there.
3. It repeats until no input elements remain.

The following is a graphical example of the insertion sort algorithm. The partially sorted list (black) initially contains only the first element in the list. One element (red) is removed from the input data and inserted in-place into the sorted list with each iteration.

![image](./image/Insertion-sort-example-300px.gif)

### Example 1:

![image](./image/sort1linked-list.jpeg)

```
Input: head = [4,2,1,3]
Output: [1,2,3,4]
```

### Example 2:

![image](./image/sort2linked-list.jpeg)

```
Input: head = [-1,5,3,4,0]
Output: [-1,0,3,4,5]
```

### Constraints:

- The number of nodes in the list is in the range `[1, 5000]`.
- `-5000 <= Node.val <= 5000`

### Related Topics

- Linked List
- Sorting
  
---

# 解題方向

`Linked List` 暖身題，控制好指針紀錄即可，直接看程式碼吧