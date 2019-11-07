// Hash strings or files
package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"io"
	"log"
	"os"
)

// GetMD5FromFile returns the MD5 hash of a file.
func GetMD5FromFile(file os.File) string {
	hasher := md5.New()
	content, _ := os.Open(file.Name())
	defer content.Close()

	if _, err := io.Copy(hasher, content); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))

}

// GetMD5 returns the MD5 hash of a string.
func GetMD5(text string) string {
	data := []byte(text)
	b := md5.Sum(data)
	return hex.EncodeToString(b[:])
}

// GetSHA1FromFile returns the SHA1 hash of a file.
func GetSHA1FromFile(file os.File) string {
	hasher := sha1.New()
	content, _ := os.Open(file.Name())
	defer content.Close()

	if _, err := io.Copy(hasher, content); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))

}

// GetSHA1 returns the SHA1 hash of a string.
func GetSHA1(text string) string {
	data := []byte(text)
	b := sha1.Sum(data)
	return hex.EncodeToString(b[:])
}

// GetSHA256FromFile returns the SHA256 hash of a file.
func GetSHA256FromFile(file os.File) string {
	hasher := sha256.New()
	content, _ := os.Open(file.Name())
	defer content.Close()

	if _, err := io.Copy(hasher, content); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))

}

// GetSHA256 returns the SHA256 hash of a string.
func GetSHA256(text string) string {
	data := []byte(text)
	b := sha256.Sum256(data)
	return hex.EncodeToString(b[:])
}

// GetSHA512FromFile returns the SHA512 hash of a file.
func GetSHA512FromFile(file os.File) string {
	hasher := sha512.New()
	content, _ := os.Open(file.Name())
	defer content.Close()

	if _, err := io.Copy(hasher, content); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hasher.Sum(nil))

}

// GetSHA512 returns the SHA512 hash of a string.
func GetSHA512(text string) string {
	data := []byte(text)
	b := sha512.Sum512(data)
	return hex.EncodeToString(b[:])
}
