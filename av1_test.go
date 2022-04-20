package av1_test

import (
	"crypto/aes"
	"encoding/hex"
	av1 "github.com/nvx/nxp-av1-kdf"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKDF(t *testing.T) {
	testData := []struct {
		name     string
		cipher   av1.BlockCipher
		key      string
		salt     string
		expected string
	}{
		{
			name:     "aes1",
			cipher:   aes.NewCipher,
			key:      "000102030405060708090A0B0C0D0E0F",
			salt:     "011A2B3C4D5E6F99",
			expected: "A5707EF8828B5648C1E0B0A6B68C851C",
		},
	}

	for _, d := range testData {
		d := d
		t.Run(d.name, func(t *testing.T) {
			key, err := hex.DecodeString(d.key)
			require.NoError(t, err)

			salt, err := hex.DecodeString(d.salt)
			require.NoError(t, err)

			expected, err := hex.DecodeString(d.expected)
			require.NoError(t, err)

			out, err := av1.KDF(d.cipher, key, salt)
			require.NoError(t, err)
			assert.Equal(t, expected, out, hex.EncodeToString(out))
		})
	}
}
