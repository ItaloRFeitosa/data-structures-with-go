package linkedlist

import (
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("should append element", func(t *testing.T) {
		expectedSize := uint(3)
		first := "first_element"
		second := "second_element"
		third := "third_element"

		list := New[string]()

		list.Append(first)
		list.Append(second)
		list.Append(third)

		if list.size != expectedSize {
			t.Errorf("expected list size %d got %d", expectedSize, list.size)
		}

		firstBlock := list.head
		if firstBlock.data != first {
			t.Errorf("expected first block data to be %s got %s", first, firstBlock.data)
		}

		secondBlock := firstBlock.next
		if secondBlock.data != second {
			t.Errorf("expected second block data to be %s got %s", second, secondBlock.data)
		}

		thirdBlock := secondBlock.next
		if thirdBlock.data != third {
			t.Errorf("expected third block data to be %s got %s", third, thirdBlock.data)
		}

	})

	t.Run("should append element", func(t *testing.T) {
		expectedSize := uint(4)
		first := "first_element"
		second := "second_element"
		third := "third_element"
		forth := "forth_element"

		list := New[string]()

		list.Append(first)
		list.Append(second)
		list.Append(third)
		list.Append(forth)

		list.Remove(0)
		expectedSize--

		if list.size != expectedSize {
			t.Errorf("expected list size %d got %d", expectedSize, list.size)
		}

		if list.head.data != second {
			t.Errorf("expected head block data to be %s got %s", second, list.head.data)
		}

		list.Remove(1)
		expectedSize--

		if list.size != expectedSize {
			t.Errorf("expected list size %d got %d", expectedSize, list.size)
		}

		newSecondBlock := list.head.next
		if newSecondBlock.data != forth {
			t.Errorf("expected new second block data to be %s got %s", forth, newSecondBlock.data)
		}
	})
}
