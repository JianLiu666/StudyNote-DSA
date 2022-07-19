# Array

> An Array is a collection of items. The items could be integers, strings, DVDs, games, books – anything really. The items are stored in neighboring (contiguous) memory locations. Because they’re stored together, checking through the entire collection of items is straightforward.
>
> from [LeetCode Explore: Arrays 101](https://leetcode.com/explore/learn/card/fun-with-arrays/521/introduction/3236/)

- [Array](#array)
  - [What is different between Array Capacity and Array Length?](#what-is-different-between-array-capacity-and-array-length)
    - [Array Capacity](#array-capacity)
  - [Basic Array Operations](#basic-array-operations)
    - [Insertion](#insertion)
      - [Inserting at the End of an Array](#inserting-at-the-end-of-an-array)
      - [Inserting at the Start of an Array](#inserting-at-the-start-of-an-array)
      - [Inserting Anywhere in the Array](#inserting-anywhere-in-the-array)
    - [Deletion](#deletion)
      - [Deleting From the End of an Array](#deleting-from-the-end-of-an-array)
      - [Deleting From the Start of an Array](#deleting-from-the-start-of-an-array)
      - [Deleting From Anywhere in the Array](#deleting-from-anywhere-in-the-array)
    - [Search](#search)
      - [Linear Search](#linear-search)
  - [Concepts](#concepts)
    - [Binary Search](#binary-search)

---
 
## What is different between Array Capacity and Array Length?

If somebody asks you how long the DVD Array is, what would your answer be?

1. The number of DVDs the box could hold, if it was full, or → `The capacity of the Array`
2. The number of DVDs currently in the box. → `The length of the Array`

### Array Capacity

- The Array’s capacity must be decided when the Array is created. *The capacity cannot be changed later.*

## Basic Array Operations

### Insertion

#### Inserting at the End of an Array
- Time Complexity : `O(1)`

```
+---+
| 8 |
+---+
  |
  v
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```

#### Inserting at the Start of an Array
- Time Complexity: `O(n)`
  
```
+---+                                  +---+
| 8 |                                  | 8 |
+---+                                  +---+
  |                                      |  ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐
  v                                      v  | v | v | v | v | v | v | v
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |  =>  | 8 | 1 | 2 | 3 | 4 | 5 | 6 | 7 |
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```
    
#### Inserting Anywhere in the Array
- Time Complexity: `O(n)`

```
                +---+                                  +---+
                | 8 |                                  | 8 |
                +---+                                  +---+
                  |                                      |  ┌-┐ ┌-┐ ┌-┐ 
                  v                                      v  | v | v | v 
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |  =>  | 1 | 2 | 3 | 4 | 8 | 5 | 6 | 7 |
  +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```
    

### Deletion

#### Deleting From the End of an Array
- Time Complexity: `O(1)`

```
                            DELETE!
                              v
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 |   |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```
    
#### Deleting From the Start of an Array
- Time Complexity: `O(n)`

```
DELETE!                                   ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐ ┌-┐
  v                                       v | v | v | v | v | v | v |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  |   | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 2 | 3 | 4 | 5 | 6 | 7 | 8 |   |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```
    
#### Deleting From Anywhere in the Array
- Time Complexity: `O(n)`

```
                DELETE!                                   ┌-┐ ┌-┐ ┌-┐ 
                  v                                       v | v | v | 
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 1 | 2 | 3 | 4 |   | 6 | 7 | 8 |  =>  | 1 | 2 | 3 | 4 | 6 | 7 | 8 |   |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```
    

### Search

#### Linear Search
- Time Complexity: `O(n)`

```
Where is the 7?
                                        NO! NO! NO! NO! NO! NO!                                      GOTCHA!
                                         v   v   v   v   v   v                                          v
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
+---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
```

## Concepts

### Binary Search
- Time Complexity: `O(logn)`

```
施工中 ...
```