package mnemonics

import (
	"bytes"
	"math/big"
	"testing"
)

// TestUnitBytesToInt probes the bytesToInt function.
func TestUnitBytesToInt(t *testing.T) {
	// Try for value {0}.
	expected := big.NewInt(0)
	result := bytesToInt([]byte{0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '0' value")
	}

	// Try for value {1}.
	expected = big.NewInt(1)
	result = bytesToInt([]byte{1})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '1' value")
	}

	// Try for value {255}.
	expected = big.NewInt(255)
	result = bytesToInt([]byte{255})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '255' value")
	}

	// Try for value {0, 0}.
	expected = big.NewInt(256)
	result = bytesToInt([]byte{0, 0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '256' value")
	}

	// Try for value {1, 0}.
	expected = big.NewInt(257)
	result = bytesToInt([]byte{1, 0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '257' value")
	}

	// Try for value {0, 1}.
	expected = big.NewInt(512)
	result = bytesToInt([]byte{0, 1})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '512' value")
	}

	// Try for value {1, 1}.
	expected = big.NewInt(513)
	result = bytesToInt([]byte{1, 1})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '513' value")
	}

	// Try for value {2, 1}.
	expected = big.NewInt(514)
	result = bytesToInt([]byte{2, 1})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '514' value")
	}

	// Try for value {2, 2}.
	expected = big.NewInt(770)
	result = bytesToInt([]byte{2, 2})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '770' value")
	}

	// Try for value {0, 255}.
	expected = big.NewInt(65536)
	result = bytesToInt([]byte{0, 255})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '65536' value")
	}

	// Try for value {0, 0, 0}.
	expected = big.NewInt(65792)
	result = bytesToInt([]byte{0, 0, 0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '65792' value")
	}

	// Try for value {1, 0, 0}.
	expected = big.NewInt(65793)
	result = bytesToInt([]byte{1, 0, 0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '65793' value")
	}

	// Try for value {0, 1, 0}.
	expected = big.NewInt(66048)
	result = bytesToInt([]byte{0, 1, 0})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '66048' value")
	}

	// Try for value {0, 0, 1}.
	expected = big.NewInt(131328)
	result = bytesToInt([]byte{0, 0, 1})
	if expected.Cmp(result) != 0 {
		t.Error("failure for '131328' value")
	}
}

// TestIntegrationConversions checks ToPhrase and FromPhrase for consistency
// and sanity.
func TestIntegrationConversions(t *testing.T) {
	// Try for value {0}.
	initial := []byte{0}
	phrase, err := ToPhrase(initial, English)
	if err != nil {
		t.Error(err)
	}
	if len(phrase) != 1 {
		t.Fatal("unexpected phrase length")
	}
	if phrase[0] != englishDictionary[0] {
		t.Error("unexpected ToPhrase result")
	}
	final, err := FromPhrase(phrase, English)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(initial, final) != 0 {
		t.Error("failure for value {0}")
	}

	// Try for value {1}.
	initial = []byte{1}
	phrase, err = ToPhrase(initial, English)
	if err != nil {
		t.Error(err)
	}
	if len(phrase) != 1 {
		t.Fatal("unexpected phrase length")
	}
	if phrase[0] != englishDictionary[1] {
		t.Error("unexpected ToPhrase result")
	}
	final, err = FromPhrase(phrase, English)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(initial, final) != 0 {
		t.Error("failure for value {1}")
	}

	// Try for value {255}.
	initial = []byte{255}
	phrase, err = ToPhrase(initial, English)
	if err != nil {
		t.Error(err)
	}
	if len(phrase) != 1 {
		t.Fatal("unexpected phrase length")
	}
	if phrase[0] != englishDictionary[255] {
		t.Error("unexpected ToPhrase result")
	}
	final, err = FromPhrase(phrase, English)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(initial, final) != 0 {
		t.Error("failure for value {255}")
	}

	// Try for value {0, 0}.
	initial = []byte{0, 0}
	phrase, err = ToPhrase(initial, English)
	if err != nil {
		t.Error(err)
	}
	if len(phrase) != 1 {
		t.Fatal("unexpected phrase length")
	}
	if phrase[0] != englishDictionary[256] {
		t.Error("unexpected ToPhrase result")
	}
	final, err = FromPhrase(phrase, English)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(initial, final) != 0 {
		t.Error("failure for value {0, 0}")
	}

	// Try for value {abbey, abbey}.
	initial = []byte{90, 5}
	phrase, err = ToPhrase(initial, English)
	if err != nil {
		t.Error(err)
	}
	if len(phrase) != 2 {
		t.Fatal("unexpected phrase length")
	}
	if phrase[0] != englishDictionary[0] {
		t.Error("unexpected ToPhrase result")
	}
	if phrase[1] != englishDictionary[0] {
		t.Error("unexpected ToPhrase result")
	}
	final, err = FromPhrase(phrase, English)
	if err != nil {
		t.Error(err)
	}
	if bytes.Compare(initial, final) != 0 {
		t.Error("failure for value {abbey, abbey}")
	}

	// Check that all values going from []byte to phrase and back result in the
	// original value, as deep as reasonable.
	for i := 0; i < 256; i++ {
		initial := []byte{byte(i)}
		phrase, err := ToPhrase(initial, English)
		if err != nil {
			t.Fatal(err)
		}
		final, err := FromPhrase(phrase, English)
		if err != nil {
			t.Fatal(err)
		}
		if bytes.Compare(initial, final) != 0 {
			t.Error("comparison failed during circular byte check")
		}
	}
	for i := 0; i < 256; i++ {
		for j := 0; j < 256; j++ {
			initial := []byte{byte(i), byte(j)}
			phrase, err := ToPhrase(initial, English)
			if err != nil {
				t.Fatal(err)
			}
			final, err := FromPhrase(phrase, English)
			if err != nil {
				t.Fatal(err)
			}
			if bytes.Compare(initial, final) != 0 {
				t.Error("comparison failed during circular byte check")
			}
		}
	}
	// It takes too long to try all numbers 3 deep, so only a handful are
	// picked. All edge numbers are checked.
	for i := 0; i < 256; i++ {
		for _, j := range []byte{0, 1, 2, 3, 16, 25, 82, 200, 252, 253, 254, 255} {
			for _, k := range []byte{0, 1, 2, 3, 9, 29, 62, 104, 105, 217, 252, 253, 254, 255} {
				initial := []byte{byte(i), j, k}
				phrase, err := ToPhrase(initial, English)
				if err != nil {
					t.Fatal(err)
				}
				final, err := FromPhrase(phrase, English)
				if err != nil {
					t.Fatal(err)
				}
				if bytes.Compare(initial, final) != 0 {
					t.Error("comparison failed during circular byte check")
				}
			}
		}
	}

	// Check that all values going from phrase to []byte and back result in the
	// original value, as deep as reasonable.
	for i := 0; i < DictionarySize; i++ {
		initial := Phrase{englishDictionary[i]}
		entropy, err := FromPhrase(initial, English)
		if err != nil {
			t.Fatal(err)
		}
		final, err := ToPhrase(entropy, English)
		if err != nil {
			t.Fatal(err)
		}
		if len(initial) != len(final) {
			t.Fatal("conversion error")
		}
		for i := range initial {
			if initial[i] != final[i] {
				t.Error("conversion error")
			}
		}
	}
	// It takes too long to try all numbers 2 deep for phrases, so the test it
	// not comprehensive. All edge numbers are checked.
	for i := 0; i < DictionarySize; i++ {
		for _, j := range []int{0, 1, 2, 3, 4, 5, 6, 25, 50, 75, 122, 266, 305, 1620, 1621, 1622, 1623, 1623, 1625} {
			initial := Phrase{englishDictionary[i], englishDictionary[j]}
			entropy, err := FromPhrase(initial, English)
			if err != nil {
				t.Fatal(err)
			}
			final, err := ToPhrase(entropy, English)
			if err != nil {
				t.Fatal(err)
			}
			if len(initial) != len(final) {
				t.Fatal("conversion error")
			}
			for i := range initial {
				if initial[i] != final[i] {
					t.Error("conversion error")
				}
			}
		}
	}
	// It takes too long to try all numbers 2 deep for phrases, so the test it
	// not comprehensive. All edge numbers are checked.
	for _, i := range []int{0, 1, 2, 3, 4, 5, 6, 25, 50, 75, 122, 266, 305, 1620, 1621, 1622, 1623, 1623, 1625} {
		for _, j := range []int{0, 1, 2, 3, 4, 5, 6, 25, 50, 75, 122, 266, 305, 1620, 1621, 1622, 1623, 1623, 1625} {
			for _, k := range []int{0, 1, 2, 3, 4, 5, 6, 25, 50, 75, 122, 266, 305, 1620, 1621, 1622, 1623, 1623, 1625} {
				initial := Phrase{englishDictionary[i], englishDictionary[j], englishDictionary[k]}
				entropy, err := FromPhrase(initial, English)
				if err != nil {
					t.Fatal(err)
				}
				final, err := ToPhrase(entropy, English)
				if err != nil {
					t.Fatal(err)
				}
				if len(initial) != len(final) {
					t.Fatal("conversion error")
				}
				for i := range initial {
					if initial[i] != final[i] {
						t.Error("conversion error")
					}
				}
			}
		}
	}
}
