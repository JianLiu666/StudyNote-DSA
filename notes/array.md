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
      - [Binary Search](#binary-search)
  - [Concepts](#concepts)
    - [Two-Pointer Techniques](#two-pointer-techniques)
      - [Head and tail pointers](#head-and-tail-pointers)
      - [Different step pointers](#different-step-pointers)
    - [In-place Algorithms](#in-place-algorithms)

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

#### Binary Search
- Time Complexity: `O(logn)`

```
施工中 ...
```

## Concepts

### Two-Pointer Techniques

Using `two pointers at the same time` to do the iteration.

#### Head and tail pointers

For example : 

 - reverse the elements in an array.

```
We can set two pointers to the given array
 - head pointer : starts from the first element.
 - tail pointer : starts from the last element.

 head                        tail
  v                           v
+---+---+---+---+---+---+---+---+
| 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |
+---+---+---+---+---+---+---+---+

Continue swapping the elements until the two pointers meet each other.

 - step.1
                                            
    head                        tail        -> head                tail <-
     v                           v              v                   v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 |  =>  | 8 | 2 | 3 | 4 | 5 | 6 | 7 | 1 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
     ^                           ^
     └---------------------------┘

 - step.2
                                            
        head                tail                -> head        tail <-
         v                   v                      v           v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 8 | 2 | 3 | 4 | 5 | 6 | 7 | 1 |  =>  | 8 | 7 | 3 | 4 | 5 | 6 | 2 | 1 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
         ^                   ^
         └-------------------┘

 - step.3
                                            
            head        tail                        -> headtail <-
             v           v                              v   v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 8 | 7 | 3 | 4 | 5 | 6 | 2 | 1 |  =>  | 8 | 7 | 6 | 4 | 5 | 3 | 2 | 1 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
             ^           ^
             └-----------┘

 - step.4
                                            
                headtail                              Congrats!
                 v   v                                    
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 8 | 7 | 3 | 4 | 5 | 6 | 2 | 1 |  =>  | 8 | 7 | 6 | 5 | 4 | 3 | 2 | 1 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
                 ^   ^
                 └---┘
```

#### Different step pointers

For example : 

 - Given an array and a value, remove all instances of that value **in-place** and return the new length.

```
Let's remove all 5 from this array, we can set two pointers to achieve it.
 - one is still used for the array iteration
 - the second one always points at the position for next legal value.

the pseudocode is :

  if nums[iter] != val {
    nums[anchor] = nums[iter]
    anchor += 1
  }

 - step.1

    iter                                    -> iter
     v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 5 | 2 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 5 | 2 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
     ^                                          ^
   anchor (replace to 1)                   -> anchor
   
 - step.2

        iter                                    -> iter
         v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 5 | 2 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 5 | 2 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
         ^                                      ^
       anchor (skip)                          anchor

 - step.3

            iter                                    -> iter
             v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 5 | 2 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 2 | 2 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
         ^                                          ^
       anchor (replace to 2)                   -> anchor
   
 - step.4

                iter                                    -> iter
                 v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 2 | 2 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 2 | 2 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
             ^                                      ^
           anchor (skip)                          anchor

 - step.5

                    iter                                    -> iter
                     v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 2 | 2 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 2 | 3 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
             ^                                          ^
           anchor (replace to 3)                   -> anchor

 - step.6

                        iter                                    -> iter
                         v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 2 | 3 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 2 | 3 | 5 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
                 ^                                      ^
               anchor (skip)                          anchor
            
 - step.7

                            iter                                    -> iter
                             v                                          v
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
   | 1 | 2 | 3 | 5 | 3 | 5 | 4 | 5 |  =>  | 1 | 2 | 3 | 4 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+      +---+---+---+---+---+---+---+---+
                 ^                                          ^
               anchor (replace to 4)                   -> anchor

 - step.8

                                iter
                                 v  
   +---+---+---+---+---+---+---+---+
   | 1 | 2 | 3 | 4 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+
                 ^                  
               anchor (skip)  

 - eventually result

    we only wanted
   |---------------|
   +---+---+---+---+---+---+---+---+
   | 1 | 2 | 3 | 4 | 3 | 5 | 4 | 5 |
   +---+---+---+---+---+---+---+---+
                 ^
               anchor
```

### In-place Algorithms

```
施工中 ...
```