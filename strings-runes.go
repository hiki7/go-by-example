package main

import (
	"fmt"
	"unicode/utf8"
)

// a Go string is a read-only slice of bytes
// strings are UTF-8 encode, which means that character might use multiple characters
// runes are int32 values representing unicode code points
func main() {
	const s = "สวัสดี"

	//length of string is the number of bytes used for representing each character in the string
	// in this case the length will be equal to 18, since there are non-latin characters which use multiple bytes
	fmt.Println("len:", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("rune count:", utf8.RuneCountInString(s))

	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	fmt.Println("\nUsing DecodeRuneInString")

	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
