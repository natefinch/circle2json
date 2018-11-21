package lib

import (
	"bufio"
	"strings"
	"testing"
)

func TestParzeZone(t *testing.T) {
	r := strings.NewReader(
		`#0
Limbo - Internal~
0 99 10 2
*
* Mobiles
M 0 1 1 1               Puff
*
S
$

`)

	line := 0
	scanner := &fileScanner{
		line:    &line,
		Scanner: bufio.NewScanner(r),
	}
	if err := scanner.MustScan(); err != nil {
		t.Fatal(err)
	}
	_, err := scanZone(scanner)
	if err != nil {
		t.Fatal(err)
	}

}
