package funcgo_test

import (
	"testing"

	"github.com/obaraelijah/funcgo"
)

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
