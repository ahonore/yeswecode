package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func ComputeSpongePatch(level, x, y int) rune {
	size := int(new(big.Int).Exp(big.NewInt(3), big.NewInt(int64(level)), nil).Int64())
	return computeSpongePatch(size, x, y)
}

func computeSpongePatch(size, x, y int) rune {
	if size <= 1 { // level 0
		return '0'
	}

	patchSize := size / 3
	// test coord in center of the patch
	if (x/patchSize == 1) && (y/patchSize == 1) {
		return '+'
	}

	return computeSpongePatch(patchSize, x%patchSize, y%patchSize)
}

func ComputeFlag(s string) (sum int) {
	line := 1
	sc := bufio.NewScanner(strings.NewReader(s))
	for sc.Scan() {
		l := sc.Text()
		var num_0, num_plus int
		for _, c := range l {
			if c == '0' {
				num_0++
			} else {
				num_plus++
			}
		}
		sum += line*num_0 + num_plus
		line++
	}
	return
}

func SpongePattern(level, x1, y1, x2, y2 int) string {
	var sb strings.Builder
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			sb.WriteRune(ComputeSpongePatch(level, x, y))
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func main() {
	const level = 39
	const x1, y1, x2, y2 = 4052555153018976131, 4052555153018973954, 4052555153018976266, 4052555153018976266

	s := SpongePattern(level, x1, y1, x2, y2)
	fmt.Print(s)
	flag := ComputeFlag(s)
	fmt.Println(flag)
	f, _ := os.Create("flag.txt")
	f.WriteString(strconv.Itoa(flag))
	f.Close()
}
