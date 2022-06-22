'''
Problem Statement
    Implement a function findMinimum(lst) which finds the smallest number in the given list.

Input
    A list of integers

Sample Input
    arr = [9,2,3,6]

Output
    The smallest number in the list

Sample Output
    2
'''

# My Solution
# Time Complexity: O(n)
def find_minimum_sol1(arr):
    if len(arr) == 0:
        return 0
    
    minimum = arr[0]
    for val in arr:
        if minimum > val:
            minimum = val
    
    return minimum

print(find_minimum_sol1([9, 2, 3, 6]))