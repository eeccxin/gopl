package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	algorithm := flag.String("algorithm", "sha256", "Hash algorithm: sha256, sha384, or sha512")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		switch *algorithm {
		case "sha256":
			hash := sha256.Sum256([]byte(input))
			fmt.Printf("SHA256: %x\n", hash)
		case "sha384":
			hash := sha512.Sum384([]byte(input))
			fmt.Printf("SHA384: %x\n", hash)
		case "sha512":
			hash := sha512.Sum512([]byte(input))
			fmt.Printf("SHA512: %x\n", hash)
		default:
			fmt.Fprintf(os.Stderr, "Invalid algorithm: %s\n", *algorithm)
			os.Exit(1)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}
}
