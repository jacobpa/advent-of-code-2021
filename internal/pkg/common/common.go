package common

import (
	"bufio"
	"log"
	"os"
)

func ParseFile(filePath string) ([]string, error) {
	out := make([]string, 0)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Couldn't close input file: %v", err)
		}
	}(f)

	fScanner := bufio.NewScanner(f)
	for fScanner.Scan() {
		out = append(out, fScanner.Text())
	}

	return out, nil
}

func ParseStdin() ([]string, error) {
	out := make([]string, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		out = append(out, s.Text())
	}

	return out, nil
}
