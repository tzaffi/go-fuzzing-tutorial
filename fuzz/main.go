package main

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) (string, error) {
	/*
		b := []byte(s)
		for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
			b[i], b[j] = b[j], b[i]
		}
		return string(b)
	*/
	if !utf8.ValidString(s) {
		return "", fmt.Errorf("Not a valid utf8 string")
	}
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r), nil

}

func main() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)
	fmt.Printf("original: %q\n", input)
	fmt.Printf("reversed: %q\n", rev)
	fmt.Printf("reversed again: %q\n", doubleRev)

	/*
		for pos, char := range "日本語" {
			fmt.Printf("character %c of type %T starts at byte position %d\n", char, char, pos)
		}
	*/
}
