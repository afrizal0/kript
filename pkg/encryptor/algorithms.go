package encryptor

import (
	"bufio"
	"crypto/aes"
	"encoding/hex"
	"os"
	"strings"
)

/*
TASK:
	1. Support data files encryption ✅[Still have some bug :(]
*/

func ReadFiles(path string) string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var data string
	s := bufio.NewScanner(file)
	for s.Scan() {
		read_line := s.Text()
		read_line = strings.TrimSpace(read_line)
		data += read_line
	}

	return data
}

func AES(plaintext string, key []byte) string {
	c, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}
