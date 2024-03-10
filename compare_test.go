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
		t.Error("Negative flag not set when it should be")
	}

	if !core.getFlag(Overflow) {
		t.Error("Overflow flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Error("Zero flag set when it should not be")
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
		t.Error("Negative flag not set when it should be")
	}

	if !core.getFlag(Overflow) {
		t.Error("Overflow flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Error("Zero flag set when it should not be")
	}
}

func TestCompareImmediate(t *testing.T) {
	bus := newBus([]uint8{0xC9, 0x05})
	core := New(&bus)
	core.Reset()
	core.acc = 0x0A
	core.Step()

	if !core.getFlag(Carry) {
		t.Error("Carry flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Error("Zero flag set when it should not be")
	}

	if core.getFlag(Negative) {
		t.Error("Negative flag set when it should not be")
	}
}

func TestCompareZeroPage(t *testing.T) {
	bus := newBus([]uint8{0xC5, 0x69})
	bus.Write(0x0069, 0x05)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0A
	core.Step()

	if !core.getFlag(Carry) {
		t.Error("Carry flag not set when it should be")
	}

	if core.getFlag(Zero) {
		t.Error("Zero flag set when it should not be")
	}

	if core.getFlag(Negative) {
		t.Error("Negative flag set when it should not be")
	}
}
