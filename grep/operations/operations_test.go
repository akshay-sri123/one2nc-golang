package operations

import (
	"fmt"
	"testing"
)

func TestSearchSubstring(t *testing.T) {
	text, err := readFromFile("test-data/test_file.txt")
	if err != nil {
		fmt.Print(err.Error())
	}
	expected := []string{"     Stealing and giving odour! Enough; no more:"}
	filterString := "giving odour"
	got := searchSubstring(text, filterString)
	if expected[0] != got[0] {
		t.Errorf("Expected %v, got %v", expected, got[0])
	}
}
