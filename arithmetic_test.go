package moscore

import (
	// "log"
	"testing"
)

func TestAddImmediate(t *testing.T) {
	bus := newBus([]uint8{0x69, 0x05})
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect sum: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAddZeroPage(t *testing.T) {
	bus := newBus([]uint8{0x65, 0x42})
	bus.Write(0x0042, 0x05)
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect sum: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAddZeroPageX(t *testing.T) {
	bus := newBus([]uint8{0x75, 0x40})
	bus.Write(0x0042, 0x05)
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.idx = 0x02
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect sum: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAddAbsolute(t *testing.T) {
	bus := newBus([]uint8{0x6D, 0x37, 0x13})
	bus.Write(0x1337, 0x05)
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.Step()

	if core.acc != 0x0a {
		t.Errorf("incorrect sum: wanted=%#02x, got=%#02x", 0x0a, core.acc)
	}
}

func TestAddAbsoluteIndexed(t *testing.T) {
	// X
	bus := newBus([]uint8{0x7D, 0x40})
	bus.Write(0x0042, 0x05)
	core := New(&bus)
	core.Reset()
	core.acc = 0x05
	core.idx = 0x02
	core.Step()

	if core.acc != 0x0a {
		t.Errorf("incorrect sum: wanted=%#02x, got=%#02x", 0x0a, core.acc)
	}

	// Y
	bus = newBus([]uint8{0x79, 0x40})
	bus.Write(0x0042, 0x05)
	core = New(&bus)
	core.Reset()
	core.acc = 0x05
	core.idy = 0x02
	core.Step()

	if core.acc != 0x0a {
		t.Errorf("incorrect sum: wanted=%#02x, got=%#02x", 0x0a, core.acc)
	}
}
