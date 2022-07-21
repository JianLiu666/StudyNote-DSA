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

**Dfinition** : Using the first *m* bits of the *key* times an irrational number.

**Formal Definition** : $h(k) = ⌊M(kA\ mod\ 1)⌋$
- *M* is usually an integer $2^m$
- *A* is an irrational number (or an approximation thereto) in the range $0 < A < 1$
  - [Knuth](https://en.wikipedia.org/wiki/Donald_Knuth) suggests that *A* can be ${\sqrt{5}-1 \over 2} ≈ 0.6180339887...$
- The modulo 1 operation removes the integer part of $k*A$

```
The multiplication method of hashing:

                                          w bits
                              |-                         -|
                              +---------------------------+
                              |             k             |
                              +---------------------------+
                              +---------------------------+
                           x  |        s = A x 2^w        |
                              +---------------------------+
-----------------------------------------------------------
+---------------------------+ +---------------------------+
|            r1             | |            r0             |
+---------------------------+ +---------------------------+
                              |-      -| extract m bits
                                 h(k)
```

As the above image, we only concerned with the most significant bits of *r0*, so the formula will be like:

- $h(k) = ⌊(kA\ mod\ 2^w) / 2^{w-m}⌋$
- ususally, we will suppose that the word size of the machine is *w* bits and that *k* fits into a single word.

**References**

- [Wikipedia | Multiplicative Hashing](https://en.wikipedia.org/wiki/Hash_function#Multiplicative_hashing)
- [Hash Table : Intro](http://alrightchiu.github.io/SecondRound/hash-tableintrojian-jie.html)

## Avoiding collisions in hash table

施工中 ...

### Chaining

施工中 ...

### Open Addressing

施工中 ...

**References**

- [Hash Table : Open Addressing](http://alrightchiu.github.io/SecondRound/hash-tableopen-addressing.html)