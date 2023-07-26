package encryptor

import (
	"crypto/aes"
	"encoding/hex"
)

func AES(plaintext string, key []byte) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}
