//
// Created by Yaz Saito on 06/15/12.
//

// fifo_queue.Queue is a simple FIFO queue. It supports pushing an
// item at the end, and popping an item from the front. It is
// implemented as a single-linked list of arrays of items.
package fifo_queue

type Item interface{}

const chunkSize = 64

type chunk struct {
	items        [chunkSize]interface{}
	start, limit int
	next         *chunk
}

type Queue struct {
	head, tail *chunk
	count      int
}

// Create a new empty FIFO queue
func NewQueue() *Queue {
	return &Queue{}
}

// Return the number of items in the queue
func (q *Queue) Len() int {
	return q.count
}

// Add an item to the end of the queue
func (q *Queue) PushBack(item Item) {
	if q.head == nil {
		q.tail = new(chunk)
		q.head = q.tail
	} else if q.tail.limit >= chunkSize {
		q.tail.next = new(chunk)
		q.tail = q.tail.next
	}
	q.tail.items[q.tail.limit] = item
	q.tail.limit++
	q.count++
}

// Remove the item at the head of the queue and return it.
//
// REQUIRES: q.Len() > 0
func (q *Queue) PopFront() Item {
	doAssert(q.count > 0)
	doAssert(q.head.start < q.head.limit)
	item := q.head.items[q.head.start]
	q.head.start++
	q.count--
	if q.head.start >= q.head.limit {
		if q.count == 0 {
			q.head.start = 0
			q.head.limit = 0
			q.head.next = nil
		} else {
			q.head = q.head.next
		}
	}
	return item
}

func doAssert(b bool) {
	if !b {
		panic("fifo_queue assertion failed")
	}
}
