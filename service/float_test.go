package service

import (
	"testing"
)

// TestMultiHundred : 測試 浮點數乘100
func TestMultiHundred(t *testing.T) {
	have := MultiHundred(1129.6)

	want := float64(1129600)
	if have != want {
		t.Fatalf("MultiHundred(1129.6)=%f want=%f", have, want)
	}
}
