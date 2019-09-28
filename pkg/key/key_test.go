package key

import (
	"testing"
)

func TestGenerate(t *testing.T) {
	testCases := []struct {
		desc   string
		length int
	}{
		{
			desc:   "should be 8 chars in length",
			length: 8,
		},
		{
			desc:   "should be 19 chars in length",
			length: 19,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			key := Generate(tC.length)

			if len(key) != tC.length {
				t.Errorf("expected the string to be %v, got %v instead", tC.length, len(key))
			}
		})
	}
}
