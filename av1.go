package av1

import (
	"crypto/cipher"
	"errors"
)

type BlockCipher func(key []byte) (cipher.Block, error)

func KDF(blockCipher BlockCipher, masterKey []byte, salt []byte) ([]byte, error) {
	c, err := blockCipher(masterKey)
	if err != nil {
		return nil, err
	}

	blockSize := c.BlockSize()

	if len(masterKey) != blockSize || len(salt) == 0 || len(salt) > blockSize {
		return nil, errors.New("invalid masterKey or salt lengths")
	}

	salt = expandSalt(salt, blockSize)
	message := make([]byte, len(salt))
	for i := range message {
		message[i] = salt[i] ^ masterKey[i]
	}

	c.Encrypt(message, message)

	return message, nil
}
