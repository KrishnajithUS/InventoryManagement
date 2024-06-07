package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

var secretKey string = "0ED8463FF7C5CE2E93309811FA892D63"

func encodeToBase64(ciphertext []byte) string {
	return hex.EncodeToString(ciphertext)
}

func decodeToBytes(ciphertext string) []byte {
	res, err := hex.DecodeString(ciphertext)
	fmt.Println("err",err)
	return res
}

func Encrypt(plaintext string) (string, error) {
	aes, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = rand.Read(nonce)
	if err != nil {
		return "", err
	}
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	encodeToBase64 := encodeToBase64(ciphertext)
	fmt.Println("Cipher data,",ciphertext)
	return encodeToBase64, nil

}

func Decrypt(cipertext string) (string, error) {
	aes, err := aes.NewCipher([]byte(secretKey))
	cipertextInBytes := decodeToBytes(cipertext)
	fmt.Println(cipertextInBytes)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(aes)
	if err != nil {
		return "", err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipertextInBytes := cipertextInBytes[:nonceSize], cipertextInBytes[nonceSize:]

	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(cipertextInBytes), nil)
	if err != nil {
		return "", err
	}
	encodeToBase64 := encodeToBase64(plaintext)
	return encodeToBase64, nil
}

func VerifyPasswordAES(password string, hashedPassword string) error {
	decryptedPassword, err := Decrypt(hashedPassword)
	fmt.Println("decrypted", decodeToBytes)
	fmt.Println("err",err)
	if err != nil {
		return err
	}
	fmt.Println("give password", encodeToBase64([]byte(password)))
	hexPassword := encodeToBase64([]byte(password))
	fmt.Println("decrypted",decryptedPassword)
	if hexPassword != decryptedPassword {
		return fmt.Errorf("password mismatch")
	}
	return nil
}
