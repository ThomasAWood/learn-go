package tests

import (
	"learngo/algorithms"
	"testing"
)

func TestEmptyPeek(t *testing.T) {
	queue := algorithms.Queue{}
	result := queue.Peek()
	if result != -1 {
		t.Errorf(`Peek() on empty queue = %v, expected -1`, result)
	}
}

func TestEmptyDequeue(t *testing.T) {
	queue := algorithms.Queue{}
	result := queue.Dequeue()
	if result != -1 {
		t.Errorf(`Dequeue() on empty queue = %v, expected -1`, result)
	}
}

func TestEnqueuePeek(t *testing.T) {
	queue := algorithms.Queue{}
	queue.Enqueue(3)
	result := queue.Peek()
	if result != 3 {
		t.Errorf(`Enqueue(3), Peek() = %v, expected 3`, result)
	}
}

func TestEnqueueDequeue(t *testing.T) {
	queue := algorithms.Queue{}
	queue.Enqueue(5)
	result := queue.Dequeue()
	if result != 5 {
		t.Errorf(`Enqueue(5), Dequeue() = %v, expected 5`, result)
	}
}

func TestMultipleEnqueue(t *testing.T) {
	queue := algorithms.Queue{}
	for i := 0; i < 10; i++ {
		queue.Enqueue(7)
	}
	for i := 0; i < 10; i++ {
		result := queue.Dequeue()
		if result != 7 {
			t.Errorf(`Enqueue(7) * 10, Dequeue() = %v`, result)
		}
	}
}

// TODO: think of more exhaustive tests
