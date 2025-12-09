package day01

import (
	"fmt"
	"testing"
)

func TestSecretSafe(t *testing.T) {
	fmt.Println("******** PART 1 ********")
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

func TestSecretSafePartTwo(t *testing.T) {
	fmt.Println("******** PART 2 ********")
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
	ans := testSafe.GetAnswerPartTwo()
	expected := 6
	if ans != expected {
		t.Errorf("Expected %d, got %d", expected, ans)
	}
	t.Logf("success, expected %d got %d", expected, ans)
}
