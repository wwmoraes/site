package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"lukechampine.com/blake3"
)

const (
	hashDigestSize = 32
	minArgs        = 2
)

func main() {
	if len(os.Args) < minArgs {
		fmt.Fprintln(os.Stderr, "usage: b3sum <FILE-1> [FILE-2..FILE-N]")
		os.Exit(2)
	}

	for _, filename := range os.Args[1:] {
		hash := blake3.New(hashDigestSize, nil)

		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("failed to read file %q: %w", filename, err))
			os.Exit(1)
		}

		data = base64.StdEncoding.AppendEncode(nil, data)

		fmt.Printf("%s\tbase64:%s\n", filename, string(data))

		data = append(data, []byte(strings.TrimPrefix(filepath.Ext(filename), "."))...)

		_, err = hash.Write(data)
		if err != nil {
			fmt.Fprintln(os.Stderr, fmt.Errorf("failed to hash data: %w", err))
		}

		fmt.Printf("%s\t%s\n", hex.EncodeToString(hash.Sum(nil))[:32], filename)
	}
}
