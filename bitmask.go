package bitmask

import "fmt"

type part uint64
const sz = 64
const msk = sz-1

type Bitmask struct {
    x, y int
    w, h int
    lines [][]part
}

func FromString(x, y, w, h int, str string) Bitmask {
    result := Bitmask{x, y, w, h, make([][]part, h)}
    for i, r := uint(0), 0; r<h; r++ {
        row := make([]part, w / sz + 1)
        result.lines[r] = row
        for c := 0; c < w; c++ {
            if str[i] != '.' && str[i] != ' ' {
                row[c / sz] |= part(1 << uint(c % sz))
            }
            i++
        }
    }
    return result
}

func (b Bitmask) Format(f fmt.State, c int) {
    for r := 0; r<b.h; r++ {
        row := b.lines[r]
        for c := 0; c < b.w; c++ {
            if row[c / sz] & part(1 << uint(c % sz)) != 0 {
                fmt.Fprint(f, "#")
            } else {
                fmt.Fprint(f, ".")
            }
        }
        fmt.Fprint(f, "\n")
    }
}

func Collides(b1 Bitmask, b2 Bitmask) bool {
    return true
}
