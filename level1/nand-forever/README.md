# NAND Forever!

Every operation in a computer are made of logical operators, like **OR**, **AND**, **NOT** and so on... However, the **NAND** operator has the particularity to be [Turing-complete](https://en.wikipedia.org/wiki/Turing_completeness). It means that all operations in a computer could be done using only this logical operator. The **NAND** operator is equivalent to a logical inverted **AND** operation.

Given two 32-bit digit numbers, return the resulting 32-bit digit number after performing **NAND** operations on bits between first and second numbers. Be aware that an operation may overlap a previous operation.

Binary numbers are written using the following logical gate format:
- `|` gate is plugged (bit is 1)
- `.` gate is unplugged (bit is 0)
Example: `|.|..|..` stands for the binary number `10100100`.

The input file is formatted this way:
- *Line 1*: the first binary digit number
- *Line 2*: the second binary digit number
- *Other lines*:  `n` `m` `r`
with `n` (resp.  `m`), the bit index of the first (resp. second) binary digit number to perform the **NAND** on. `r` is the index of the bit in the resulting number.

*By default, all logical gates of the resulting number are unplugged.*

Example of such an input of two 8-bit numbers for clarity:

```
|..|..|.
||..|.||
0 2 0
7 7 1
2 6 5
```

The flag will be: `..|....|`

What is the flag after operations from the [input file](./input.txt)?
