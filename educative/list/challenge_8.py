'''
Rearrange Positive & Negative Values

Problem Statement:
    Implement a function rearrange(lst) which rearranges the elements such that all the negative elements appear on the left and positive elements appear at the right of the list. 
    Note that it is not necessary to maintain the sorted order of the input list.

    Generally zero is NOT positive or negative, we will treat zero as a positive integer for this challenge! So, zero will be placed at the right.

Sample Input:
    [10,-1,20,4,5,-9,-6]

Output:
    A list with negative elements at the left and positive elements at the right

Sample Output:
    [-1,-9,-6,10,20,4,5]
'''

# My Solution
# Time Complexity: O(n)
def rearrange_sol1(lst):
    l, r = 0, len(lst)-1
    
    while l < r:
        if lst[l] >= 0 and lst[r] < 0:
            lst[l], lst[r] = lst[r], lst[l]
            l += 1
            r -= 1
        elif lst[l] < 0:
            l += 1
        elif lst[r] >= 0:
            r -= 1
    
    return lst

print(rearrange_sol1([10, -1, 20, 4, 5, -9, -6]))

# Solution 2: Rearranging in Place
# Time Complexity: O(n)
def rearrange_sol2(lst):
    idx = 0

    for cur in range (len(lst)):
        if lst[cur] < 0:
            lst[cur], lst[idx] = lst[idx], lst[cur]
            idx += 1
    
    return lst

print(rearrange_sol2([10, -1, 20, 4, 5, -9, -6]))
