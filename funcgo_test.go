package funcgo

import (
	"reflect"
	"testing"
)

func TestImmutableList(t *testing.T) {
	list := Cons(1, Cons(2, Cons(3, EmptyList[int]())))
	if list.IsEmpty() || list.Head() != 1 {
		t.Errorf("List creation failed: got %+v", list)
	}

	mapped := list.Map(func(x int) int { return x * 2 })
	if !reflect.DeepEqual(mapped.ToSlice(), []int{2, 4, 6}) {
		t.Errorf("List Map failed: got %v", mapped.ToSlice())
	}

	filtered := mapped.Filter(func(x int) bool { return x > 2 })
	if !reflect.DeepEqual(filtered.ToSlice(), []int{4, 6}) {
		t.Errorf("List Filter failed: got %v", filtered.ToSlice())
	}

	empty := EmptyList[int]()
	if !empty.IsEmpty() {
		t.Errorf("EmptyList failed: got %+v", empty)
	}
}

func TestCurry(t *testing.T) {
	add := Curry2(func(a, b int) int { return a + b })
	add5 := add(5)
	if add5(3) != 8 {
		t.Errorf("Curry2 failed: got %v, want 8", add5(3))
	}

	sum3 := Curry3(func(a, b, c int) int { return a + b + c })
	sum3_5 := sum3(5)
	sum3_5_3 := sum3_5(3)
	if sum3_5_3(2) != 10 {
		t.Errorf("Curry3 failed: got %v, want 10", sum3_5_3(2))
	}
}
