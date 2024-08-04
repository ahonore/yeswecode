package main

import (
	"testing"
)

func TestSpongePattern(t *testing.T) {
	tests := map[string]struct {
		level, x1, y1, x2, y2 int
		result                string
		flag                  int
	}{
		"level2": {
			level: 2,
			x1:    5,
			y1:    4,
			x2:    7,
			y2:    6,
			result: `+0+
+00
000
`,
			flag: 17,
		},
		"level3": {
			level: 3,
			x1:    1,
			y1:    3,
			x2:    10,
			y2:    9,
			result: `00+++00000
+0+++0+00+
00+++00000
0000000000
+00+00+00+
0000000000
00000000++
`,
			flag: 240,
		},
		"level6": {
			level: 6,
			x1:    701,
			y1:    34,
			x2:    719,
			y2:    56,
			result: `+0+00+00+00+00+00+0
+000000000000000000
+000000000+++++++++
+0+00+00+0+++++++++
+000000000+++++++++
+000+++000+++++++++
+0+0+++0+0+++++++++
+000+++000+++++++++
+000000000+++++++++
+0+00+00+0+++++++++
+000000000+++++++++
+000000000000000000
+0+00+00+00+00+00+0
+000000000000000000
+000+++000000+++000
+0+0+++0+00+0+++0+0
+000+++000000+++000
+000000000000000000
+0+00+00+00+00+00+0
+000000000000000000
0000000000000000000
00+00+00+00+00+00+0
0000000000000000000
`,
			flag: 3830,
		},
	}

	for k, v := range tests {
		v := v
		t.Run(k, func(t *testing.T) {
			s := SpongePattern(v.level, v.x1, v.y1, v.x2, v.y2)
			if s != v.result {
				t.Error(s, v.result)
			}

			flag := ComputeFlag(s)
			if flag != v.flag {
				t.Error(flag, v.flag)
			}
		})
	}
}
