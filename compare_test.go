package moscore

import "testing"

func TestBitZeroPage(t *testing.T) {
	bus := newBus([]uint8{0x24, 0x69})
	bus.Write(0x0069, 0b1100_0000)
	core := New(&bus)
	core.Reset()
	core.acc = 0xFF
	core.Step()

	if !core.getFlag(Negative) {
		t.Errorf("Negative flag not set when it should be")
	}

	if !core.getFlag(Overflow) {
		t.Errorf("Overflow flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Errorf("Zero flag set when it should not be")
	}
}

func TestBitAbsolute(t *testing.T) {
	bus := newBus([]uint8{0x2C, 0x37, 0x13})
	bus.Write(0x1337, 0b1100_0000)
	core := New(&bus)
	core.Reset()
	core.acc = 0xFF
	core.Step()

	if !core.getFlag(Negative) {
		t.Errorf("Negative flag not set when it should be")
	}

	if !core.getFlag(Overflow) {
		t.Errorf("Overflow flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Errorf("Zero flag set when it should not be")
	}
}
