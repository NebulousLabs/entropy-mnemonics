package mnemonics

import (
	"testing"
)

// TestEnglishDictionary checks that the english dictionary is well formed.
func TestEnglishDictionary(t *testing.T) {
	// Check for sane constants.
	if English != "english" {
		t.Error("unexpected identifier for english dictionary")
	}
	if EnglishUniquePrefix != 3 {
		t.Error("unexpected prefix len for english dictionary")
	}

	// Check that the dictionary has well formed elements, and no repeats.
	engMap := make(map[string]struct{})
	for i, word := range englishDictionary {
		if len(word) < EnglishUniquePrefix {
			t.Fatal("found a short word at index", i)
		}

		_, exists := engMap[word[:3]]
		if exists {
			t.Error("found a prefix conflict at index", i)
		}
		engMap[word[:3]] = struct{}{}
	}
}
