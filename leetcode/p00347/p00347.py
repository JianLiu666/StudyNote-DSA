from typing import Counter, List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        counter = Counter(nums)
        buckets = [[] for _ in range(len(nums)+1)]

        for num, count in counter.items():
            buckets[count].append(num)

        res = []
        for i in range(len(buckets)-1, -1, -1):
            if len(buckets[i]) > 0:
                res += buckets[i]
                if len(res) >= k:
                    return res
        
        return res

if __name__ == '__main__':
    s = Solution()

    assert s.topKFrequent([1,1,1,2,2,3], 2) == [1,2]
    assert s.topKFrequent([1], 1) == [1]
    assert s.topKFrequent([-1,-1], 1) == [-1]
    assert s.topKFrequent([1,2], 2) == [1,2]