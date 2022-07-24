from secrets import choice

class RandomizedSet:

    def __init__(self):
        self.lookup = {}
        self.nums = []

    def insert(self, val: int) -> bool:
        if val in self.lookup:
            return False
        
        self.nums.append(val)
        self.lookup[val] = len(self.nums)-1

        return True

    def remove(self, val: int) -> bool:
        if val not in self.lookup:
            return False

        val_idx = self.lookup[val]
        last_val = self.nums[-1]
        self.lookup[last_val] = val_idx
        self.nums[val_idx] = last_val

        del self.lookup[val]
        self.nums.pop()

        return True

    def getRandom(self) -> int:
        return choice(self.nums)

if __name__ == '__main__':
    obj = RandomizedSet()
    assert obj.insert(1) == True
    assert obj.remove(2) == False
    assert obj.insert(2) == True
    obj.getRandom()
    assert obj.remove(1) == True
    assert obj.insert(2) == False
    assert obj.getRandom() == 2