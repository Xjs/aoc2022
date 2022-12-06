package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(marker(b, 4))
	fmt.Println(marker(b, 14))
}

func marker(s []byte, l int) int {
	if len(s) < l {
		return -1
	}

	for i := l - 1; i < len(s); i++ {
		chars := make(map[byte]struct{})
		var haveEqual bool
		for j := i - (l - 1); j <= i; j++ {
			r := s[j]
			if _, ok := chars[r]; ok {
				haveEqual = true
				break
			}
			chars[r] = struct{}{}
		}
		if !haveEqual {
			// Elves start counting at 1
			return i + 1
		}
	}

	return -1
}
