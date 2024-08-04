# SpongeBob is Injured

You're a doctor and the next patient entering is called SpongeBob. He told you that he's hurt himself, pointing to the injury on his body. By observing his body, you noticed the shape had some repeating patterns, like a fractal.

Assuming SpongeBob's body is a square sponge, you inspect at different levels of detail.

Level 0:

```
0
```

Level 1:

```
000
0+0
000
```

Level 2:

```
000000000
0+00+00+0
000000000
000+++000
0+0+++0+0
000+++000
000000000
0+00+00+0
000000000
```

Level 3:

```
000000000000000000000000000
0+00+00+00+00+00+00+00+00+0
000000000000000000000000000
000+++000000+++000000+++000
0+0+++0+00+0+++0+00+0+++0+0
000+++000000+++000000+++000
000000000000000000000000000
0+00+00+00+00+00+00+00+00+0
000000000000000000000000000
000000000+++++++++000000000
0+00+00+0+++++++++0+00+00+0
000000000+++++++++000000000
000+++000+++++++++000+++000
0+0+++0+0+++++++++0+0+++0+0
000+++000+++++++++000+++000
000000000+++++++++000000000
0+00+00+0+++++++++0+00+00+0
000000000+++++++++000000000
000000000000000000000000000
0+00+00+00+00+00+00+00+00+0
000000000000000000000000000
000+++000000+++000000+++000
0+0+++0+00+0+++0+00+0+++0+0
000+++000000+++000000+++000
000000000000000000000000000
0+00+00+00+00+00+00+00+00+0
000000000000000000000000000
```

etc.

SpongeBob was quite accurate about the location of the injury and told you its bounding box, defined by the top-left and the bottom-right points. For example, at level 3, if the bounding box is (1,3)-(10,9), the pattern at this location would be:

```
00+++00000
+0+++0+00+
00+++00000
0000000000
+00+00+00+
0000000000
00000000++
```

The flag, representing the injury location is defined as follow:
- For each line, count each `0` and `+` and, get the line number to compute `num_of_line * num_of_0 + num_of_+`.
- Sum the value of each line to get the flag.

The line number starts at 1. In our example, the final value would be:

```
1*6+3 + 2*4+5 + 3*6+3 + 4*10+0 + 5*6+3 + 6*10+0 + 7*9+1 = 240
```

SpongeBob told you the bounding box is `(4052555153018976131,4052555153018973954)-(4052555153018976266,4052555153018976266)` at level 39. Give the flag of the injured pattern.
