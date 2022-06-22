'''
Problem Statement

    Implement a function that merges two sorted lists of m and n elements respectively, into another sorted list.
    Name it merge_lists(lst1, lst2).

Input
    Two sorted lists.

Sample Input
    list1 = [1,3,4,5]  
    list2 = [2,6,7,8]

Output
    A merged and sorted list consisting of all elements of both input lists.

Sample Output
    arr = [1,2,3,4,5,6,7,8]
'''

# My Solution: Merging in Place
# Time Complexity: O(m+n-1)
def merge_lists(lst1, lst2):
    # Creating 2 new variable to track the 'current index'
    idx1, idx2 = 0, 0
    iCount = 0
    
    # iterate lst2 把裡面的 element 逐一與 lst1 比較，如果比 lst1[idx1] 還小時就往前插入
    while idx1 < len(lst1) and idx2 < len(lst2):
        if lst1[idx1] > lst2[idx2]:
            lst1.insert(idx1, lst2[idx2])
            idx2 += 1
        
        idx1 += 1
        iCount += 1

    # 確保 lst2 的資料都有處理完畢
    if idx2 < len(lst2):
        lst1 += lst2[idx2:]
    
    return iCount, lst1


print(merge_lists([4, 5, 6], [-2, -1, 0, 7]))