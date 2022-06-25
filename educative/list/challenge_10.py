'''
Maximum Sum Sublist

Problem statement:
    Given an integer list, return the maximum sublist sum. 
    The list may contain both positive and negative integers and is unsorted.

Input:
    a list lst

Sample input:
    lst = [-4, 2, -5, 1, 2, 3, 6, -5, 1]

Output:
    a number (maximum subarray sum)

Sample output:
    largest_sum = 12
'''

# My Solution
# Time Complexity: O(n)
def find_max_sum_sublist_sol1(lst): 
  if len(lst) == 0:
    return 0

  result = lst[0]
  idx = 0

  # 找到起始位置
  for i in range (len(lst)):
    if lst[i] >= 0:
      result = lst[i]
      idx = i
      break
    
    if lst[i] > result:
      result = lst[i]
      idx = i

  # 整個 list 都是負數的話就不用在找了
  if result < 0 or idx == len(lst)-1:
    return result

  temp = result
  cur = idx+1

  # 從定位點開始往後查找
  while cur < len(lst):
    # 確保累加這個數字是有意義的，不會越加越少
    if lst[idx] + lst[cur] > 0:
      temp += lst[cur]
    # 不然就跳到下一個位置重新累計
    elif cur+2 < len(lst):
      temp = lst[cur+1]
      cur += 2
      idx += 2
      continue

    if temp > result:
      result = temp
    
    cur += 1
    idx += 1

  return result

print(find_max_sum_sublist_sol1([-4, 2, -5, 1, 2, 3, 6, -5, 1]));

# Solution: Kadane's Algorithm
# Time Complexity: O(n)
def find_max_sum_sublist_sol2(lst): 
    if len(lst) == 0:
        return 0

    cur_max, global_max = lst[0], lst[0]
    for i in range (1, len(lst)):
        if cur_max < 0:
            cur_max = lst[i]
        else:
            cur_max += lst[i]

        if cur_max > global_max:
            global_max = cur_max

    return global_max


print(find_max_sum_sublist_sol2([-4, 2, -5, 1, 2, 3, 6, -5, 1]));
