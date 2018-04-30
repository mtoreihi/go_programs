package main

import (
"crypto/aes"
"crypto/cipher"
"errors"
"fmt"
"log"
"encoding/base64"
)

func main() {
	text := []byte("My name is Mehran")
	key := []byte("the-key-has-to-be-32-bytes-long!")

	ciphertext, err := encrypt(text, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	b64 := base64.StdEncoding.EncodeToString(ciphertext)
	fmt.Printf("%s => %s\n", text, b64)

	b64_decoded,_ := base64.StdEncoding.DecodeString(b64)
	plaintext, err := decrypt(b64_decoded, key)
	if err != nil {
		// TODO: Properly handle error
		log.Fatal(err)
	}
	fmt.Printf("%s => %s\n", b64, plaintext)
}

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)

	if err != nil {
		return nil, err
	}

	//nonce := make([]byte, gcm.NonceSize())
	nonce := []byte("1234567890--")
	/*if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}*/

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
