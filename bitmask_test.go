package bitmask

import (
        "fmt"
        "rand"
        "reflect"
        "testing"
        "testing/quick"
)

type exp_t map[int][]bool;

func TestAbsRelExplicitData(t *testing.T) {
    bmsk := FromString(4, -3, 3, 3, ".#."+
                                    "#.#"+
                                    ".##")
    expectations := exp_t { // [y][x] -> t/f
        0: []bool {false,  true, false, /* out of bounds */ false},
        1: []bool { true, false,  true},
        2: []bool {false,  true,  true},
        // out of bounds
        3: []bool {false,  false},
    }
    doAbsRelTest(bmsk, expectations, t)
}

func TestCollidesEdgeCase(t *testing.T) {
    b1 := FromString(0, 0, 65, 1, "................................................................#")
    b2 := FromString(1, 0, 64, 1, "...............................................................#")
    if !b1.Collides(b2) {
        t.Error(b1, b2)
    }
}

func TestDumbCollidesOneElem(t *testing.T) {
    testCollidesOneElem(func(b1, b2 Bitmask) bool {return b1.dumbCollides(b2)}, t)
}

func TestCollidesOneElem(t *testing.T) {
    testCollidesOneElem(func(b1, b2 Bitmask) bool {return b1.Collides(b2)}, t)
}

func TestCollides(t *testing.T) {
    f := func(b1 Bitmask, b2 Bitmask) bool {
        return b1.Collides(b2) == b1.dumbCollides(b2)
    }
    if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

// helper functions

func (b Bitmask) Generate(rand *rand.Rand, size int) reflect.Value {
    result := Bitmask {
        X: size - rand.Intn(size),
        Y: size - rand.Intn(size),
        w: rand.Intn(size),
        h: rand.Intn(size),
    }
    result.lines = make([][]part, result.h)
    for y := 0; y<result.h; y++ {
        result.lines[y] = make([]part, result.w / sz + 1)
        for x := 0; x < result.w; x++ {
            result.SetRel(x, y, rand.Intn(size/4) == 0)
        }
    }

    return reflect.NewValue(result)
}

func doAbsRelTest(bmsk Bitmask, expectations exp_t, t *testing.T) {
    for y, yexpect := range expectations {
        for x, val := range yexpect {
            if (val != bmsk.abs(x + bmsk.X, y + bmsk.Y)) {
                t.Error(x + bmsk.X, y + bmsk.Y, bmsk, "abs should be", val)
            }
            if (val != bmsk.rel(x, y)) {
                t.Error(x, y, bmsk, "rel should be", val)
            }
        }
    }
}

func (b1 Bitmask) dumbCollides(b2 Bitmask) bool {
    for y := 0; y < b1.h; y++ {
        for x := 0; x < b1.w; x++ {
            if b1.abs(b1.X + x, b1.Y + y) && b2.abs(b1.X + x, b1.Y + y) {
                return true;
            }
        }
        for y := 0; y < b2.h; y++ {
            for x := 0; x < b2.w; x++ {
                if b1.abs(b2.X + x, b2.Y + y) && b2.abs(b2.X + x, b2.Y + y) {
                    return true;
                }
            }
        }
    }
    return false;
}

func staticallyAssertBitmaskIsGenerator(t *testing.T) quick.Generator {
    return FromString(0, 0, 1, 1, ".")
}

func testCollidesOneElem(collisionTest func(Bitmask, Bitmask)bool, t *testing.T) {
    f := func(b Bitmask) bool {
        for y := 0; y<b.h; y++ {
            for x := 0; x < b.w; x++ {
                oneEl := FromString(b.X + x, b.Y + y, 1, 1, "#")
                if b.rel(x, y) != b.dumbCollides(oneEl) {
                    fmt.Printf("failed at %d, %d, %v %v:\n", x, y, b.rel(x, y),  b.dumbCollides(oneEl))
                    return false
                }
            }
        }
        return true
    }
    if err := quick.Check(f, nil); err != nil {
		t.Error(err)
	}
}

