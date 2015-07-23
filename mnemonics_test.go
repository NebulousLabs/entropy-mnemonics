package mnemonics

import (
	"math/big"
	"testing"
)

// TestBytesToInt probes the bytesToInt function.
func TestBytesToInt(t *testing.T) {
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
