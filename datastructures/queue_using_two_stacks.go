package datastructures

import (
	"strconv"
	"strings"
)

// https://www.hackerrank.com/challenges/queue-using-two-stacks/problem?isFullScreen=false

type QueueUsingStacksStack struct {
	items []int32
}

func (stack *QueueUsingStacksStack) push(value int32) {
	stack.items = append(stack.items, value)
}

func (stack *QueueUsingStacksStack) pop() int32 {
	n := len(stack.items)
	if n == 0 {
		return -1
	}

	res := stack.items[n-1]
	stack.items = stack.items[0 : n-1]
	return res
}

func (stack *QueueUsingStacksStack) peek() int32 {
	return stack.items[len(stack.items)-1]
}

func (stack *QueueUsingStacksStack) isEmpty() bool {
	return len(stack.items) == 0
}

type QueueUsingStacksQueue struct {
	active int
	stack1 *QueueUsingStacksStack
	stack2 *QueueUsingStacksStack
}

func NewQueueUsingStacksQueue() *QueueUsingStacksQueue {
	return &QueueUsingStacksQueue{
		active: 1,
		stack1: &QueueUsingStacksStack{},
		stack2: &QueueUsingStacksStack{},
	}
}

func (queue *QueueUsingStacksQueue) enqueue(value int32) {
	var source *QueueUsingStacksStack
	var target *QueueUsingStacksStack

	if queue.active == 1 {
		source = queue.stack1
		target = queue.stack2
		queue.active = 2
	} else {
		source = queue.stack2
		target = queue.stack1
		queue.active = 1
	}

	source.push(value)
	for !source.isEmpty() {
		target.push(source.pop())
	}
}

func (queue *QueueUsingStacksQueue) dequeue() int32 {
	if queue.active == 1 {
		return queue.stack1.pop()
	} else {
		return queue.stack2.pop()
	}
}

func (queue *QueueUsingStacksQueue) print() int32 {
	if queue.active == 1 {
		return queue.stack1.peek()
	} else {
		return queue.stack2.peek()
	}
}

func queueUsingTwoStacks(operations []string) []int32 {
	var result []int32
	queue := NewQueueUsingStacksQueue()

	for _, op := range operations {
		splittedOp := strings.Split(op, " ")
		operation := splittedOp[0]

		if operation == "1" {
			param := splittedOp[1]
			value, _ := strconv.ParseInt(param, 10, 64)
			queue.enqueue(int32(value))
		} else if operation == "2" {
			queue.dequeue()
		} else if operation == "3" {
			result = append(result, queue.print())
		}
	}
	return result
}
