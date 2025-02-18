package orderset

import (
	"testing"
)

func TestOrderSet_Add(t *testing.T) {
	var set *OrderSet[int]
	set = set.Add(99)
	set = set.Add(88)
	set = set.Add(77)

	if !set.Contains(99) && !set.Contains(88) && !set.Contains(77) {
		t.Error("Contain error")
	}
}

func TestOrderSet_Contains(t *testing.T) {
	var set *OrderSet[int]
	set = set.Add(99)
	set = set.Add(88)
	set = set.Add(77)

	b1 := set.Contains(99)
	if !b1 {
		t.Error("Contain error")
	}

	b2 := set.Contains(88)
	if !b2 {
		t.Error("Contain error")
	}

	set = nil
	b3 := set.Contains(88)
	if b3 {
		t.Error("Contain error")
	}
}

func TestOrderSet_Index(t *testing.T) {
	var set = NewOrderSet[int]()
	set.Add(99)
	set.Add(88)
	set.Add(77)
	set.Add(66)

	set.Del(88)
	if set.Index(99) != 0 {
		t.Error("Index error")
	}

	set.Add(77)
	if set.Index(77) != 1 {
		t.Error("Index error")
	}

	set.Add(66)
	if set.Index(66) != 2 {
		t.Error("Index error")
	}

	if set.Index(88) != -1 {
		t.Error("Index error")
	}

	set = nil
	if set.Index(99) != -1 {
		t.Error("Index error")
	}
}

func TestOrderSet_Count(t *testing.T) {
	var set = NewOrderSet[int]()
	set.Add(99)
	set.Add(88)
	set.Add(77)
	set.Add(66)

	set.Del(88)
	if set.Count() != 3 {
		t.Error("Count number error")
	}

	set.Add(77)
	if set.Count() != 3 {
		t.Error("Contain error")
	}

	set.Add(88)
	if set.Count() != 4 {
		t.Error("Contain error")
	}

	set = nil
	if set.Count() != 0 {
		t.Error("Contain error")
	}
}

func TestOrderSet_Del(t *testing.T) {
	var set = NewOrderSet[int]()
	set.Add(99)
	set.Add(88)
	set.Add(77)

	set.Del(88)
	if set.Contains(88) {
		t.Error("Contain error")
	}
	set.Del(77)
	if set.Contains(77) {
		t.Error("Contain error")
	}
}

func TestOrderSet_ToSlice(t *testing.T) {
	var set = NewOrderSet[int]()
	set.Add(99)
	set.Add(88)
	set2 := set
	set.Add(77)

	s := set2.ToSlice()

	count := len(s)
	if count != 3 {
		t.Error("Size error")
	}

	if s[0] != 99 {
		t.Error("Position error")
	}
	if s[2] != 77 {
		t.Error("Position error")
	}

	set = nil
	s = set.ToSlice()
	if len(s) != 0 {
		t.Error("Size error")
	}
}
