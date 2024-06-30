package main

import (
	"log"
	"math/rand"
	"os"
	"strings"
)

var convertToHomoglyph = map[rune][]rune{
	'!':  {'\uFF01'},
	'"':  {'\uff02'},
	'$':  {'\uff04'},
	'%':  {'\uff05'},
	'&':  {'\uff06'},
	'\'': {'\uff07'},
	'(':  {'\uff08'},
	')':  {'\uff09'},
	'*':  {'\uff0A'},
	'+':  {'\uff0B'},
	',':  {'\uff0C'},
	'-':  {'\uff0D'},
	'.':  {'\uff0E'},
	'/':  {'\uff0F'},
	'0':  {'\uff10'},
	'1':  {'\uff11'},
	'2':  {'\uff12'},
	'3':  {'\uff13'},
	'4':  {'\uff14'},
	'5':  {'\uff15'},
	'6':  {'\uff16'},
	'7':  {'\uff17'},
	'8':  {'\uff18'},
	'9':  {'\uff19'},
	':':  {'\uff1A'},
	';':  {'\uff1B'},
	'<':  {'\uff1C'},
	'=':  {'\uff1D'},
	'>':  {'\uff1E'},
	'?':  {'\uff1F'},
	'@':  {'\uff20'},
	'A':  {'\uff21', '\u0391', '\u0410'},
	'B':  {'\uff22', '\u0392', '\u0412'},
	'C':  {'\uff23', '\u03F9', '\u216D'},
	'D':  {'\uff24'},
	'E':  {'\uff25', '\u0395', '\u0415'},
	'F':  {'\uff26'},
	'G':  {'\uff27'},
	'H':  {'\uff28', '\u0397', '\u041D'},
	'I':  {'\uff29', '\u0399', '\u0406'},
	'J':  {'\uff2A'},
	'K':  {'\uff2B', '\u039A', '\u212A'},
	'L':  {'\uff2C'},
	'M':  {'\uff2D', '\u039C', '\u041C'},
	'N':  {'\uff2E'},
	'O':  {'\uff2F', '\u039F', '\u041E'},
	'P':  {'\uff30', '\u03A1', '\u0420'},
	'Q':  {'\uff31'},
	'R':  {'\uff32'},
	'S':  {'\uff33'},
	'T':  {'\uff34', '\u03A4', '\u0422'},
	'U':  {'\uff35'},
	'V':  {'\uff36', '\u0474', '\u2164'},
	'W':  {'\uff37'},
	'X':  {'\uff38', '\u03A7', '\u2169'},
	'Y':  {'\uff39', '\u03A5', '\u04AE'},
	'Z':  {'\uff3A'},
	'[':  {'\uff3B'},
	'\\': {'\uff3C'},
	']':  {'\uff3D'},
	'^':  {'\uff3E'},
	'_':  {'\uff3F'},
	'`':  {'\uff40'},
	'a':  {'\uff41'},
	'b':  {'\uff42'},
	'c':  {'\uff43', '\u03F2', '\u0441'},
	'd':  {'\uff44'},
	'e':  {'\uff45'},
	'f':  {'\uff46'},
	'g':  {'\uff47'},
	'h':  {'\uff48'},
	'i':  {'\uff49', '\u0456', '\u2170'},
	'j':  {'\uff4A'},
	'k':  {'\uff4B'},
	'l':  {'\uff4C'},
	'm':  {'\uff4D'},
	'n':  {'\uff4E'},
	'o':  {'\uff4F', '\u03BF', '\u043E'},
	'p':  {'\uff50'},
	'q':  {'\uff51'},
	'r':  {'\uff52'},
	's':  {'\uff53'},
	't':  {'\uff54'},
	'u':  {'\uff55'},
	'v':  {'\uff56', '\u03BD', '\u2174'},
	'w':  {'\uff57'},
	'x':  {'\uff58', '\u0445', '\u2179'},
	'y':  {'\uff59'},
	'z':  {'\uff5A'},
	'{':  {'\uff5B'},
	'|':  {'\uff5C'},
	'}':  {'\uff5D'},
	'~':  {'\uff5E'},
	' ': {
		'\u2000',
		'\u2001',
		'\u2002',
		'\u2003',
		'\u2004',
		'\u2005',
		'\u2006',
		'\u2007',
		'\u2008',
		'\u2009',
		'\u200A',
		'\u2028',
		'\u2029',
		'\u202F',
		'\u205F',
	},
}

const (
	flagFile      = "flag.txt"
	initTweetFile = "initial_tweet.txt"
	inputFile     = "input.txt"

	seed int64 = 8086614821050898000
)

func main() {
	rand.New(rand.NewSource(seed))
	bf, err := os.ReadFile(flagFile)
	if err != nil {
		log.Fatal(err)
	}

	bit, err := os.ReadFile(initTweetFile)
	if err != nil {
		log.Fatal(err)
	}

	sf := string(bf)
	sit := string(bit)

	var sb strings.Builder
	rf := []rune(sf)
	var rfid int
	for _, rit := range sit {
		if rfid >= len(rf) {
			sb.WriteRune(rit)
			continue
		}

		if rit == rf[rfid] {
			choices := convertToHomoglyph[rit]
			sb.WriteRune(choices[rand.Intn(len(choices))])
			rfid++
			continue
		}

		sb.WriteRune(rit)
	}
	if err = os.WriteFile(inputFile, []byte(sb.String()), 0644); err != nil {
		log.Fatal(err)
	}
}
