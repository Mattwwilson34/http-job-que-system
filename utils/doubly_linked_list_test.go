package utils

import (
	"http-job-que-system/types"
	"testing"
	"time"
)

// Helper function to create test jobs
func createTestJob(id, name string) types.Job {
	return types.Job{
		Id:              id,
		Name:            name,
		Body:            "",
		CreatedDateTime: time.Now().UTC().Format(time.RFC3339),
	}
}

func TestIsEmpty(t *testing.T) {
	dll := CreateNewDoubleLinkedList()

	if !dll.IsEmpty() {
		t.Errorf("new list should be empty, got IsEmpty() = %t", dll.IsEmpty())
	}

	if dll.size != 0 {
		t.Errorf("new list should have size 0, got %d", dll.size)
	}
}

func TestAppend(t *testing.T) {
	t.Run("append to empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Verify initial state
		if !dll.IsEmpty() || dll.size != 0 {
			t.Fatal("new list should be empty with size 0")
		}

		// Append first job
		job1 := createTestJob("1", "job_1")
		node1 := dll.Append(job1)

		// Test return value
		if node1 == nil {
			t.Error("Append should return the created node, got nil")
		}

		// Test list state after first append
		if dll.size != 1 {
			t.Errorf("size after first append = %d, want 1", dll.size)
		}

		if dll.IsEmpty() {
			t.Error("list should not be empty after append")
		}

		// Test head and tail point to same node
		if dll.head != dll.tail {
			t.Error("head and tail should point to same node in single-item list")
		}

		// Test job content
		if got := dll.head.value.Id; got != job1.Id {
			t.Errorf("head job ID = %s, want %s", got, job1.Id)
		}

		// Test node pointers for single item
		if dll.head.next != nil || dll.head.prev != nil {
			t.Error("single node should have nil next and prev pointers")
		}
	})

	t.Run("append to non-empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Add first job
		job1 := createTestJob("1", "job_1")
		dll.Append(job1)

		// Add second job
		job2 := createTestJob("2", "job_2")
		node2 := dll.Append(job2)

		// Test return value
		if node2 == nil {
			t.Error("Append should return the created node, got nil")
		}

		// Test list size
		if dll.size != 2 {
			t.Errorf("size after second append = %d, want 2", dll.size)
		}

		// Test head remains unchanged
		if got := dll.head.value.Id; got != job1.Id {
			t.Errorf("head job ID = %s, want %s (should remain job1)", got, job1.Id)
		}

		// Test tail updated to new job
		if got := dll.tail.value.Id; got != job2.Id {
			t.Errorf("tail job ID = %s, want %s", got, job2.Id)
		}

		// Test head and tail are different
		if dll.head == dll.tail {
			t.Error("head and tail should be different nodes in two-item list")
		}

		// Test forward linking (head -> tail)
		if dll.head.next != dll.tail {
			t.Errorf("head.next should point to tail, got %p want %p", dll.head.next, dll.tail)
		}

		// Test backward linking (tail -> head)
		if dll.tail.prev != dll.head {
			t.Errorf("tail.prev should point to head, got %p want %p", dll.tail.prev, dll.head)
		}

		// Test boundary conditions
		if dll.head.prev != nil {
			t.Error("head.prev should be nil")
		}

		if dll.tail.next != nil {
			t.Error("tail.next should be nil")
		}
	})

	t.Run("append multiple jobs", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		jobs := []types.Job{
			createTestJob("1", "job_1"),
			createTestJob("2", "job_2"),
			createTestJob("3", "job_3"),
		}

		// Append all jobs
		for i, job := range jobs {
			node := dll.Append(job)

			if node == nil {
				t.Errorf("Append %d should return node, got nil", i+1)
			}

			expectedSize := i + 1
			if dll.size != expectedSize {
				t.Errorf("size after append %d = %d, want %d", i+1, dll.size, expectedSize)
			}
		}

		// Verify final state
		if dll.size != 3 {
			t.Errorf("final size = %d, want 3", dll.size)
		}

		// Verify head is first job
		if got := dll.head.value.Id; got != jobs[0].Id {
			t.Errorf("head job ID = %s, want %s", got, jobs[0].Id)
		}

		// Verify tail is last job
		if got := dll.tail.value.Id; got != jobs[2].Id {
			t.Errorf("tail job ID = %s, want %s", got, jobs[2].Id)
		}

		// Verify linking integrity by traversing forward
		current := dll.head
		for i, expectedJob := range jobs {
			if current == nil {
				t.Fatalf("node %d is nil during forward traversal", i)
			}

			if got := current.value.Id; got != expectedJob.Id {
				t.Errorf("forward traversal node %d ID = %s, want %s", i, got, expectedJob.Id)
			}

			current = current.next
		}

		// Should end with nil
		if current != nil {
			t.Error("forward traversal should end with nil")
		}

		// Verify linking integrity by traversing backward
		current = dll.tail
		for i := len(jobs) - 1; i >= 0; i-- {
			if current == nil {
				t.Fatalf("node %d is nil during backward traversal", i)
			}

			if got := current.value.Id; got != jobs[i].Id {
				t.Errorf("backward traversal node %d ID = %s, want %s", i, got, jobs[i].Id)
			}

			current = current.prev
		}

		// Should end with nil
		if current != nil {
			t.Error("backward traversal should end with nil")
		}
	})
}

func TestAppendEdgeCases(t *testing.T) {
	t.Run("append zero value job", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		var zeroJob types.Job

		node := dll.Append(zeroJob)

		if node == nil {
			t.Error("should handle zero-value jobs and return node")
		}

		if dll.size != 1 {
			t.Errorf("size = %d, want 1", dll.size)
		}

		// Zero-value job should have empty string values
		if dll.head.value.Id != "" || dll.head.value.Name != "" {
			t.Error("zero-value job should have empty string fields")
		}
	})

	t.Run("append to nil receiver panics", func(t *testing.T) {
		var dll *DoublyLinkedList

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when calling Append on nil receiver")
			}
		}()

		job := createTestJob("1", "test")
		dll.Append(job)
	})
}
