from typing import List

class Solution:
    # Time Complexity: O(logn)
    # Space Complexity: O(1)
    def nextGreatestLetter(self, letters: List[str], target: str) -> str:
        if target >= letters[-1]:
            return letters[0]
        
        low, high = 0, len(letters)-1

        while low < high:
            mid = low + (high-low)//2
            if letters[mid] <= target:
                low = mid + 1
            else:
                high = mid
        
        return letters[low]

if __name__ == '__main__':
    s = Solution()

    assert s.nextGreatestLetter(["c","f","j"],"a") == "c"
    assert s.nextGreatestLetter(["c","f","j"],"c") == "f"
    assert s.nextGreatestLetter(["c","f","j"],"d") == "f"