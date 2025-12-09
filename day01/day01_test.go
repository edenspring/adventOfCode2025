package day01

import (
	"testing"
)

func TestSecretSafe (t *testing.T) {
	directions := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	testSafe := NewSecretSafe(directions)
	ans := testSafe.GetAnswer()
	expected := 3
	if ans != expected {
		t.Errorf("Expected %d, got %d", expected, ans)
	}
	t.Logf("success, expected %d got %d", expected, ans)
}