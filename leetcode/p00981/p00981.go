package p00981

type Element struct {
	value     string
	timestamp int
}

type TimeMap struct {
	data map[string][]*Element
}

func Constructor() TimeMap {
	return TimeMap{data: map[string][]*Element{}}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, exists := this.data[key]; !exists {
		this.data[key] = []*Element{}
	}
	this.data[key] = append(this.data[key], &Element{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	if _, exists := this.data[key]; !exists {
		return ""
	}
	if this.data[key][0].timestamp > timestamp {
		return ""
	}

	arr := this.data[key]
	head, tail := 0, len(arr)-1
	for head <= tail {
		mid := (head + tail) / 2
		if arr[mid].timestamp == timestamp {
			return arr[mid].value
		}
		if arr[mid].timestamp < timestamp {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}

	return arr[head-1].value
}
