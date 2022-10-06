package p00981

type Info struct {
	Value     string
	Timestamp int
}

type TimeMap struct {
	items map[string][]*Info
}

func Constructor() TimeMap {
	return TimeMap{
		items: map[string][]*Info{},
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.items[key] == nil {
		this.items[key] = []*Info{}
	}
	this.items[key] = append(this.items[key], &Info{
		Value:     value,
		Timestamp: timestamp,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if arr, ok := this.items[key]; ok {
		head, tail := 0, len(arr)-1
		for head < tail {
			mid := (head + tail) / 2
			if arr[mid].Timestamp == timestamp {
				return arr[mid].Value
			}
			if arr[mid].Timestamp > timestamp {
				tail = mid
			} else {
				head = mid + 1
			}
		}

		if arr[tail].Timestamp > timestamp {
			if tail == 0 {
				return ""
			}
			return arr[tail-1].Value
		}
		return arr[tail].Value
	}

	return ""
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
