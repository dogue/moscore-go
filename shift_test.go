package moscore

import "testing"

func TestAslAccumulator(t *testing.T) {
	bus := newBus([]uint8{0x0A})
	core := New(&bus)
	core.Reset()
	core.acc = 0b0110_0110
	core.Step()

	if core.acc != 0b1100_1100 {
		t.Errorf("Incorrect value in ACC: wanted=%#08b, got=%#08b", 0b1100_1100, core.acc)
	}
}

func TestAslZeroPage(t *testing.T) {
	bus := newBus([]uint8{0x06, 0x42})
	bus.Write(0x0042, 0b0110_0110)
	core := New(&bus)
	core.Reset()
	core.Step()
	byte := bus.Read(0x0042)

	if byte != 0b1100_1100 {
		t.Errorf("Incorrect value: wanted=%#08b, got=%#08b", 0b1100_1100, byte)
	}
}

func TestAslZeroPageX(t *testing.T) {
	bus := newBus([]uint8{0x16, 0x40})
	bus.Write(0x0042, 0b0110_0110)
	core := New(&bus)
	core.Reset()
	core.idx = 0x02
	core.Step()
	byte := bus.Read(0x0042)

	if byte != 0b1100_1100 {
		t.Errorf("Incorrect value: wanted=%#08b, got=%#08b", 0b1100_1100, byte)
	}
}

func TestAslAbsolute(t *testing.T) {
	bus := newBus([]uint8{0x0E, 0x37, 0x13})
	bus.Write(0x1337, 0b0110_0110)
	core := New(&bus)
	core.Reset()
	core.Step()
	byte := bus.Read(0x1337)

	if byte != 0b1100_1100 {
		t.Errorf("Incorrect value: wanted=%#08b, got=%#08b", 0b1100_1100, byte)
	}
}

func TestAslAbsoluteX(t *testing.T) {
	bus := newBus([]uint8{0x1E, 0x33, 0x13})
	bus.Write(0x1337, 0b0110_0110)
	core := New(&bus)
	core.Reset()
	core.idx = 0x04
	core.Step()
	byte := bus.Read(0x1337)

	if byte != 0b1100_1100 {
		t.Errorf("Incorrect value: wanted=%#08b, got=%#08b", 0b1100_1100, byte)
	}
}
