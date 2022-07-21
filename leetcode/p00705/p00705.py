class MyHashSet:
    def __init__(self):
        self.buckets = [[] for _ in range(1<<13)]

    def hasher(self, key):
        return ((key*786433) & (1<<20)-1) >> 7
    
    def add(self, key: int) -> None:
        hashed = self.hasher(key)
        if key not in self.buckets[hashed]:
            self.buckets[hashed].append(key)

    def remove(self, key: int) -> None:
        hashed = self.hasher(key)
        if key in self.buckets[hashed]:
            self.buckets[hashed].remove(key)

    def contains(self, key: int) -> bool:
        hashed = self.hasher(key)
        return key in self.buckets[hashed]

if __name__ == '__main__':
    obj = MyHashSet()
    
    obj.add(1)
    obj.add(2)
    assert obj.contains(1) == True
    assert obj.contains(3) == False
    obj.add(2)
    assert obj.contains(2) == True
    obj.remove(2)
    assert obj.contains(2) == False
