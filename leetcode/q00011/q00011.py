from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(1)
    def maxArea(self, height: List[int]) -> int:
        head = 0
        tail = len(height)-1

        max_area = 0
        while head != tail:
            h = min(height[head], height[tail])
            w = tail - head
            area = h * w
            if area > max_area:
                max_area = area

            if height[head] < height[tail]:
                head += 1
            else:
                tail -= 1

        return max_area

if __name__ == '__main__':
    s = Solution()

    assert s.maxArea([1,8,6,2,5,4,8,3,7]) == 49
    assert s.maxArea([1,1]) == 1