package gcollections

import (
	"container/list"
	"sync"

	"github.com/spf13/cast"
)

type Deque struct {
	data *list.List
	mut  *sync.Mutex
}

func NewDeque() (dq *Deque) {
	dq = new(Deque)
	dq.data = list.New()
	dq.mut = new(sync.Mutex)
	return dq
}

//Append push back a item
func (dq *Deque) Append(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushBack(v)
}

// AppendLeft push front a item
func (dq *Deque) AppendLeft(v interface{}) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	dq.data.PushFront(v)
}

//Pop pop a item from back, if dequeue
func (dq *Deque) Pop() (v interface{}, canPop bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		return elem.Value, true
	}
	return nil, false
}

// PopInt pop a int, canPop is true if dq is not empty, error if the item is not int
func (dq *Deque) PopInt() (v int, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		v, err = cast.ToIntE(v)
		return v, true, err
	}
	return 0, false, nil
}

// PopBool pop a int, canPop is true if dq is not empty, error if the item is not bool
func (dq *Deque) PopBool() (v bool, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		v, err = cast.ToBoolE(v)
		return v, true, err
	}
	return false, false, nil
}

// PopFloat pop a int, canPop is true if dq is not empty, error if the item is not float
func (dq *Deque) PopFloat() (v float64, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		v, err = cast.ToFloat64E(v)
		return v, true, err
	}
	return 0, false, nil
}

// PopString pop a int, canPop is true if dq is not empty, error if the item is not string
func (dq *Deque) PopString() (v string, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Back()
		dq.data.Remove(elem)
		v, err = cast.ToStringE(v)
		return v, true, err
	}
	return v, false, nil
}

//PopLeft pop a item from front
func (dq *Deque) PopLeft() (v interface{}, canGet bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		return elem.Value, true
	} else {
		return nil, false
	}
}

// PopLeftInt pop a int left, canPop is true if dq is not empty, err is nil if the item is int
func (dq *Deque) PopLeftInt() (v int, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		v, err = cast.ToIntE(v)
		return v, true, err
	}
	return 0, false, nil
}

// PopLeftBool pop a bool left， canPop is true if dq is not empty, err is nil if the item is bool
func (dq *Deque) PopLeftBool() (v bool, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		v, err = cast.ToBoolE(v)
		return v, true, err
	}
	return false, false, nil
}

// PopLeftFloat pop a float left, canPop is true if dq is not empty, err is nil if the item is float
func (dq *Deque) PopLeftFloat() (v float64, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		v, err = cast.ToFloat64E(v)
		return v, true, err
	}
	return 0, false, nil
}

// PopLeftString pop a string left, canPop is true if dq is not empty, err is nil if the item is string
func (dq *Deque) PopLeftString() (v string, canPop bool, err error) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() > 0 {
		elem := dq.data.Front()
		dq.data.Remove(elem)
		v, err = cast.ToStringE(v)
		return v, true, err
	}
	return v, false, nil
}

//Clear clear a dequeue
func (dq *Deque) Clear() {
	dq.data = list.New()
}

// Remove remove the first elem match v
func (dq *Deque) Remove(v interface{}) (removed bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if dq.data.Len() == 0 {
		return false
	}
	find := false
	current := dq.data.Front()
	var needDelete *list.Element
	for i := 0; i < dq.data.Len(); i++ {
		if current.Value == v {
			find = true
			needDelete = current
			break
		} else {
			current = current.Next()
		}
	}
	if find {
		dq.data.Remove(needDelete)
		return true
	} else {
		return false
	}
}

// Index return the first index between start(include) and end(not include) which match v, find will be false if can not find
func (dq *Deque) Index(v interface{}, start int, end int) (position int, find bool) {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	if start >= end || start < 0 || start >= dq.data.Len() {
		return 0, false
	} else {
		for current, i := dq.data.Front(), 0; current != nil; current, i = current.Next(), i+1 {
			if current.Value == v && start <= i && end > i {
				return i, true
			}
		}
		return 0, false
	}
}

//ExtendLeft push front a container/list
func (dq *Deque) ExtendLeft(other *Deque) {
	if dq != other {
		dq.mut.Lock()
		defer dq.mut.Unlock()
		dq.data.PushFrontList(other.data)
	}
}

//Extend push back a container/list
func (dq *Deque) Extend(other *Deque) {
	if dq != other {
		dq.mut.Lock()
		defer dq.mut.Unlock()
		dq.data.PushBackList(other.data)
	}
}

// Rotate right shift step, if step < 0 then left shift step
func (dq *Deque) Rotate(step int) {
	if dq.data.Len() == 0 {
		return
	}
	if step > 0 {
		for ; step > 0; step-- {
			v, _ := dq.Pop()
			dq.AppendLeft(v)
		}
	} else {
		for ; step < 0; step++ {
			v, _ := dq.PopLeft()
			dq.Append(v)
		}
	}
}

// Size return length of deque
func (dq *Deque) Size() int {
	dq.mut.Lock()
	defer dq.mut.Unlock()
	return dq.data.Len()
}
