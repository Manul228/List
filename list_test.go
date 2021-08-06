package list

import (
	"testing"
)

func TestAddAtHead(t *testing.T) {
	testList := Constructor()
	testList.AddAtHead(1)
	testList.AddAtHead(2)
	testList.AddAtHead(3)

	current := testList.Head

	if current.Prev != nil {
		t.Error("Head Prev is not nil")
	}

	for i := 1; i < testList.Length+1; i++ {
		if current.Val != testList.Length-i+1 {
			t.Errorf("wrong element %v != %v", current.Val, testList.Length-i)
		}

		if current.Next != nil {
			current = current.Next
		}
	}
	if testList.Tail.Val != 1 {
		t.Error("wrong Tail")
	}
}

func TestAddAtTail(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)

	current := testList.Head

	if current.Prev != nil {
		t.Error("Head Prev is not nil")
	}

	for i := 1; i < testList.Length+1; i++ {
		if current.Val != i {
			t.Errorf("wrong element %v != %v", current.Val, i)
		}

		if current.Next != nil {
			current = current.Next
		}
	}
	if testList.Tail.Val != 3 {
		t.Error("wrong Tail")
	}
}

func TestGet(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)

	for i := 0; i < testList.Length; i++ {
		val := testList.Get(i)
		if val != i+1 {
			t.Errorf("wrong element %v at index %v", val, i)
		}
	}

	if testList.Get(123) != -1 {
		t.Error("index verification failed")
	}

	if testList.Get(0) != 1 {
		t.Error("invalid Get Head ")
	}
}

func TestInsertAtIndex(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)

	testList.AddAtIndex(3, 4)

	val := testList.Get(3)
	if val != 4 {
		t.Error("add at Tail by index has failed: ", val, 4)
	}

	testList.AddAtIndex(3, 33)

	val = testList.Get(3)
	if val != 33 {
		t.Error("add at Tail by index has failed: ", val, 4)
	}
}

func TestLoopTrue(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)

	testList.Head.Next.Next.Next = testList.Head.Next

	if HasCycle(testList.Head) != true {
		t.Error("wrong")
	}
}

func TestLoopFalse(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)

	if HasCycle(testList.Head) != false {
		t.Error("wrong")
	}
}

func TestLoopEmptyList(t *testing.T) {
	testList := Constructor()

	if HasCycle(testList.Head) != false {
		t.Error("wrong")
	}
}

func TestLoopOneTwo(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)

	if HasCycle(testList.Head) != false {
		t.Error("wrong")
	}
}

func TestLoopDetectSimple(t *testing.T) {
	testList := Constructor()
	testList.AddAtTail(1)
	testList.AddAtTail(2)
	testList.AddAtTail(3)
	testList.AddAtTail(4)

	pos := testList.Head.Next

	testList.Head.Next.Next.Next.Next = pos

	if DetectCycle(testList.Head) != pos {
		t.Error("wrong")
	}
}

func TestIntersectSimple(t *testing.T) {
	tl1 := Constructor()
	tl2 := Constructor()

	tl1.AddAtTail(1)
	tl1.AddAtTail(2)
	tl1.AddAtTail(3)
	tl1.AddAtTail(4)

	tl2.AddAtTail(5)
	tl2.AddAtTail(6)
	tl2.AddAtTail(7)
	tl2.AddAtTail(8)

	tl := Constructor()
	tl.AddAtTail(9)
	tl.AddAtTail(10)

	tl1.Tail.Next = tl.Head
	tl2.Tail.Next = tl.Head
	tl1.Tail = tl.Tail
	tl2.Tail = tl.Tail

	PrintMe(tl1.Head)
	PrintMe(tl2.Head)
}
