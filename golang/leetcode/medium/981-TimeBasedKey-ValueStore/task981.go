package main

type Value struct {
	value     string
	timestamp int
}

type TimeMap struct {
	data map[string]*[]Value
}

func Constructor() TimeMap {
	data := make(map[string]*[]Value)
	return TimeMap{data: data}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	data, ok := this.data[key]

	if !ok {
		data = new([]Value)
		this.data[key] = data
	}

	*data = append(*data, Value{value: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	dataRow, ok := this.data[key]

	if !ok {
		return ""
	}

	data := *dataRow

	l, r := 0, len(data)-1

	var index int
	for l <= r {
		m := (l + r) / 2

		if data[m].timestamp <= timestamp {
			index = m
			l = m + 1
		} else {
			r = m - 1
		}
	}

	if data[index].timestamp <= timestamp {
		return data[index].value
	} else {
		return ""
	}
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
