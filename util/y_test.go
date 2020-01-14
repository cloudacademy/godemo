package util

import "testing"

func TestGetY(t *testing.T) {
	ans := GetY()
	if ans != 200 {
		t.Errorf("GetX() = %d; want 200", ans)
	}
}
