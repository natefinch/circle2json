package lib

import (
	"fmt"
	"strconv"
)

// BitsTonames converts a bitvector into a list of bit names
func BitsToNames(vector string, maxbit int, bits map[int]string, chars map[rune]string) ([]string, error) {
	values := []string{}
	if num, err := strconv.Atoi(vector); err == nil {
		// number-style bitvector
		for bit := 1; bit <= maxbit; bit = bit << 1 {
			if num&bit != 0 {
				values = append(values, bits[bit])
			}
		}
		return values, nil
	}
	for _, r := range []rune(vector) {
		s, ok := chars[r]
		if !ok {
			return nil, fmt.Errorf("unknown bit vector letter: %v", r)
		}
		values = append(values, s)
	}
	return values, nil
}
