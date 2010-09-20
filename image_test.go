package bitmask

import (
    "testing"
    "os"
    "image/png"
    "testing/quick"
    )

func fromPng(filename string) *Bitmask {
    f, err := os.Open(filename, os.O_RDONLY, 0)
    if err != nil {
        panic(err.String())
    }
    defer f.Close()
    img, err := png.Decode(f)
    if err != nil {
        panic(err.String())
    }
    return MakeBitmask(&ImageTemplate{img})
}

func TestLoadFromPng(t *testing.T) {
    strGlider := fromString(0, 0, 3, 3, ".x."+"..x"+"xxx")
    pngGlider := fromPng("test_data/slider.png")
    g := func(x, y int) bool {
        b := fromString(x % 4, y % 4, 1, 1, "x")
        return strGlider.Collides(b) == pngGlider.Collides(b)
    }
    if err := quick.Check(g, nil); err != nil {
		t.Error(err)
	}
}
