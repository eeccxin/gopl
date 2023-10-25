package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	algorithm := flag.String("algorithm", "sha256", "Hash algorithm: sha256, sha384, or sha512")
	flag.Parse()

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
		os.Exit(1)
	}

	switch *algorithm {
	case "sha256":
		hash := sha256.Sum256(input)
		fmt.Printf("SHA256: %x\n", hash)
	case "sha384":
		hash := sha512.Sum384(input)
		fmt.Printf("SHA384: %x\n", hash)
	case "sha512":
		hash := sha512.Sum512(input)
		fmt.Printf("SHA512: %x\n", hash)
	default:
		fmt.Fprintf(os.Stderr, "Invalid algorithm: %s\n", *algorithm)
		os.Exit(1)
	}
}
