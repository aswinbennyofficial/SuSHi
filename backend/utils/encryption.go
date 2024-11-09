package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"

	// "errors"

	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/pbkdf2"
)

func EncryptString(text, password, salt string) (string,string,error) {
	// use pbkdf2 to generate key
	key:=pbkdf2.Key([]byte(password),[]byte(salt),10000,32,sha256.New)
	log.Debug().Msgf("EncryptString() : Key generated")

	

	// create a new cipher block
	block,err:=aes.NewCipher(key)
	if err != nil {
		return "", "", err
	}

	log.Debug().Msgf("Block generated")

	// Generate a random IV
	ivString:=make([]byte, aes.BlockSize)
	_,err=rand.Read(ivString)
	if err != nil {
		return "", "", err
	}

	log.Debug().Msgf("IV generated")


	// Encrypt the text
	ciphertext := make([]byte, len(text))
    stream := cipher.NewCFBEncrypter(block, ivString)
    stream.XORKeyStream(ciphertext, []byte(text))

	log.Debug().Msgf("Text encrypted")

	// Encode ciphertext and IV to base64 for easy storage/transmission
    encodedCiphertext := base64.StdEncoding.EncodeToString(ciphertext)
    encodedIV := base64.StdEncoding.EncodeToString(ivString)

    return encodedCiphertext, encodedIV, nil

}

func DecryptString(encodedCiphertext, encodedIV, password, salt string) (string,error) {
	// Decode base64 ciphertext and IV
    ciphertext, err := base64.StdEncoding.DecodeString(encodedCiphertext)
    if err != nil {
		log.Error().Msgf("Error decoding base64 ciphertext: %v", err)
        return "", err
    }
    ivString, err := base64.StdEncoding.DecodeString(encodedIV)
    if err != nil {
        return "", err
    }

	// use pbkdf2 to generate key
    key := pbkdf2.Key([]byte(password), []byte(salt), 10000, 32, sha256.New)

    // create a new cipher block
    block, err := aes.NewCipher(key)
    if err != nil {
		log.Error().Msgf("Error creating new cipher block: %v", err)
        return "", err
    }

    // Decrypt the text
    // if len(ciphertext) < aes.BlockSize {
    //     return "", errors.New("ciphertext too short")
    // }

    plaintext := make([]byte, len(ciphertext))
    stream := cipher.NewCFBDecrypter(block, ivString)
    stream.XORKeyStream(plaintext, ciphertext)

    return string(plaintext), nil

}

func GenerateUUID(len int) (string,error) {
    bytes := make([]byte, len/2)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }
    return hex.EncodeToString(bytes)[:len], nil
}