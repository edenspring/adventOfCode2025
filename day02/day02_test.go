package day02

import (
	"slices"
	"testing"
)

func TestDay02PartOne(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	expectedTotal := 1227775554
	expectedBadIDs := []int{11,22,99,1010,1188511885,222222,446446,38593859}

	testChecker := NewIDChecker(input)
	actualBadIDs := testChecker.GetBadIDs()
	actualTotal := testChecker.GetTotal()

	if len(expectedBadIDs) != len(actualBadIDs) {
		t.Errorf("ERR, Mismatched lengths in expected and actual: %d != %d", len(expectedBadIDs), len(actualBadIDs))
	}

	for _, num := range actualBadIDs {
		if !slices.Contains(expectedBadIDs, num) {
			t.Errorf("ERR, Found unexpected id: %d", num)
		}
	}

	if expectedTotal != actualTotal {
		t.Errorf("ERR, Expected total does not match actual total: %d != %d", expectedTotal, actualTotal)
	}
}

func TestDay02PartTwo(t *testing.T) {
	input := `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124`
	expectedTotal := 4174379265
	expectedBadIDs := []int{11,22,99,111,999,1010,1188511885,222222,446446,38593859,565656,824824824,2121212121}

	testChecker := NewIDChecker(input)
	actualBadIDs := testChecker.GetBadIDsPartTwo()
	actualTotal := testChecker.GetTotal()

	if len(expectedBadIDs) != len(actualBadIDs) {
		t.Errorf("ERR, Mismatched lengths in expected and actual: %d != %d", len(expectedBadIDs), len(actualBadIDs))
	}

	for _, num := range actualBadIDs {
		if !slices.Contains(expectedBadIDs, num) {
			t.Errorf("ERR, Found unexpected id: %d", num)
		}
	}

	if expectedTotal != actualTotal {
		t.Errorf("ERR, Expected total does not match actual total: %d != %d", expectedTotal, actualTotal)
	}
}

