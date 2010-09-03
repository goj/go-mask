package bitmask

import (
        "testing"
)

func TestThatFails(t *testing.T) {
    bmsk := FromString(0, 0, 3, 3, ".x.xxx.x.")
    t.Error(bmsk)
}
