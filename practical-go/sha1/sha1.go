package main

import (
	"compress/gzip"
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println(SHA1Sig("http.log.gz"))
	fmt.Println(SHA1Sig("sha1.go"))
}

// cat http.log.gz | gunzip | sha1sum
func SHA1Sig(fileName string) (string, error) {
	// cat http.log.gz
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// | gunzip if decompression is needed
	r, err := initReader(file, fileName)
	if err != nil {
		return "", err
	}
	defer r.Close()

	// | sha1sum
	w := sha1.New()

	if _, err := io.Copy(w, r); err != nil {
		return "", fmt.Errorf("%q - copy: %w", fileName, err)
	}

	sig := w.Sum(nil)

	return fmt.Sprintf("%x", sig), nil
}

// Decompress only if filename ends in ".gz"
func initReader(file *os.File, fileName string) (io.ReadCloser, error) {
	if strings.HasSuffix(fileName, ".gz") {
		// | gunzip
		r, err := gzip.NewReader(file)
		if err != nil {
			return nil, fmt.Errorf("init gzip reader: %w", err)
		}

		return r, nil
	}

	return file, nil
}
