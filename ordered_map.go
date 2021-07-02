package gcollections

import (
	"sync"

	"github.com/spf13/cast"
)

type OrderedMap struct {
	mut  *sync.Mutex
	dict map[interface{}]*LinkedNode
	list *LinkedList
}

type item struct {
	Key   interface{}
	Value interface{}
}

func NewOrderedMap() *OrderedMap {
	om := new(OrderedMap)
	om.dict = make(map[interface{}]*LinkedNode)
	om.list = NewLinkedList()
	om.mut = new(sync.Mutex)
	return om
}

// Set set k-v and push back the k-v, the order will not change if the key already exist unless putEndIfExists is true
func (om *OrderedMap) Set(key interface{}, value interface{}, putEndIfExists bool) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if k, ok := om.dict[key]; ok {
		if putEndIfExists {
			om.list.Remove(k)
			om.dict[key] = om.list.Append(key, value)
		} else {
			om.dict[key].value = value
		}
	} else {
		om.dict[key] = om.list.Append(key, value)
	}
}

func (om *OrderedMap) Get(key interface{}) (interface{}, bool) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if v, ok := om.dict[key]; ok {
		return v.value, true
	}
	return nil, false
}

// GetInt get int value, return error if cast failed
func (om *OrderedMap) GetInt(key interface{}) (int, bool, error) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if _, ok := om.dict[key]; ok {
		res, err := cast.ToIntE(om.dict[key])
		return res, true, err
	}
	return 0, false, nil
}

// GetBool get bool value, return error if cast failed
func (om *OrderedMap) GetBool(key interface{}) (bool, bool, error) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if _, ok := om.dict[key]; ok {
		res, err := cast.ToBoolE(om.dict[key])
		return res, true, err
	}
	return false, false, nil
}

// GetString get string value, return error if cast failed
func (om *OrderedMap) GetString(key interface{}) (string, bool, error) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if _, ok := om.dict[key]; ok {
		res, err := cast.ToStringE(om.dict[key])
		return res, true, err
	}
	return "", false, nil
}

// GetFloat get float value, return error if cast failed
func (om *OrderedMap) GetFloat(key interface{}) (float64, bool, error) {
	om.mut.Lock()
	defer om.mut.Unlock()
	if _, ok := om.dict[key]; ok {
		res, err := cast.ToFloat64E(om.dict[key])
		return res, true, err
	}
	return 0, false, nil
}

// Del del a key from ordered map, return false if key not exist
func (om *OrderedMap) Del(key interface{}) bool {
	om.mut.Lock()
	defer om.mut.Unlock()
	if n, ok := om.dict[key]; ok {
		if ok := om.list.Remove(n); !ok {
			return false
		}
		delete(om.dict, key)
		return true
	}
	return false
}

// Iter return k-v in order(set order)
func (om *OrderedMap) Iter() chan *item {
	om.mut.Lock()
	defer om.mut.Unlock()
	ch := make(chan *item)
	go func() {
		for n := range om.list.Iter() {
			ch <- &item{
				Key:   n.key,
				Value: n.value,
			}
		}
		close(ch)
	}()
	return ch
}
