package mnemonics

import (
	"errors"
	"math/big"
	"strings"
)

const (
	// DictionarySize specifies the size of the dictionaries that are used by
	// the mnemonics package. All dictionaries are the same length so that the
	// same []byte can be encoded into multiple languages and all results will
	// resemble eachother.
	DictionarySize = 1626
)

var (
	errEmptyInput      = errors.New("input has len 0 - not valid for conversion")
	errUnknownLanguage = errors.New("language not recognized")
	errUnknownWord     = errors.New("word not found in dictionary for given language")
)

type (
	// Language is a type-safe identifier that indicates which dictionary
	// should be used.
	Language string

	// Dictionary is a DictionarySize list of words which can be used to create
	// human-friendly entropy.
	Dictionary [DictionarySize]string

	// Phrase is the human readable version of a random []byte. Most typically,
	// a phrase is displayed to the user using the String method.
	Phrase []string
)

func bytesToInt(bs []byte) *big.Int {
	base := big.NewInt(256)
	exp := big.NewInt(1)
	result := big.NewInt(-1)
	for i := 0; i < len(bs); i++ {
		tmp := big.NewInt(int64(bs[i]))
		tmp.Add(tmp, big.NewInt(1))
		tmp.Mul(tmp, exp)
		exp.Mul(exp, base)
		result.Add(result, tmp)
	}
	return result
}

func intToBytes(bi *big.Int) (bs []byte) {
	base := big.NewInt(256)
	for bi.Cmp(base) >= 0 {
		i := new(big.Int).Mod(bi, base).Int64()
		bs = append(bs, byte(i))
		bi.Sub(bi, base)
		bi.Div(bi, base)
	}
	bs = append(bs, byte(bi.Int64()))
	return bs
}

func intToPhrase(bi *big.Int, l Language) (p Phrase, err error) {
	// Determine which dictionary to use based on the input language.
	var dict Dictionary
	switch {
	case l == English:
		dict = englishDictionary
	default:
		return nil, errUnknownLanguage
	}

	base := big.NewInt(DictionarySize)
	for bi.Cmp(base) >= 0 {
		i := new(big.Int).Mod(bi, base).Int64()
		p = append(p, dict[i])
		bi.Sub(bi, base)
		bi.Div(bi, base)
	}
	p = append(p, dict[bi.Int64()])
	return p, nil
}

func phraseToInt(p Phrase, l Language) (*big.Int, error) {
	// Determine which dictionary to use based on the input language.
	var dict Dictionary
	var prefixLen int
	switch {
	case l == English:
		dict = englishDictionary
		prefixLen = EnglishUniquePrefix
	default:
		return nil, errUnknownLanguage
	}

	base := big.NewInt(1626)
	exp := big.NewInt(1)
	result := big.NewInt(-1)
	for i := 0; i < len(p); i++ {
		// Find the index associated with the phrase.
		var tmp *big.Int
		found := false
		for j, word := range dict {
			if strings.HasPrefix(word, p[i][:prefixLen]) {
				tmp = big.NewInt(int64(j))
				found = true
				break
			}
		}
		if !found {
			return nil, errUnknownWord
		}

		// Add the index to the int.
		tmp.Add(tmp, big.NewInt(1))
		tmp.Mul(tmp, exp)
		exp.Mul(exp, base)
		result.Add(result, tmp)
	}
	return result, nil
}

// ToPhrase converts an input []byte to a human-friendly phrase. The conversion
// is reversible.
func ToPhrase(entropy []byte, l Language) (Phrase, error) {
	if len(entropy) == 0 {
		return nil, errEmptyInput
	}
	intEntropy := bytesToInt(entropy)
	return intToPhrase(intEntropy, l)
}

// FromPhrase converts an input phrase back to the original []byte.
func FromPhrase(p Phrase, l Language) ([]byte, error) {
	if len(p) == 0 {
		return nil, errEmptyInput
	}
	intEntropy, err := phraseToInt(p, l)
	if err != nil {
		return nil, err
	}
	println(intEntropy.Int64())
	return intToBytes(intEntropy), nil
}

// String combines a phrase into a single string by concatenating the
// individual words with space separation.
func (p Phrase) String() string {
	return strings.Join(p, " ")
}
