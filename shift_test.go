package moscore

import "testing"

func TestASL(t *testing.T) {
	bus := newBus([]uint8{0x0A})
	core := New(&bus)
	core.Reset()
	core.acc = 0b0110_0110
	core.Step()

	if core.acc != 0b1100_1100 {
		t.Errorf("incorrect shifted value: wanted=%#08b, got=%#08b", 0b1100_1100, core.acc)
	}
}
