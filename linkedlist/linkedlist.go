package linkedlist

type LinkedList[T any] struct {
	head *Block[T]
	size uint
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (ll *LinkedList[T]) Append(data T) {
	var block Block[T]
	block.data = data

	ll.size++

	if ll.head == nil {
		ll.head = &block
		return
	}

	cursor := ll.head

	for cursor.next != nil {
		cursor = cursor.next
	}

	cursor.next = &block

}

func (ll *LinkedList[T]) Remove(index uint) {
	var cursorPosition uint

	ll.size--

	if cursorPosition == index {
		currentHead := ll.head
		ll.head = currentHead.next
		currentHead.next = nil
		return
	}

	cursor := ll.head

	for cursor.next != nil {

		cursorPosition++
		if cursorPosition == index {
			toRemove := cursor.next
			cursor.next = toRemove.next
			toRemove.next = nil
			return
		}

		cursor = cursor.next
	}
}

type Block[T any] struct {
	data T
	next *Block[T]
}
