package utility

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type params struct {
	time    uint32
	memory  uint32
	threads uint8
	keylen  uint32
}

func decodeHash(passwordHash string) (hash []byte, salt []byte, p *params, err error) {
	// "" argon2id v m,t,p salt hash
	p = &params{}
	values := strings.Split(passwordHash, "$")
	salt, err = base64.RawStdEncoding.Strict().DecodeString(values[4])
	if err != nil {
		return nil, nil, nil, err
	}
	hash, err = base64.RawStdEncoding.Strict().DecodeString(values[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.keylen = uint32(len(hash))
	_, err = fmt.Sscanf(values[3], "m=%d,t=%d,p=%d", &p.memory, &p.time, &p.threads)
	if err != nil {
		return nil, nil, nil, err
	}
	return hash, salt, p, nil
}

func generateSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	return salt, err
}

func HashPassword(password string) (string, error) {
	var passwordHash string
	passwordBytes := []byte(password)
	salt, err := generateSalt(16)
	if err != nil {
		return passwordHash, err
	}
	p := &params{
		time:    2,
		memory:  19 * 1024,
		threads: 1,
		keylen:  32,
	}
	hash := argon2.IDKey(
		passwordBytes,
		salt,
		p.time,
		p.memory,
		p.threads,
		p.keylen,
	)
	encodedSalt := base64.RawStdEncoding.EncodeToString(salt)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)
	passwordHash = fmt.Sprintf(
		"$argon2id$%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version,
		p.memory,
		p.time,
		p.threads,
		encodedSalt,
		encodedHash,
	)
	return passwordHash, nil
}

func VerifyPassword(passwordHash string, password string) (bool, error) {
	hash, salt, params, err := decodeHash(passwordHash)
	if err != nil {
		return false, err
	}
	passwordBytes := []byte(password)
	newHash := argon2.IDKey(passwordBytes, salt, params.time, params.memory, params.threads, params.keylen)
	result := subtle.ConstantTimeCompare(hash, newHash)
	if result == 0 {
		return false, nil
	}
	return true, nil
}
