// generate input file and flag file of the challenge
package main

import (
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var seed int64 = 1711531345694243359
var inputFile = "input.txt"
var flagFile = "flag.txt"

func convToGates(v uint32) (s string) {
	for i := 31; i >= 0; i-- {
		c := '.'
		if ((v >> i) & 1) == 1 {
			c = '|'
		}

		s += string(c)
	}
	return
}

func main() {
	f, _ := os.Create(inputFile)
	defer f.Close()
	r := rand.New(rand.NewSource(seed))
	n1 := r.Uint32()
	n2 := r.Uint32()
	f.WriteString(convToGates(n1))
	f.Write([]byte{'\n'})
	f.WriteString(convToGates(n2))
	f.Write([]byte{'\n'})
	var res uint32
	for i := 0; i < 8192; i++ {
		n1Id := uint32(r.Int31n(32))
		n2Id := uint32(r.Int31n(32))
		rId := uint32(r.Int31n(32))
		b := ^((n1 >> n1Id) & (n2 >> n2Id)) & 1
		if b == 1 {
			res |= (1 << rId)
		} else {
			res &= ^(1 << rId)
		}

		f.WriteString(strings.Join([]string{strconv.Itoa(int(n1Id)), strconv.Itoa(int(n2Id)), strconv.Itoa(int(rId))}, " "))
		f.Write([]byte{'\n'})
	}

	ff, _ := os.Create(flagFile)
	defer ff.Close()
	ff.WriteString(convToGates(res))
	ff.Write([]byte{'\n'})
}
