package funcgo_test

import (
	"strings"
	"testing"

	"github.com/obaraelijah/funcgo"
)

// TestForeach This function will test slices
func TestForeach(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := make([]int, 0)
	funcgo.Foreach(func(elem int) { b = append(b, elem) })(a)

	if len(a) != len(b) {
		t.Error("mismatch of slice len")
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			t.Error("elements order mismatch")
		}
	}
}

func TestMap(t *testing.T) {
	a := []int{1, 2, 3, 4, 5}

	b := funcgo.Map(func(elem int) int { return elem * elem })(a)

	if len(a) != len(b) {
		t.Error("mismatch of slice len")
	}

	for i := 0; i < len(a); i++ {
		if b[i] != a[i]*a[i] {
			t.Error("function was not executed correctly")
		}
	}

}

func TestFilter(t *testing.T) {
	s := []string{"abc", "def", "fgh"}
	a := funcgo.Filter(func(t string) bool { return strings.HasPrefix(t, "d") })(s)

	if len(a) != 1 {
		t.Error("length mismatch")
	}

	if a[0] != "def" {
		t.Error("Wrong element in slice")
	}
}

func TestAny(t *testing.T) {
	s := []string{"abc", "def", "fgh"}

	a := funcgo.Any(func(t string) bool { return false })(s)

	if a {
		t.Error("Any should not be true")
	}

	a = funcgo.Any(func(t string) bool { return t == "abc" })(s)

	if !a {
		t.Error("Any should be true")
	}
}
