// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"fmt"
	"reflect"
)

func main() {
	// counts := make(map[rune]int)    // counts of Unicode characters
	// var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	// invalid := 0                    // count of invalid UTF-8 characters

	// in := bufio.NewReader(os.Stdin)
	// for {
	// 	r, n, err := in.ReadRune() // returns rune, nbytes, error
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	unicode.IsLetter(r)
	// 	if err != nil {
	// 		fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// 	if r == unicode.ReplacementChar && n == 1 {
	// 		invalid++
	// 		continue
	// 	}
	// 	counts[r]++
	// 	utflen[n]++
	// }
	// fmt.Printf("rune\tcount\n")
	// for c, n := range counts {
	// 	fmt.Printf("%q\t%d\n", c, n)
	// }
	// fmt.Print("\nlen\tcount\n")
	// for i, n := range utflen {
	// 	if i > 0 {
	// 		fmt.Printf("%d\t%d\n", i, n)
	// 	}
	// }
	// if invalid > 0 {
	// 	fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	// }
	arr := [...]int{1: 1, 2: 2}
	fmt.Println(reflect.TypeOf(arr))
	sli := arr[1:]
	fmt.Println(reflect.TypeOf(sli))
	// zero(&arr)
	fmt.Printf("arr: %v\n", arr)
}

func zero(ptr *[2]int) {
	fmt.Printf("ptr: %x\n", ptr)
	fmt.Printf("(*ptr): %x\n", (*ptr))
	for i := range ptr {
		(*ptr)[i] = 0
	}
}

// func zero(ptr *[2]int) {
// 	*ptr = [2]int{}
// }

//!-
