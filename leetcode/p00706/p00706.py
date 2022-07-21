class Node:
    def __init__(self, key=0, val=0):
        self.key = key
        self.val = val
class MyHashMap:
    def __init__(self):
        self.buckets = [[] for _ in range(1 << 13)]

    def hasher(self, key: int) -> int:
        return ((key*786433) & (1<<20-1)) >> 7

    def put(self, key: int, value: int) -> None:
        hashed = self.hasher(key)
        
        for i in range(len(self.buckets[hashed])):
            if self.buckets[hashed][i].key == key:
                self.buckets[hashed][i].val = value
                return

        self.buckets[hashed].append(Node(key, value))

    def get(self, key: int) -> int:
        hashed = self.hasher(key)
        for i in range(len(self.buckets[hashed])):
            if self.buckets[hashed][i].key == key:
                return self.buckets[hashed][i].val

        return -1

    def remove(self, key: int) -> None:
        hashed = self.hasher(key)
        
        idx = -1
        for i in range(len(self.buckets[hashed])):
            if self.buckets[hashed][i].key == key:
                idx = i
                break
        
        if idx == -1:
            return
        if idx == 0:
            self.buckets[hashed] = self.buckets[hashed][1:]
        elif idx == len(self.buckets[hashed])-1:
            self.buckets[hashed].pop()
        else:
            self.buckets[hashed] = self.buckets[hashed][:idx] + self.buckets[hashed][idx+1:]

if __name__ == '__main__':
    obj = MyHashMap()
    
    obj.put(1,1)
    obj.put(2,2)
    assert obj.get(1) == 1
    assert obj.get(3) == -1
    obj.put(2,1)
    assert obj.get(2) == 1
    obj.remove(2)
    assert obj.get(2) == -1
