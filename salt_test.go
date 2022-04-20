package av1

import (
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestExpandSalt(t *testing.T) {
	testData := []struct {
		name     string
		len      int
		in       string
		expected string
	}{
		{
			name:     "exact",
			len:      16,
			in:       "011A2B3C4D5E6F99011A2B3C4D5E6F99",
			expected: "011A2B3C4D5E6F99011A2B3C4D5E6F99",
		},
		{
			name:     "double",
			len:      16,
			in:       "011A2B3C4D5E6F99",
			expected: "011A2B3C4D5E6F99011A2B3C4D5E6F99",
		},
		{
			name:     "larger",
			len:      16,
			in:       "011A2B3C4D5E6F99011A2B3C4D5E6F99BEEF",
			expected: "011A2B3C4D5E6F99011A2B3C4D5E6F99",
		},
	}

	for _, d := range testData {
		d := d
		t.Run(d.name, func(t *testing.T) {
			in, err := hex.DecodeString(d.in)
			require.NoError(t, err)

			expected, err := hex.DecodeString(d.expected)
			require.NoError(t, err)

			out := expandSalt(in, d.len)
			assert.Equal(t, expected, out)
		})
	}
}

func TestMakeSalt(t *testing.T) {
	testData := []struct {
		name     string
		keyNo    uint8
		uid      string
		expected string
	}{
		{
			name:     "keyNo0",
			keyNo:    0,
			uid:      "1A2B3C4D5E6F99",
			expected: "001A2B3C4D5E6F99",
		},
		{
			name:     "keyNo1",
			keyNo:    1,
			uid:      "1A2B3C4D5E6F99",
			expected: "011A2B3C4D5E6F99",
		},
	}

	for _, d := range testData {
		d := d
		t.Run(d.name, func(t *testing.T) {
			uid, err := hex.DecodeString(d.uid)
			require.NoError(t, err)

			expected, err := hex.DecodeString(d.expected)
			require.NoError(t, err)

			out := MakeSalt(d.keyNo, uid)
			assert.Equal(t, expected, out)
		})
	}
}
