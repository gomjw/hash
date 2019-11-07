package hash

import (
	"fmt"
	"testing"
)

func TestGetHash(t *testing.T) {

	testCases := []struct {
		name     string
		expected string
		result   string
	}{
		{"MD5", GetMD5("Hello World"), "b10a8db164e0754105b7a99be72e3fe5"},
		{"SHA1", GetSHA1("Hello World"), "0a4d55a8d778e5022fab701977c5d840bbc486d0"},
		{"SHA256", GetSHA256("Hello World"), "a591a6d40bf420404a011733cfb7b190d62c65bf0bcda32b57b277d9ad9f146e"},
		{"SHA512", GetSHA512("Hello World"), "2c74fd17edafd80e8447b0d46741ee243b7eb74dd2149a0ab1b9246fb30382f27e853d8585719e0e67cbda0daa8f51671064615d645ae27acb15bfb1447f459b"},
	}

	for _, h := range testCases {

		t.Run(fmt.Sprintf("%s", h.name), func(t *testing.T) {
			if h.result != h.expected {
				t.Errorf("Expected: %s | Got: %s", h.expected, h.result)
			}
		})

	}
}
