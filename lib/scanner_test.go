package lib

import (
	"bufio"
	"io/ioutil"
	"os"
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

func TestScanMob(t *testing.T) {
	const mobs = `#0
coodie~
a big coodie~
A big coodie attacks you!
~
It's a big coodie!
~
321615 168 1000 W3
127 75 -38 1d100+16000 10d10+70
52000 100000
8 8 2
18 100 18 18 18 18 0
8 2 0
-380 0 6 0 0 0
10 10 10 10 10
0 0 0
6 0
~
#1
puff Puff dragon~
Puff the Magic Dragon~
Puff the Magic Dragon is standing here.
~
Puff is a very large, very green, female dragon.  Her scales are an iridescent
green, and the light in the room reflects off of them, giving the whole room an
eerie green glow.  There are tiny tendrils of smoke rising from her jaws and
nostrils, and you suddenly feel very warm in her presence.  Her tail is very
long, with spikes down the back of it.  She loves to curl it tightly around
her "captive audience" and draw them close to her while they converse.
~
24587 1032 1000 W3
54 21 -10 16d10+4000 1d8+30
0 300000
8 8 2
13 0 13 13 13 13 0
0 2 0
-100 0 6 0 0 0
10 10 10 10 10
0 0 0
6 0
$
`
	f, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		f.Close()
		os.RemoveAll(f.Name())
	}()

	_, err = f.WriteString(mobs)
	if err != nil {
		t.Fatal(err)
	}
	name := f.Name()
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}

	ms, err := ParseMobFile(name)
	if err != nil {
		t.Fatal(err)
	}
	if len(ms) != 2 {
		t.Fatalf("expected 2 mobs, got %v", len(ms))
	}
}
