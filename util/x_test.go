package util

import "testing"

func TestGetX(t *testing.T) {
	ans := GetX()
	if ans != 100 {
		t.Errorf("GetX() = %d; want 100", ans)
	}
}
