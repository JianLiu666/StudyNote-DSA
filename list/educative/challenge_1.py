'''
Remove Even Integers from List

Problem Statement:
    Implement a function that removes all the even elements from a given list.
    Name it remove_even(lst).

Input:
    A list with random integers.

Sample Input:
    my_list = [1,2,4,5,10,6,3]

Output:
    A list with only odd integers.

Sample Output:
    my_list = [1,5,3]
'''

# My Solution: Doing it "by hand"
# Time Complexity: O(n)
def remove_even_sol1(lst):
    odds = []
    for number in lst:
        if number % 2 != 0:
            odds.append(number)
    return odds

print(remove_even_sol1([3, 2, 41, 3, 34]))

# Solution 2: Using list comprehension
# Time Complexity: O(n)
def remove_even_sol2(lst):
    return [number for number in lst if number % 2 != 0]

print(remove_even_sol2([3, 2, 41, 3, 34]))
