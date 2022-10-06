[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/time-based-key-value-store/)
![image](https://img.shields.io/badge/Difficulty-Medium-yellow)

---

# 981. Time Based Key-Value Store

Design a time-based key-value data structure that can store multiple values for the same key at different time stamps and retrieve the key's value at a certain timestamp.

Implement the `TimeMap` class:

- `TimeMap()` Initializes the object of the data structure.
- void `set(String key, String value, int timestamp)` Stores the key `key` with the value `value` at the given time `timestamp`.
- `String get(String key, int timestamp)` Returns a value such that `set` was called previously, with `timestamp_prev <= timestamp`. If there are multiple such values, it returns the value associated with the largest `timestamp_prev`. If there are no values, it returns `""`.

### Example 1:

```
Input
["TimeMap", "set", "get", "get", "set", "get", "get"]
[[], ["foo", "bar", 1], ["foo", 1], ["foo", 3], ["foo", "bar2", 4], ["foo", 4], ["foo", 5]]
Output
[null, null, "bar", "bar", null, "bar2", "bar2"]

Explanation
 - TimeMap timeMap = new TimeMap();
 - timeMap.set("foo", "bar", 1);  // store the key "foo" and value "bar" along with timestamp = 1.
 - timeMap.get("foo", 1);         // return "bar"
 - timeMap.get("foo", 3);         // return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
 - timeMap.set("foo", "bar2", 4); // store the key "foo" and value "bar2" along with timestamp = 4.
 - timeMap.get("foo", 4);         // return "bar2"
 - timeMap.get("foo", 5);         // return "bar2"
```

### Constraints:

- `1 <= key.length, value.length <= 100`
- `key` `and` value consist of lowercase English letters and digits.
- 1 <= `timestamp` <= $10^7$
- All the timestamps `timestamp` of `set` are strictly increasing.
- At most 2 * $10^5$ calls will be made to `set` and `get`.

### Related Topics

- Hash Table
- String
- Binary Search
- Design
  
---

# 解題方向

先用 `Hash Table` 維護不同的 `key`，接著對於 `key` 相同的 `values` 用 `array` 來維護，由於題目說明 `timstamp` 會保證遞增，所以可以放心的將新的 `value` 直接塞到 `array` 最後一個位置即可，不需重新排序

最後只要實作 `Get()` 時的查找即可，透過 `Binary Search` 加速收斂搜尋區間
