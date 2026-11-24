package main

import "testing"

func TestQueue(t *testing.T) {
	q := NewQueue(3)
	if _, ok := q.Dequeue(); ok {
		t.Error("dequeue on empty must fail")
	}
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	if q.Enqueue(4) {
		t.Error("enqueue on full must fail")
	}
	if v, _ := q.Front(); v != 1 {
		t.Errorf("Front = %d, want 1", v)
	}
	if v, _ := q.Dequeue(); v != 1 {
		t.Errorf("Dequeue = %d, want 1 (FIFO)", v)
	}
	// Wrap-around: the freed slot is reused.
	if !q.Enqueue(4) {
		t.Error("enqueue after dequeue must succeed")
	}
	want := []int{2, 3, 4}
	for _, w := range want {
		if v, _ := q.Dequeue(); v != w {
			t.Errorf("Dequeue = %d, want %d", v, w)
		}
	}
	if q.Size() != 0 {
		t.Errorf("Size = %d, want 0", q.Size())
	}
}

func TestWrapStress(t *testing.T) {
	q := NewQueue(2)
	for i := 0; i < 10; i++ {
		if !q.Enqueue(i) {
			t.Fatalf("enqueue %d failed", i)
		}
		if v, _ := q.Dequeue(); v != i {
			t.Fatalf("got %d, want %d", v, i)
		}
	}
}
