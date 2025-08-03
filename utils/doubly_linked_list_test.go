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

func TestPrepend(t *testing.T) {
	t.Run("prepend to empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Verify initial state
		if !dll.IsEmpty() || dll.size != 0 {
			t.Fatal("new list should be empty with size 0")
		}

		// Prepend first job
		job1 := createTestJob("1", "job_1")
		node1 := dll.Prepend(job1)

		// Test return value
		if node1 == nil {
			t.Error("Prepend should return the created node, got nil")
		}

		// Test list state after first prepend
		if dll.size != 1 {
			t.Errorf("size after first prepend = %d, want 1", dll.size)
		}

		if dll.IsEmpty() {
			t.Error("list should not be empty after prepend")
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

		// Verify returned node is actually the head
		if node1 != dll.head {
			t.Error("returned node should be the head")
		}
	})

	t.Run("prepend to non-empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Add first job (will be tail after prepend)
		job1 := createTestJob("1", "job_1")
		dll.Append(job1)

		// Prepend second job (will become new head)
		job2 := createTestJob("2", "job_2")
		node2 := dll.Prepend(job2)

		// Test return value
		if node2 == nil {
			t.Error("Prepend should return the created node, got nil")
		}

		// Test list size
		if dll.size != 2 {
			t.Errorf("size after prepend = %d, want 2", dll.size)
		}

		// Test head updated to new job
		if got := dll.head.value.Id; got != job2.Id {
			t.Errorf("head job ID = %s, want %s", got, job2.Id)
		}

		// Test tail remains original job
		if got := dll.tail.value.Id; got != job1.Id {
			t.Errorf("tail job ID = %s, want %s (should remain job1)", got, job1.Id)
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

		// Verify returned node is actually the new head
		if node2 != dll.head {
			t.Error("returned node should be the new head")
		}
	})

	t.Run("prepend multiple jobs", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		jobs := []types.Job{
			createTestJob("1", "job_1"),
			createTestJob("2", "job_2"),
			createTestJob("3", "job_3"),
		}

		// Prepend all jobs (they'll be in reverse order)
		for i, job := range jobs {
			node := dll.Prepend(job)

			if node == nil {
				t.Errorf("Prepend %d should return node, got nil", i+1)
			}

			expectedSize := i + 1
			if dll.size != expectedSize {
				t.Errorf("size after prepend %d = %d, want %d", i+1, dll.size, expectedSize)
			}

			// Each prepend should make the new node the head
			if node != dll.head {
				t.Errorf("prepend %d: returned node should be new head", i+1)
			}
		}

		// Verify final state
		if dll.size != 3 {
			t.Errorf("final size = %d, want 3", dll.size)
		}

		// Verify head is last prepended job (job_3)
		if got := dll.head.value.Id; got != jobs[2].Id {
			t.Errorf("head job ID = %s, want %s", got, jobs[2].Id)
		}

		// Verify tail is first prepended job (job_1)
		if got := dll.tail.value.Id; got != jobs[0].Id {
			t.Errorf("tail job ID = %s, want %s", got, jobs[0].Id)
		}

		// Verify order is reversed: [job_3, job_2, job_1]
		expectedOrder := []string{"3", "2", "1"}

		// Forward traversal should show reversed order
		current := dll.head
		for i, expectedId := range expectedOrder {
			if current == nil {
				t.Fatalf("node %d is nil during forward traversal", i)
			}

			if got := current.value.Id; got != expectedId {
				t.Errorf("forward traversal node %d ID = %s, want %s", i, got, expectedId)
			}

			current = current.next
		}

		// Should end with nil
		if current != nil {
			t.Error("forward traversal should end with nil")
		}

		// Backward traversal should show original order: [job_1, job_2, job_3]
		expectedReverseOrder := []string{"1", "2", "3"}
		current = dll.tail
		for i, expectedId := range expectedReverseOrder {
			if current == nil {
				t.Fatalf("node %d is nil during backward traversal", i)
			}

			if got := current.value.Id; got != expectedId {
				t.Errorf("backward traversal node %d ID = %s, want %s", i, got, expectedId)
			}

			current = current.prev
		}

		// Should end with nil
		if current != nil {
			t.Error("backward traversal should end with nil")
		}
	})
}

func TestPrependEdgeCases(t *testing.T) {
	t.Run("prepend zero value job", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		var zeroJob types.Job

		node := dll.Prepend(zeroJob)

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

		if node != dll.head {
			t.Error("returned node should be the head")
		}
	})

	t.Run("prepend to nil receiver panics", func(t *testing.T) {
		var dll *DoublyLinkedList

		defer func() {
			if r := recover(); r == nil {
				t.Error("expected panic when calling Prepend on nil receiver")
			}
		}()

		job := createTestJob("1", "test")
		dll.Prepend(job)
	})

	t.Run("prepend maintains node isolation", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		job1 := createTestJob("1", "first")
		job2 := createTestJob("2", "second")

		node1 := dll.Prepend(job1)
		node2 := dll.Prepend(job2)

		// Nodes should be different objects
		if node1 == node2 {
			t.Error("different prepends should create different node objects")
		}

		// But linking should be correct
		if node2.next != node1 {
			t.Error("second prepend should link to first")
		}

		if node1.prev != node2 {
			t.Error("first node should link back to second")
		}
	})
}

func TestRemoveFirst(t *testing.T) {
	t.Run("remove from empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Verify initial empty state
		if !dll.IsEmpty() || dll.size != 0 {
			t.Fatal("list should start empty")
		}

		// Remove from empty list
		node := dll.RemoveFirst()

		// Should return nil
		if node != nil {
			t.Errorf("RemoveFirst() on empty list = %v, want nil", node)
		}

		// List should remain empty
		if !dll.IsEmpty() || dll.size != 0 {
			t.Error("list should remain empty after removing from empty list")
		}

		if dll.head != nil || dll.tail != nil {
			t.Error("head and tail should remain nil after removing from empty list")
		}
	})

	t.Run("remove from single node list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		job := createTestJob("1", "single_job")
		dll.Append(job)

		// Verify single node state
		if dll.size != 1 || dll.head != dll.tail {
			t.Fatal("list should have single node")
		}

		// Remove the single node
		node := dll.RemoveFirst()

		// Should return the node
		if node == nil {
			t.Fatal("RemoveFirst() should return the node, got nil")
		}

		// Verify returned node content
		if got := node.value.Id; got != job.Id {
			t.Errorf("returned node ID = %s, want %s", got, job.Id)
		}

		// Verify pointer cleanup on returned node
		if node.next != nil || node.prev != nil {
			t.Error("returned node should have clean pointers (next=nil, prev=nil)")
		}

		// Verify list is now empty
		if !dll.IsEmpty() || dll.size != 0 {
			t.Errorf("list should be empty after single node removal, size=%d", dll.size)
		}

		// Verify head and tail are reset
		if dll.head != nil || dll.tail != nil {
			t.Error("head and tail should be nil after single node removal")
		}
	})

	t.Run("remove from multi node list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		jobs := []types.Job{
			createTestJob("1", "job_1"),
			createTestJob("2", "job_2"),
			createTestJob("3", "job_3"),
		}

		// Build list: [job_1] ↔ [job_2] ↔ [job_3]
		for _, job := range jobs {
			dll.Append(job)
		}

		// Store original head and second node for verification
		originalHead := dll.head
		originalSecond := dll.head.next

		// Remove first node
		node := dll.RemoveFirst()

		// Verify returned node
		if node == nil {
			t.Fatal("RemoveFirst() should return the node, got nil")
		}

		if got := node.value.Id; got != jobs[0].Id {
			t.Errorf("returned node ID = %s, want %s", got, jobs[0].Id)
		}

		// Verify pointer cleanup on returned node
		if node.next != nil || node.prev != nil {
			t.Error("returned node should have clean pointers (next=nil, prev=nil)")
		}

		// Verify returned node is the original head
		if node != originalHead {
			t.Error("returned node should be the original head node")
		}

		// Verify list state
		if dll.size != 2 {
			t.Errorf("list size = %d, want 2", dll.size)
		}

		// Verify new head is the original second node
		if dll.head != originalSecond {
			t.Error("new head should be the original second node")
		}

		// Verify new head content
		if got := dll.head.value.Id; got != jobs[1].Id {
			t.Errorf("new head ID = %s, want %s", got, jobs[1].Id)
		}

		// Verify new head has clean prev pointer
		if dll.head.prev != nil {
			t.Error("new head prev should be nil")
		}

		// Verify tail unchanged
		if got := dll.tail.value.Id; got != jobs[2].Id {
			t.Errorf("tail should remain unchanged, got ID %s, want %s", got, jobs[2].Id)
		}

		// Verify bidirectional linking integrity
		if dll.head.next != dll.tail {
			t.Error("head.next should point to tail")
		}

		if dll.tail.prev != dll.head {
			t.Error("tail.prev should point to head")
		}
	})
}

func TestRemoveLast(t *testing.T) {
	t.Run("remove from empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Verify initial empty state
		if !dll.IsEmpty() || dll.size != 0 {
			t.Fatal("list should start empty")
		}

		// Remove from empty list
		node := dll.RemoveLast()

		// Should return nil
		if node != nil {
			t.Errorf("RemoveLast() on empty list = %v, want nil", node)
		}

		// List should remain empty
		if !dll.IsEmpty() || dll.size != 0 {
			t.Error("list should remain empty after removing from empty list")
		}

		if dll.head != nil || dll.tail != nil {
			t.Error("head and tail should remain nil after removing from empty list")
		}
	})

	t.Run("remove from single node list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		job := createTestJob("1", "single_job")
		dll.Append(job)

		// Verify single node state
		if dll.size != 1 || dll.head != dll.tail {
			t.Fatal("list should have single node")
		}

		// Remove the single node
		node := dll.RemoveLast()

		// Should return the node
		if node == nil {
			t.Fatal("RemoveLast() should return the node, got nil")
		}

		// Verify returned node content
		if got := node.value.Id; got != job.Id {
			t.Errorf("returned node ID = %s, want %s", got, job.Id)
		}

		// Verify pointer cleanup on returned node
		if node.next != nil || node.prev != nil {
			t.Error("returned node should have clean pointers (next=nil, prev=nil)")
		}

		// Verify list is now empty
		if !dll.IsEmpty() || dll.size != 0 {
			t.Errorf("list should be empty after single node removal, size=%d", dll.size)
		}

		// Verify head and tail are reset
		if dll.head != nil || dll.tail != nil {
			t.Error("head and tail should be nil after single node removal")
		}
	})

	t.Run("remove from multi node list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		jobs := []types.Job{
			createTestJob("1", "job_1"),
			createTestJob("2", "job_2"),
			createTestJob("3", "job_3"),
		}

		// Build list: [job_1] ↔ [job_2] ↔ [job_3]
		for _, job := range jobs {
			dll.Append(job)
		}

		// Store original tail and second-to-last node for verification
		originalTail := dll.tail
		originalSecondLast := dll.tail.prev

		// Remove last node
		node := dll.RemoveLast()

		// Verify returned node
		if node == nil {
			t.Fatal("RemoveLast() should return the node, got nil")
		}

		if got := node.value.Id; got != jobs[2].Id {
			t.Errorf("returned node ID = %s, want %s", got, jobs[2].Id)
		}

		// Verify pointer cleanup on returned node
		if node.next != nil || node.prev != nil {
			t.Error("returned node should have clean pointers (next=nil, prev=nil)")
		}

		// Verify returned node is the original tail
		if node != originalTail {
			t.Error("returned node should be the original tail node")
		}

		// Verify list state
		if dll.size != 2 {
			t.Errorf("list size = %d, want 2", dll.size)
		}

		// Verify new tail is the original second-to-last node
		if dll.tail != originalSecondLast {
			t.Error("new tail should be the original second-to-last node")
		}

		// Verify new tail content
		if got := dll.tail.value.Id; got != jobs[1].Id {
			t.Errorf("new tail ID = %s, want %s", got, jobs[1].Id)
		}

		// Verify new tail has clean next pointer
		if dll.tail.next != nil {
			t.Error("new tail next should be nil")
		}

		// Verify head unchanged
		if got := dll.head.value.Id; got != jobs[0].Id {
			t.Errorf("head should remain unchanged, got ID %s, want %s", got, jobs[0].Id)
		}

		// Verify bidirectional linking integrity
		if dll.head.next != dll.tail {
			t.Error("head.next should point to tail")
		}

		if dll.tail.prev != dll.head {
			t.Error("tail.prev should point to head")
		}
	})
}

func TestRemoveConsistency(t *testing.T) {
	t.Run("single node removal consistency", func(t *testing.T) {
		// Test that RemoveFirst and RemoveLast behave identically on single-node lists
		job := createTestJob("1", "test_job")

		// Test RemoveFirst on single node
		dll1 := CreateNewDoubleLinkedList()
		dll1.Append(job)
		node1 := dll1.RemoveFirst()

		// Test RemoveLast on single node
		dll2 := CreateNewDoubleLinkedList()
		dll2.Append(job)
		node2 := dll2.RemoveLast()

		// Both should return equivalent nodes (same content, clean pointers)
		if node1 == nil || node2 == nil {
			t.Fatal("both methods should return nodes for single-node removal")
		}

		if node1.value.Id != node2.value.Id {
			t.Error("both methods should return nodes with same content")
		}

		if node1.next != nil || node1.prev != nil || node2.next != nil || node2.prev != nil {
			t.Error("both returned nodes should have clean pointers")
		}

		// Both lists should be in identical empty state
		if dll1.size != dll2.size || dll1.size != 0 {
			t.Error("both lists should have same empty size")
		}

		if dll1.head != dll2.head || dll1.head != nil {
			t.Error("both lists should have nil head")
		}

		if dll1.tail != dll2.tail || dll1.tail != nil {
			t.Error("both lists should have nil tail")
		}
	})

	t.Run("sequential operations on empty list", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()

		// Add single node and remove it
		job := createTestJob("1", "temp_job")
		dll.Append(job)
		first := dll.RemoveFirst()

		// Verify first removal worked
		if first == nil || !dll.IsEmpty() {
			t.Fatal("first removal should work and leave empty list")
		}

		// Try removing from now-empty list
		second := dll.RemoveLast()

		// Should handle empty list gracefully
		if second != nil {
			t.Error("removing from empty list should return nil")
		}

		if !dll.IsEmpty() {
			t.Error("list should remain empty")
		}
	})
}

func TestRemovePointerIntegrity(t *testing.T) {
	t.Run("removed node isolation", func(t *testing.T) {
		dll := CreateNewDoubleLinkedList()
		jobs := []types.Job{
			createTestJob("1", "job_1"),
			createTestJob("2", "job_2"),
			createTestJob("3", "job_3"),
		}

		for _, job := range jobs {
			dll.Append(job)
		}

		// Remove first node
		removedFirst := dll.RemoveFirst()

		// Remove last node
		removedLast := dll.RemoveLast()

		// Verify removed nodes are completely isolated
		if removedFirst.next != nil || removedFirst.prev != nil {
			t.Error("removed first node should have no connections")
		}

		if removedLast.next != nil || removedLast.prev != nil {
			t.Error("removed last node should have no connections")
		}

		// Verify remaining list is intact
		if dll.size != 1 {
			t.Errorf("remaining list size = %d, want 1", dll.size)
		}

		if dll.head != dll.tail {
			t.Error("single remaining node should be both head and tail")
		}

		if dll.head.next != nil || dll.head.prev != nil {
			t.Error("remaining single node should have clean pointers")
		}

		// Verify content of remaining node
		if got := dll.head.value.Id; got != jobs[1].Id {
			t.Errorf("remaining node ID = %s, want %s", got, jobs[1].Id)
		}
	})
}
