package math

import "testing"

func TestAdd(t *testing.T) {
	ans := Add(2, 5)
	if ans != 7 {
		t.Errorf("Add(2, 5) = %d; want 7", ans)
	}
}
