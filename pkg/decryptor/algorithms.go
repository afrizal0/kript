package decryptor

import (
	"crypto/aes"
	"encoding/hex"
)

func AES(plaintext string, key []byte) string {
	k := []byte(key)
	c, err := aes.NewCipher(k)
	if err != nil {
		panic(err)
	}
	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}
