'''
Right Rotate List

Problem Statement:
    Implement a function right_rotate(lst, k) which will rotate the given list by k.
    This means that the right-most elements will appear at the left-most position in the list and so on.
    You only have to rotate the list by one element at a time.

Input:
    A list and a positive number by which to rotate that list

Sample Input:
    lst = [10,20,30,40,50]
    k = 3

Output:
    The given list rotated by k elements

Sample Output:
    lst = [30,40,50,10,20]
'''

# My Solution
# Time Complexity: O(n)
def right_rotate_sol1(lst, k):
    len_lst = len(lst)

    if len_lst < 2 or len_lst == k:
        return lst

    idx = k
    if k > len_lst:
        idx = k % len_lst
    
    return lst[len_lst-idx:] + lst[:len_lst-idx]

print(right_rotate_sol1([10, 20, 30, 40, 50], 3))

# Solution 2: Pythonic Rotation
# Time Complexity: O(n)
def right_rotate_sol2(lst, k):
    if len(lst) == 0:
        k = 0
    else:
        k = k % len(lst)
    
    return lst[-k:] + lst[:-k]

print(right_rotate_sol2([10, 20, 30, 40, 50], 3))
