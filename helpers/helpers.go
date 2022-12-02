package helpers

import (
	"io"
	"os"
)

// OpenAndReadFile opens a file and returns the contents as a byte slice
func OpenAndReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}
