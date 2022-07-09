[![image](https://img.shields.io/badge/Leetcode-Link-blue?logo=leetcode)](https://leetcode.com/problems/excel-sheet-column-number/)
![image](https://img.shields.io/badge/Difficulty-Easy-green)

---

# 171. Excel Sheet Column Number

Given a string `columnTitle` that represents the column title as appears in an Excel sheet, return its corresponding column number.

For example:

```
A -> 1
B -> 2
C -> 3
...
Z -> 26
AA -> 27
AB -> 28 
...
```

### Example 1:

```
Input: columnTitle = "A"
Output: 1
```

### Example 2:

```
Input: columnTitle = "AB"
Output: 28
```

### Example 3:

```
Input: columnTitle = "ZY"
Output: 701
```

### Constraints:

- `1 <= columnTitle.length <= 7`
- `columnTitle` consists only of uppercase English letters.
- `columnTitle` is in the range `["A", "FXSHRXW"]`.

### Related Topics

- Math
- String

---

# 解題方向

基本的字串操作練習

| Binary   | Oct | Dec | Hex | Glyph |
| :------: | :-: | :-: | :-: | :---: |
| 100 0001 | 101 | 65  | 41  | A     |
| 100 0010 | 102 | 66  | 42  | B     |
| ...      | ... | ... | ... | ...   |
| 101 1010 | 132 | 90  | 5A  | Z     |