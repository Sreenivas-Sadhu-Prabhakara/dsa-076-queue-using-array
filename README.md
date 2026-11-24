# Implement Queue Using an Array

> Striver SDE Sheet 076 · Day 13 Stack & Queue — a circular buffer: head index plus count, arithmetic modulo capacity, nothing ever shifts. Go, stdlib only, tested.

## Problem

A fixed-capacity FIFO queue over an array: Enqueue, Dequeue, Front, Size, with full/empty reported. The naive array queue shifts every element on dequeue (O(n)) or leaks the front slots forever — the exercise is avoiding both.

## Approach

Treat the array as a ring: the back lives at `(head + size) % capacity`, and dequeue just advances `head` modulo capacity, so freed slots are reused automatically and nothing moves. Storing `head + size` (rather than head and tail pointers) kills the classic ambiguity where head == tail could mean either empty or full. This structure is what real systems use for bounded queues — ring buffers in io_uring, audio pipelines, and channel implementations are this exact object.

**Complexity:** all operations O(1); capacity ints of storage.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
