package lib

import (
	"bufio"
	"errors"
	"strings"
)

type fileScanner struct {
	line *int
	*bufio.Scanner
}

func (f *fileScanner) Scan() bool {
	b := f.Scanner.Scan()
	if b {
		(*(f.line))++
	}
	return b
}

func (f *fileScanner) MustScan() error {
	if !f.Scan() {
		if err := f.Err(); err != nil {
			return err
		}
		return errors.New("unexpected EOF")
	}
	return nil
}

func (f *fileScanner) ScanUntil(terminator string) (string, error) {
	var lines []string
	for {
		if err := f.MustScan(); err != nil {
			return "", err
		}
		s := f.Text()
		if strings.HasSuffix(s, terminator) {
			lines = append(lines, s[:len(s)-len(terminator)])
			return strings.Join(lines, "\n"), nil
		}
		lines = append(lines, s)
	}
}

func (f *fileScanner) ScanUntilPrefix(prefixes []string) (string, error) {
	var lines []string
	for {
		if err := f.MustScan(); err != nil {
			return "", err
		}
		s := f.Text()
		for _, pref := range prefixes {
			if strings.HasPrefix(s, pref) {
				return strings.Join(lines, "\n"), nil
			}
		}
		lines = append(lines, s)
	}
}
