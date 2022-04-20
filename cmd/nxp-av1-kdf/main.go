package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	av1 "github.com/nvx/nxp-av1-kdf"
	"os"
)

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:\n\tnxp-av1-kdf <key> <salt>")
	os.Exit(2)
}

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	key, err := hex.DecodeString(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing key hex string", err)
		usage()
	}

	salt, err := hex.DecodeString(os.Args[2])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error parsing salt hex string", err)
		usage()
	}

	derivedKey, err := av1.KDF(aes.NewCipher, key, salt)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error deriving key", err)
		usage()
	}

	fmt.Println(hex.EncodeToString(derivedKey))
}
