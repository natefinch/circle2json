package lib

import (
	"bufio"
	"strings"
	"testing"
)

func TestScanUntil(t *testing.T) {
	r := strings.NewReader(
		`


foo


~
and more`)
	line := 0
	scanner := &fileScanner{
		line:    &line,
		Scanner: bufio.NewScanner(r),
	}
	output, err := scanner.ScanUntil("~")
	if err != nil {
		t.Fatal(err)
	}
	if line != 7 {
		t.Errorf("miscounting lines, should be 7, got %v", line)
	}
	expected := "\n\n\nfoo\n\n\n"
	if output != expected {
		t.Fatalf("expected output %q but got %q", expected, output)
	}
}

func TestScanLongDoorDesc(t *testing.T) {
	r := strings.NewReader(
		`#10600
Entrance to the Great Chessboard~
You find yourself looking up at a huge archway that looks like an entrance





to what appears to be a gigantic chessboard.  Beyond the archway you see a





checkerboard of black and white squares.





~
106 0 2
0d0+0 0 0
D1
~
A black square.





~
0 -1 10601
D3
~
~
0 -1 15003
S
`)

	line := 0
	scanner := &fileScanner{
		line:    &line,
		Scanner: bufio.NewScanner(r),
	}
	scanner.Scan()
	_, err := scanRoom(scanner)
	if err != nil {
		t.Fatal(err)
	}
	if line != 38 {
		t.Errorf("expected line count to be 37, but was %v", line)
	}
}

func TestScanDoorFlags(t *testing.T) {
	r := strings.NewReader(
		`#10813
New Years Day Quest Staging Area~
This is the room where you get yer group together to enter into time.  If
(when) you die.. tell MotherTime "transme" and she will bring you back to
this room.  This is the only room that you can recall, relocate or quit from.
Once the quest begins, there will not be anyone else allowed to join.. unless
the group that is present votes to allow them to join.  If you leave early
for what ever reason you will not get "your share" of the quest points.
~
108 32792 0
0d0+0 0 0
D5
~
time~
3 -1 10801
S
#10814
~
~
108 14360 0
0d0+0 0 0
D7
~
~
0 -1 10808
S
`)
	line := 0
	scanner := &fileScanner{
		line:    &line,
		Scanner: bufio.NewScanner(r),
	}
	scanner.Scan()
	_, err := scanRoom(scanner)
	if err != nil {
		t.Fatal(err)
	}
}
