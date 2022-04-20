package av1

func MakeSalt(keyNo uint8, uid []byte) []byte {
	salt := make([]byte, len(uid)+1)
	salt[0] = keyNo
	copy(salt[1:], uid)
	return salt
}

func expandSalt(salt []byte, l int) []byte {
	if len(salt) >= l {
		return salt[0:l]
	}

	newSalt := make([]byte, l)
	for i := 0; i < l; i++ {
		newSalt[i] = salt[i%len(salt)]
	}

	return newSalt
}
