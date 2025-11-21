package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"dm_loanservice/drivers/goconf"
	"encoding/base64"
	"io"

	"github.com/pkg/errors"
	"github.com/speps/go-hashids/v2"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPassword(password string) (string, error) {
	bytePassword := []byte(password)
	hash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", errors.WithMessage(err, "error hash")
	}

	return string(hash), nil
}

func DecryptPassword(hashedPassword, password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, errors.WithMessage(err, "invalid password")
	}

	return true, nil
}

func createHash() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = goconf.Config().GetString("salt")
	hd.MinLength = 16
	h, _ := hashids.NewWithData(hd)
	return h
}

func Encrypt(data []int) (string, error) {
	hash := createHash()
	dataEncrypted, err := hash.Encode(data)
	if err != nil {
		return "", err
	}
	return dataEncrypted, nil
}

func Decrypt(data string) (int, error) {
	hash := createHash()
	dataDecrypted, err := hash.DecodeWithError(data)
	if err != nil {
		return 0, err
	}
	return dataDecrypted[0], nil
}

// EncryptSecret encrypts plaintext with AES-256-GCM.
func EncryptSecret(key []byte, plaintext string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("AES-256 key must be 32 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)

	// nonce||ciphertext then base64
	out := append(nonce, ciphertext...)
	return base64.RawStdEncoding.EncodeToString(out), nil
}

// DecryptSecret decrypts base64(nonce||ciphertext) with AES-256-GCM.
func DecryptSecret(key []byte, encoded string) (string, error) {
	if len(key) != 32 {
		return "", errors.New("AES-256 key must be 32 bytes")
	}

	data, err := base64.RawStdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	ns := gcm.NonceSize()
	if len(data) < ns {
		return "", errors.New("ciphertext too short")
	}

	nonce := data[:ns]
	ciphertext := data[ns:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
