package queue

import (
    // the errors library can be used to define and throw an error
    "errors"
)

type Queue struct {
    // our queue will be an empty interface slice
    // empty interfaces can hold any primitive type (int, float, string, etc),
    // thus making our queues more flexible
    member []interface{}
}

/*
    a *Queue means that we have initialized a new copy of the Queue struct, and 
    are expecting to fill it with our own custom values. Passing around this 
    *Queue will then refer to whatever queue is passed back from NewQueue()
*/
func NewQueue() *Queue {
    return &Queue{member: make([]interface{}, 0)}
}

/*
    A q *Queue before the Add declaration specifies that we are operating NOT
    on the library level, but on the returned queue level. We are not calling
    queue.Add(), we're calling q.Add(), where q is a queue returned from 
    NewQueue().
*/
func (q *Queue) Add(val interface{}) {
    q.member = append(q.member, val)
}

func (q *Queue) Remove() (interface{}, error) {
    // remove the first member - but test if q.member is 0 sized first
    if len(q.member) == 0 {
        // create a new error and return it
        return nil, errors.New("Remove from empty queue")
    }
    val := q.member[0]
    q.member = append(q.member[:0], q.member[1:]...)
    return val, nil
}
