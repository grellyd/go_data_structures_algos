package main

import "unicode/utf8"

/*
// IsUnique checks if a string has all unique characters
func IsUnique(s string) bool {
	for i, char := range s {
		for j := i+1; j < len(s); j++ {
			if char == s[j] {
				return false
			}
		}
	}
	return true
}
*/


// IsUniqueRange checks if a string has all unique characters
func IsUniqueRange(s string) bool {
	for i, r1 := range s {
		for _, r2 := range s[i+1:] {
            if r1 == r2 {
                return false
            }
        }
	}
	return true
}

// IsUniqueDecoding checks if a string has all unique characters
func IsUniqueDecoding(s string) bool {
	for i, r1 := range s {
		b := []byte(s[i+1:])
        for len(b) > 0 {
            r2, size := utf8.DecodeRune(b)
            if r1 == r2 {
                return false
            }
            b = b[size:]
        }
	}
	return true
}

// IsUniqueDecoded checks if a string has all unique characters
func IsUniqueDecoded(s string) bool {
    runes := runeArray(s)
	for i, r := range s {
		for j := i+1; j < len(runes); j++ {
			if r == runes[j] {
				return false
			}
		}
	}
	return true
}

func runeArray(s string) (runes []rune) {
    b := []byte(s)
    for len(b) > 0 {
        r, size := utf8.DecodeRune(b)
        runes = append(runes, r)
        b = b[size:]
    }
    return runes
}
