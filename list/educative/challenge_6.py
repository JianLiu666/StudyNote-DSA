'''
Problem Statement
    Implement a function find_second_maximum(lst) which returns the second largest element in the list.

Input
    A List

Sample Input
    [9,2,3,6]

Output
    Second largest element in the list

Sample Output
    6
'''

# My Solution
# Time Complexity: O(n)
def find_second_maximum(lst):
    if len(lst) < 2:
        return None

    first, second = lst[0], lst[1]
    if second > first:
        first, second = second, first

    for i in range(2, len(lst)):
        if lst[i] > first:
            second = first
            first = lst[i]
        elif lst[i] > second:
            second = lst[i]
    
    return second