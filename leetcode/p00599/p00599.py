from typing import List

class Solution:
    # Time Complexity: O(n)
    # Space Complexity: O(n)
    def findRestaurant(self, list1: List[str], list2: List[str]) -> List[str]:
        memo = {}
        for idx, name in enumerate(list1):
            memo[name] = idx
        
        minimum = float("inf")
        res = []
        for idx, name in enumerate(list2):
            if name in memo:
                l = idx + memo[name]
                if l == minimum:
                    res.append(name)
                elif l < minimum:
                    res = [name]
                    minimum = l
        
        return res


if __name__ == '__main__':
    s = Solution()

    assert s.findRestaurant(["Shogun","Tapioca Express","Burger King","KFC"], ["Piatti","The Grill at Torrey Pines","Hungry Hunter Steakhouse","Shogun"]) == ["Shogun"]
    assert s.findRestaurant(["Shogun","Tapioca Express","Burger King","KFC"], ["KFC","Shogun","Burger King"]) == ["Shogun"]
    assert s.findRestaurant(["Shogun","Tapioca Express","Burger King","KFC"], ["KFC","Burger King","Tapioca Express","Shogun"]) == ["KFC","Burger King","Tapioca Express","Shogun"]