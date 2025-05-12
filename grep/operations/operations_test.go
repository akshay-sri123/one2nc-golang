package operations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSearchSubstringTable(t *testing.T) {
	text, err := readFromFile("test-data/test_file.txt")
	if err != nil {
		fmt.Print(err.Error())
	}
	var runTests = []struct {
		name         string
		filterString string
		expected     []string
	}{
		{
			name:         "Return single substring from file",
			filterString: "giving odour",
			expected: []string{
				"     Stealing and giving odour! Enough; no more:",
			},
		},
		{
			name:         "Return multiple substring from file",
			filterString: "what",
			expected: []string{
				"     Of what validity and pitch soe'er,",
				"     How now! what news from her?",
			},
		},
		{
			name:         "Return multiple substring from file 2",
			filterString: "like",
			expected: []string{
				"     O, it came o'er my ear like the sweet sound,",
				"     And my desires, like fell and cruel hounds,",
				"     But, like a cloistress, she will veiled walk",
			},
		},
	}

	for _, tt := range runTests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchSubstring(text, tt.filterString)
			//  Check the length of the result set
			if len(tt.expected) != len(got) {
				t.Errorf("Expected %d results, got %d", len(tt.expected), len(got))
			}
			// Compare the values of the result set
			if !reflect.DeepEqual(tt.expected, got) {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}
