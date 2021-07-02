package gcollections

import (
	"sync"
	"time"

	"github.com/spf13/cast"
)

type Pair struct {
	Key   interface{} // key add to counter
	Value int         // store counts of the key
}

type PairList []Pair

type Counter struct {
	count map[interface{}]int
	mut   *sync.Mutex
}

// NewCounterFromMap new counter from map<key><count>
func NewCounterFromMap(m map[interface{}]int) *Counter {
	if m == nil {
		m = make(map[interface{}]int)
	}
	return &Counter{
		count: m,
		mut:   new(sync.Mutex),
	}
}

// NewCounterFromSlice new counter from slice
func NewCounterFromSlice(s []interface{}) *Counter {
	c := &Counter{
		count: make(map[interface{}]int),
		mut:   new(sync.Mutex),
	}
	for i := range s {
		c.count[s[i]] += 1
	}
	return c
}

// Add add a elem(key) into the counter and incr the counts of the key
func (c *Counter) Add(elem interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()
	c.count[elem] += 1
}

// GetCount get counts of a element
func (c *Counter) GetCount(elem interface{}) int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return c.count[elem]
}

// Elements return elements of the counter
func (c *Counter) Elements() []interface{} {
	c.mut.Lock()
	defer c.mut.Unlock()
	elements := make([]interface{}, 0)
	for k, v := range c.count {
		for i := 0; i < v; i++ {
			elements = append(elements, k)
		}
	}
	return elements
}

// IntElements return int elements of the counter
func (c *Counter) IntElements() ([]int, error) {
	return cast.ToIntSliceE(c.Elements())
}

// BoolElements return bool elements of the counter
func (c *Counter) BoolElements() ([]bool, error) {
	return cast.ToBoolSliceE(c.Elements())
}

// StringElements return string elements of the counter
func (c *Counter) StringElements() ([]string, error) {
	return cast.ToStringSliceE(c.Elements())
}

// DurationElements return duration elements of the counter
func (c *Counter) DurationElements() ([]time.Duration, error) {
	return cast.ToDurationSliceE(c.Elements())
}

//MostCommon return top frequently keys and their counts
func (c *Counter) MostCommon(top int) PairList {
	c.mut.Lock()
	defer c.mut.Unlock()
	if top > len(c.count) || top < 0 {
		top = len(c.count)
	}

	p := make(PairList, 0)
	pq := NewPriorityQueue()

	for k, v := range c.count {
		pq.PushItem(&Item{
			Value:    k,
			Priority: v,
		})
	}

	for i := 0; i < top; i++ {
		item := pq.PopItem()
		p = append(p, Pair{
			Key:   item.Value,
			Value: item.Priority,
		})
	}
	return p
}

// Del delete a key from a counter, if key exist return true else return false
func (c *Counter) Del(key interface{}) bool {
	c.mut.Lock()
	defer c.mut.Unlock()
	if _, ok := c.count[key]; ok {
		delete(c.count, key)
		return true
	}
	return false
}

// Len return length of a counter
func (c *Counter) Len() int {
	c.mut.Lock()
	defer c.mut.Unlock()
	return len(c.count)
}

// Clear reset all counts
func (c *Counter) Clear() {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.count = make(map[interface{}]int)
}
