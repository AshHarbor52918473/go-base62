// Package base62 encodes non-negative integers to Base62, dependency-free.
package base62

import "strings"

const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func Encode(n uint64) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{alphabet[n%62]}, b...)
		n /= 62
	}
	return string(b)
}

func Decode(s string) uint64 {
	var n uint64
	for _, c := range s {
		n = n*62 + uint64(strings.IndexRune(alphabet, c))
	}
	return n
}
