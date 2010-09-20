package bitmask

type BitmaskTemplate interface {
	CollidesAbs(x, y int) bool
	Left() int
	Right() int
	Top() int
	Bottom() int
}

type StringTemplate struct {
	x, y, w, h int
	str        string
}

func (sc StringTemplate) CollidesAbs(x, y int) bool {
	relX, relY := x-sc.x, y-sc.y
	idx := sc.w*relY + relX
	return idx >= 0 && idx < len(sc.str) && sc.str[idx] != ' ' && sc.str[idx] != '.'
}

func (sc StringTemplate) Left() int   { return sc.x }
func (sc StringTemplate) Right() int  { return sc.x + sc.w - 1 }
func (sc StringTemplate) Top() int    { return sc.y }
func (sc StringTemplate) Bottom() int { return sc.y + sc.h - 1 }

func MakeBitmask(bt BitmaskTemplate) *Bitmask {
	l, r := bt.Left(), bt.Right()
	t, b := bt.Top(), bt.Bottom()
	w := r - l + 1
	h := b - t + 1
	result := &Bitmask{l, t, w, h, make([][]part, h)}
	for y := t; y <= b; y++ {
		col := uint(0)
		row := make([]part, w>>sft+1)
		result.lines[y-t] = row
		for x := l; x <= r; x++ {
			if bt.CollidesAbs(x, y) {
				row[col>>sft] |= part(1 << (col & msk))
			}
			col++
		}
	}
	return result
}

