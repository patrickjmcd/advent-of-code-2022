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

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func StringSliceEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func StringSliceAllSameChar(a []string) bool {
	if len(a) == 0 {
		return false
	}
	for i := range a {
		if a[i] != a[0] {
			return false
		}
	}
	return true
}

func StringSliceAllUnique(a []string) bool {
	if len(a) == 0 {
		return false
	}
	seen := make(map[string]bool)
	for _, v := range a {
		if seen[v] {
			return false
		}
		seen[v] = true
	}
	return true
}
