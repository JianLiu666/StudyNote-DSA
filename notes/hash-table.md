# Hash Table

施工中 ...

--- 

- [Hash Table](#hash-table)
  - [Hash Function](#hash-function)
    - [What makes a good hash function?](#what-makes-a-good-hash-function)
    - [Division Method](#division-method)
    - [Multiplication Method](#multiplication-method)
  - [Avoiding collisions in hash table](#avoiding-collisions-in-hash-table)
    - [Chaining](#chaining)
    - [Open Addressing](#open-addressing)

---

## Hash Function

### What makes a good hash function?

> *from Introduction to Algorithms Ch.11 Hash Tables*

A good hash function satisfies (approximately) the assumption of simple uniform hashing:
- each key is equally likely to hash to any of the *m* slots, independently of where any other key has hashedt to.  

Unfortunately, we typically have no way to check this condition, since we rarely know the probability distribution from which the keys are drawn. Moreover, the keys might not be drawn independently.

**Example:**

If we know that the keys are random real numbers *k* independently and uniformly distributed in the range $0 \leq k < 1$, then the hash function can easily be:

- $h(k) = ⌊km⌋$

In practice, we can often employ heuristic techniques to create a hash function that performs well. Qualitative information about the distribution of keys may be useful in this design process.

### Division Method

施工中 ...

### Multiplication Method

施工中 ...

## Avoiding collisions in hash table

施工中 ...

### Chaining

施工中 ...

### Open Addressing

施工中 ...