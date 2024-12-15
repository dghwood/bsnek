package permutations

import "testing"

func compareSlices(r, e [][]int, t *testing.T) {
	if len(r) != len(e) {
		t.Fatalf("Perms incorrect length found: %d, expected: %d", len(r), len(e))
	}
	for i := 0; i < len(r); i++ {
		r1 := r[i]
		e1 := e[i]
		if len(r1) != len(e1) {
			t.Fatalf("Entry wrong length %d", len(r1))
		}
		for j := 0; j < len(r1); j++ {
			if r1[j] != e1[j] {
				t.Fatal("Wrong perms", i, j, r1, e1)
			}
		}
	}
}
func TestPermutations(t *testing.T) {
	r := Permutations([]int{2, 2, 2})
	e := [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}
	compareSlices(r, e, t)
}

func TestPermutations2(t *testing.T) {
	r := Permutations([]int{3, 2, 0})
	e := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{1, 0, 0},
		{1, 1, 0},
		{2, 0, 0},
		{2, 1, 0},
	}
	compareSlices(r, e, t)
}

func TestPermutations3(t *testing.T) {
	r := Permutations([]int{1, 2})
	e := [][]int{
		{0, 0},
		{0, 1},
	}
	compareSlices(r, e, t)
}
