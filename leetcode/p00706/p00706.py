class HashMap_BruteForce:

    def __init__(self):
        self.arr = [-1] * 1000001

    def put(self, key: int, value: int) -> None:
        self.arr[key] = value

    def get(self, key: int) -> int:
        return self.arr[key]

    def remove(self, key: int) -> None:
        self.arr[key] = -1

if __name__ == '__main__':
    obj = HashMap_BruteForce()
    
    obj.put(1,1)
    obj.put(2,2)
    assert obj.get(1) == 1
    assert obj.get(3) == -1
    obj.put(2,1)
    assert obj.get(2) == 1
    obj.remove(2)
    assert obj.get(2) == -1
