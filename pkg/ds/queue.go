package ds

import (
	"container/list"
	"fmt"
	"github.com/pkg/errors"
)

// linked list
// 1
func bax() {
	// new linked list
	queue := list.New()

	// Simply append to enqueue.
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)

	// Dequeue
	front := queue.Front()
	fmt.Println(front.Value)
	queue.Remove(front)
}

// 2
type customQueue struct {
	queue *list.List
}

func (c *customQueue) Enqueue(value string) {
	c.queue.PushBack(value)
}

func (c *customQueue) Dequeue() error {
	if c.queue.Len() > 0 {
		ele := c.queue.Front()
		c.queue.Remove(ele)
	}
	return fmt.Errorf("pop Error: Queue is empty")
}

func (c *customQueue) Front() (string, error) {
	if c.queue.Len() > 0 {
		if val, ok := c.queue.Front().Value.(string); ok {
			return val, nil
		}
		return "", fmt.Errorf("peep Error: Queue Datatype is incorrect")
	}
	return "", fmt.Errorf("peep Error: Queue is empty")
}

func (c *customQueue) Size() int {
	return c.queue.Len()
}

func (c *customQueue) Empty() bool {
	return c.queue.Len() == 0
}

// slice
// 1

func enqueueSimple(queue []int, element int) []int {
	queue = append(queue, element) // Simply append to enqueue.
	fmt.Println("Enqueued:", element)
	return queue
}
func dequeueSimple(queue []int) (int, []int) {
	element := queue[0] // The first element is the one to be dequeued.
	if len(queue) == 1 {
		var tmp = []int{}
		return element, tmp

	}
	return element, queue[1:]
}

func doo() {
	var queue = make([]int, 0)
	queue = enqueueSimple(queue, 10)
	// ...

}

// 2

type Queue struct {
	Elements []int
	Size     int
}

func (q *Queue) Enqueue(elem int) {
	if q.GetLength() == q.Size {
		fmt.Println("Overflow")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue() int {
	if q.IsEmpty() {
		fmt.Println("UnderFlow")
		return 0
	}
	element := q.Elements[0]
	if q.GetLength() == 1 {
		q.Elements = nil
		return element
	}
	q.Elements = q.Elements[1:]
	return element // Slice off the element once it is dequeued.
}

func (q *Queue) GetLength() int {
	return len(q.Elements)
}

func (q *Queue) IsEmpty() bool {
	return len(q.Elements) == 0
}

func (q *Queue) Peek() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("empty queue")
	}
	return q.Elements[0], nil
}
