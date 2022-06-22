'''
Problem Statement
    In this problem, you have to implement the find_sum(lst,k) function which will take a number k as input and return two numbers that add up to k.

Input
    A list and a number k

Sample Input
    lst = [1,21,3,14,5,60,7,6]
    k = 81

Output
    A list with two integers a and b that add up to k

Sample Output
    lst = [21,60]
'''

# My Solution: Brute Force
# Time Complexity: O(n(n-1)/2)
def find_sum_sol1(lst, k):
    for n1 in range (len(lst)):
        for n2 in range (n1+1, len(lst)):
            if lst[n1]+lst[n2] == k:
                return [lst[n1], lst[n2]]

print(find_sum_sol1([1, 7, 3, 14, 5, 6, 21, 60], 81))

# Solution 2: Binary search
# Time Complexity: O(nlog(n))
def binary_search(lst, element):
    head = 0
    tail = len(lst)-1
    index = -1
    found = False

    while head <= tail and found == False:
        mid = (head + tail) // 2
        if lst[mid] == element:
            index = mid
            found = True
        else:
            if lst[mid] < element:
                head = mid + 1
            else:
                tail = mid - 1
    
    return index

def find_sum_sol2(lst, k):
    lst.sort()

    for i in range (len(lst)):
        idx = binary_search(lst, k-lst[i])
        if idx != -1 and idx != i:
            return [lst[i], lst[idx]]

print(find_sum_sol2([1, 21, 3, 14, 5, 60, 7, 6], 81))

# Solution 3: Moving indices
# Time Complexity: O(nlog(n))
def find_sum_sol3(lst, k):
    lst.sort()

    head, tail = 0, len(lst)-1
    while head != tail:
        sum = lst[head] + lst[tail]
        if sum == k:
            return [lst[head], lst[tail]]
        else:
            if sum < k:
                head += 1
            else:
                tail -= 1

    return []

print(find_sum_sol3([1, 21, 3, 14, 5, 60, 7, 6], 81))