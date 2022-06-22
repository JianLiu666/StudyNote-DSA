'''
Problem Statement
    Implement a function, find_product(lst), which modifies a list so that each index has a product of all the numbers present in the list except the number stored at that index.

Input
    A list of numbers (could be floating points or integers)

Sample Input
    arr = [1,2,3,4]

Output
    A list such that each index has a product of all the numbers in the list except the number stored at that index.

Sample Output
    arr = [24,12,8,6]
'''

# My Solution
# Time Complexity: O(2n)
def find_product_sol1(lst):
    result = []

    # 紀錄 list 乘積與 0 的出現次數
    temp, zero_count = 1, 0
    for num in lst:
        if num != 0:
            temp *= num
        else:
            zero_count += 1
    
    # 兩個 0 以上表示結果全為 0
    if zero_count >= 2:
        return [0] * len(lst)
    
    # 再重新跑過一次 list 計算結果
    for num in lst:
        if zero_count == 0:
            result.append(temp/num)
        elif zero_count != 0 and num == 0:
            result.append(temp)
        else:
            result.append(0)
    
    return result

print(find_product_sol1([1, 2, 3, 4]))

# Solution 2: Optimizing the number of multiplications
# Time Complexity: O(n)
def find_product_sol2(lst):
    # get product start from left
    left = 1
    product = []
    for ele in lst:
        product.append(left)
        left = left * ele
    # get product starting from right
    print(product)
    
    right = 1
    for i in range(len(lst)-1, -1, -1):
        product[i] = product[i] * right
        right = right * lst[i]

    return product


print(find_product_sol2([1, 2, 3, 4]))