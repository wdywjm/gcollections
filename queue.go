package gcollections

import (
	"container/list"
	"sync"

	"github.com/spf13/cast"
)

type Queue struct {
	data *list.List
	mut  *sync.Mutex
}

func NewQueue() (queue *Queue) {
	queue = new(Queue)
	queue.data = list.New()
	queue.mut = new(sync.Mutex)
	return
}

func (queue *Queue) Push(v interface{}) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	queue.data.PushFront(v)
}

func (queue *Queue) IsEmpty() (isEmpty bool) {
	if queue.data.Len() > 0 {
		return false
	} else {
		return true
	}
}

// Pop pop a item, canGet is true if queue is not empty
func (queue *Queue) Pop() (v interface{}, canGet bool) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return nil, false
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		return elem.Value, true
	}
}

// PopInt pop a int, canGet is true if queue is not empty, err is nil if the item is int
func (queue *Queue) PopInt() (v int, canGet bool, err error) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return 0, false, nil
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		v, err = cast.ToIntE(elem.Value)
		return v, true, err
	}
}

// PopBool pop a bool, canGet is true if queue is not empty, err is nil if the item is bool
func (queue *Queue) PopBool() (v bool, canGet bool, err error) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return false, false, nil
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		v, err = cast.ToBoolE(elem.Value)
		return v, true, err
	}
}

// PopFloat pop a float, canGet is true if queue is not empty, err is nil if the item is float
func (queue *Queue) PopFloat() (v float64, canGet bool, err error) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return 0, false, nil
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		v, err = cast.ToFloat64E(elem.Value)
		return v, true, err
	}
}

// PopString pop a string, canGet is true if queue is not empty, err is nil if the item is string
func (queue *Queue) PopString() (v string, canGet bool, err error) {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	if queue.IsEmpty() {
		return "", false, nil
	} else {
		elem := queue.data.Back()
		queue.data.Remove(elem)
		v, err = cast.ToStringE(elem.Value)
		return v, true, err
	}
}

// Size return the length of queue
func (queue *Queue) Size() int {
	queue.mut.Lock()
	defer queue.mut.Unlock()
	return queue.data.Len()
}
