package bitmask

import "fmt"

type part uint64
const sz = 64
const msk = sz-1
const sft = 6 // sz = 1 << sft

type Bitmask struct {
    X, Y int
    w, h int
    lines [][]part
}

func FromString(x, y, w, h int, str string) Bitmask {
    result := Bitmask{x, y, w, h, make([][]part, h)}
    for i, r := uint(0), 0; r<h; r++ {
        row := make([]part, w >> sft + 1)
        result.lines[r] = row
        for c := 0; c < w; c++ {
            if str[i] != '.' && str[i] != ' ' {
                row[c >> sft] |= part(1 << uint(c & msk))
            }
            i++
        }
    }
    return result
}

func (b Bitmask) SetRel(x, y int, val bool) {
    if (y < 0 || y >= b.h) {return}
    if (x < 0 || x >= b.w) {return}
    if val {
        b.lines[y][x >> sft] |=  part(1 << uint(x & msk))
    } else {
        b.lines[y][x >> sft] &^= part(1 << uint(x & msk))
    }
}

func (b Bitmask) Format(f fmt.State, c int) {
    fmt.Fprintf(f, "\nBM+(%d, %d):\n", b.X, b.Y)
    for y := 0; y<b.h; y++ {
        for x := 0; x < b.w; x++ {
            if b.rel(x, y) {
                fmt.Fprint(f, "#")
            } else {
                fmt.Fprint(f, ".")
            }
        }
        fmt.Fprint(f, "\n")
    }
}

func (b1 Bitmask) Collides (b2 Bitmask) bool {
    dx, dy := b2.X - b1.X, b2.Y - b1.Y
    if (dx < 0) {return b2.Collides(b1)}
    miny, maxy := max(-dy, 0), min(b1.h - dy, b2.h)
    mini, maxi := max(-dx >> sft, 0), min((b1.w-dx) >> sft + 1, b2.w >> sft + 1)
    di, shift1, shift2 := part(dx >> sft), part(dx & msk), sz - part(dx & msk)
    for y := miny; y < maxy; y++ {
        for i := part(mini); i < part(maxi); i++ {
            bits := b1.lines[y + dy][i + di]
            if i>0 && bits << shift2 & b2.lines[y][i-1] != 0 { return true }
            if        bits >> shift1 & b2.lines[y][i]   != 0 { return true }
        }
    }
    return false
}

// wtf, wtf, wtf...
func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// private functions for testing - don't have to be fast

func (b Bitmask) abs(x, y int) bool {
    return b.rel(x - b.X, y - b.Y)
}

func (b Bitmask) rel(x, y int) bool {
    if (y < 0 || y >= b.h) {return false}
    if (x < 0 || x >= b.w) {return false}
    return b.lines[y][x >> sft] & part(1 << uint(x & msk)) != 0
}
