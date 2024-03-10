package moscore

import (
	// "log"
	"testing"
)

func TestADC(t *testing.T) {
	bus := newBus([]uint8{0x69, 0x05})
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("incorrect sum: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}
