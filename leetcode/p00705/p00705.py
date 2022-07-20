class HashSet_BruteForce:
    def __init__(self):
        self.arr = [False] * 1000001

    def add(self, key: int) -> None:
        self.arr[key] = True

    def remove(self, key: int) -> None:
        self.arr[key] = False

    def contains(self, key: int) -> bool:
        return self.arr[key]

if __name__ == '__main__':
    obj = HashSet_BruteForce()
    
    obj.add(1)
    obj.add(2)
    assert obj.contains(1) == True
    assert obj.contains(3) == False
    obj.add(2)
    assert obj.contains(2) == True
    obj.remove(2)
    assert obj.contains(2) == False
