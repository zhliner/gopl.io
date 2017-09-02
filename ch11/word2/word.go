// Copyright 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 305.
//!+

// Package word provides utilities for word games.
package word

import "unicode"
import "crypto/md5"
import "crypto/sha1"
import "crypto/sha256"
import "crypto/sha512"

// IsPalindrome reports whether s reads the same forward and backward.
// Letter case is ignored, as are non-letters.
func IsPalindrome(s string) bool {
	var letters []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
	for i := range letters {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

func IsPalindrome2(s string) bool {
	var letters = make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) {
			letters = append(letters, unicode.ToLower(r))
		}
	}
        n := len(letters)/2
        for i := 0; i < n; i++ {
		if letters[i] != letters[len(letters)-1-i] {
			return false
		}
	}
	return true
}

// Hzh 哈希效率比较
func Md5sum(s string) [16]byte {
    return md5.Sum([]byte(s))
}

func Sha1sum(s string) [20]byte {
    return sha1.Sum([]byte(s))
}

func Sha256sum(s string) [32]byte {
    return sha256.Sum256([]byte(s))
}

func Sha512sum(s string) [64]byte {
    return sha512.Sum512([]byte(s))
}
//!-
