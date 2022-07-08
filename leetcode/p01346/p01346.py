from typing import List

class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        need = {}
        for num in arr:
            if num in need:
                return True
            
            if num % 2 == 0:
                need[num/2] = 1
            need[num*2] = 1
        
        return False
                
if __name__ == '__main__':
    s = Solution()

    assert s.checkIfExist([10,2,5,3]) == True
    assert s.checkIfExist([7,1,14,11]) == True
    assert s.checkIfExist([3,1,14,11]) == False