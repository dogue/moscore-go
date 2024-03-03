package moscore

import "testing"

func TestAndImmediate(t *testing.T) {
	bus := newBus([]uint8{0x29, 0x0A})
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndZeroPage(t *testing.T) {
	bus := newBus([]uint8{0x25, 0x42})
	bus.Write(0x0042, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndZeroPageX(t *testing.T) {
	bus := newBus([]uint8{0x35, 0x40})
	bus.Write(0x0042, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.idx = 0x02
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndAbsolute(t *testing.T) {
	bus := newBus([]uint8{0x2D, 0x37, 0x13})
	bus.Write(0x1337, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndAbsoluteX(t *testing.T) {
	bus := newBus([]uint8{0x3D, 0x33, 0x13})
	bus.Write(0x1337, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.idx = 0x04
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndAbsoluteY(t *testing.T) {
	bus := newBus([]uint8{0x39, 0x33, 0x13})
	bus.Write(0x1337, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.idy = 0x04
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndIndexedIndirect(t *testing.T) {
	bus := newBus([]uint8{0x21, 0x60})
	bus.Write(0x0069, 0x37)
	bus.Write(0x006A, 0x13)
	bus.Write(0x1337, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.idx = 0x09
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}

func TestAndIndirectIndexed(t *testing.T) {
	bus := newBus([]uint8{0x31, 0x69})
	bus.Write(0x0069, 0x33)
	bus.Write(0x006A, 0x13)
	bus.Write(0x1337, 0x0A)
	core := New(&bus)
	core.Reset()
	core.acc = 0x0F
	core.idy = 0x04
	core.Step()

	if core.acc != 0x0A {
		t.Errorf("Incorrect value in ACC: wanted=%#02x, got=%#02x", 0x0A, core.acc)
	}
}
