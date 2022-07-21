[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/design-hashset/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 705. Design HashSet

Design a HashSet without using any built-in hash table libraries.

Implement `MyHashSet` class:

- `void add(key)` Inserts the value `key` into the HashSet.
- `bool contains(key)` Returns whether the value `key` exists in the HashSet or not.
- `void remove(key)` Removes the value `key` in the HashSet. If `key` does not exist in the HashSet, do nothing.

### Example 1:

```
Input
["MyHashSet", "add", "add", "contains", "contains", "add", "contains", "remove", "contains"]
[[], [1], [2], [1], [3], [2], [2], [2], [2]]
Output
[null, null, null, true, false, null, true, null, false]

Explanation
 - MyHashSet myHashSet = new MyHashSet();
 - myHashSet.add(1);      // set = [1]
 - myHashSet.add(2);      // set = [1, 2]
 - myHashSet.contains(1); // return True
 - myHashSet.contains(3); // return False, (not found)
 - myHashSet.add(2);      // set = [1, 2]
 - myHashSet.contains(2); // return True
 - myHashSet.remove(2);   // set = [1]
 - myHashSet.contains(2); // return False, (already removed)
```

### Constraints:

- 0 <= key <= $10^6$
- At most $10^4$ calls will be made to `add`, `remove`, and `contains`.

### Related Topics

- Array
- Hash Table
- Linked List
- Design
- Hash Function
  
---

# 解題方向

### Solved using brute-force concept

題目說了資料範圍落在 `0 - 1,000,000` 之間，所以可以直接宣告一個足夠大的 `array[1000001]` 解決，我就爛，讚啦！

### Solved using multiplicative hashing

筆記在此 [Hash Table](./../../notes/hash-table.md)

- A : 給定一個夠大的質數: $3*2^n+1$
- M : 給定的空間，題目表示至多只會有 10,000 次的 function call，所以我選了 $2^13$ 作為空間，沒啥特別原因
- W : 雖然通常是用 size word 作為常數使用，但只要比 M 還大即可，所以我選了 $2^20$，也沒啥特別原因

#### Avoid collision in hash table by using open addressing technique

施工中 ...

- References
  - [LeetCode Discuss | [Python] Easy Multiplicative Hash, explained](https://leetcode.com/problems/design-hashset/discuss/768659/Python-Easy-Multiplicative-Hash-explained)
  - [LeetCode Discuss | Real Python Solution, no cheating, open addressing](https://leetcode.com/problems/design-hashset/discuss/210494/Real-Python-Solution-no-cheating-open-addressing)


