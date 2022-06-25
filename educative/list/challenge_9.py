'''
Rearrange Sorted List in Max/Min Form

Problem Statement:
    Implement a function called max_min(lst) which will re-arrange the elements of a sorted list such that the 0th index will have the largest number, the 1st index will have the smallest, and the 2nd index will have second-largest, and so on. 
    In other words, all the even-numbered indices will have the largest numbers in the list in descending order and the odd-numbered indices will have the smallest numbers in ascending order.

Input:
    A sorted list

Sample Input:
    lst = [1,2,3,4,5]

Output:
    A list with elements stored in max/min form

Sample Output:
    lst = [5,1,4,2,3]
'''

# My Solution
# Time Complexity: O(n(n-1)/2) -> O(n^2)
def max_min_sol1(lst):
    if len(lst) < 2:
        return lst
    
    idx = 0
    while idx < len(lst):
        maximum, maxIdx = float('-inf'), idx
        for curr in range (idx, len(lst)):
            if lst[curr] > maximum:
                maximum = lst[curr]
                maxIdx = curr
        lst[idx], lst[maxIdx] = lst[maxIdx], lst[idx]
        idx += 1

        if idx >= len(lst):
            continue
        
        minimum, minIdx = float('inf'), idx
        for curr in range (idx, len(lst)):
            if minimum > lst[curr]:
                minimum = lst[curr]
                minIdx = curr
        lst[idx], lst[minIdx] = lst[minIdx], lst[idx]
        idx += 1

    return lst

print(max_min_sol1([1, 2, 3, 4, 5, 6, 7]))

# Solution 1: Creating a new list
# Time Complexity: O(n/2) -> O(n)
def max_min_sol2(lst):
    result = []

    # iterate half left
    for cur in range (len(lst)//2):
        # append corresponding last element
        result.append(lst[-(cur+1)])
        # append current element
        result.append(lst[cur])
    
    # if middle value then append
    if len(lst) % 2 == 1:
        result.append(lst[len(lst)//2])

    return result

print(max_min_sol2([1, 2, 3, 4, 5, 6, 7]))

# Solution 2: Using O(1) extra space
# Time Complexity: O(2n) -> O(n) 
def max_min_sol3(lst):
    if len(lst) == 0:
        return []

    minIdx, maxIdx = 0, len(lst)-1
    modulus = lst[len(lst)-1]+1
    
    for i in range (len(lst)):
        if i % 2 == 0:
            lst[i] += (lst[maxIdx] % modulus) * modulus
            maxIdx -= 1
        else:
            lst[i] += (lst[minIdx] % modulus) * modulus
            minIdx += 1

    for i in range (len(lst)):
        lst[i] = lst[i] // modulus

    return lst

print(max_min_sol3([1, 2, 3, 4, 5, 6, 7]))