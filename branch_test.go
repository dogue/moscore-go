package moscore

import "testing"

func TestBranchCarryClear(t *testing.T) {
	bus := newBus([]uint8{0x90, 0x05})
	core := New(&bus)
	core.Reset()
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchCarrySet(t *testing.T) {
	bus := newBus([]uint8{0xB0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Carry, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchIfEqual(t *testing.T) {
	bus := newBus([]uint8{0xF0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Zero, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchIfMinus(t *testing.T) {
	bus := newBus([]uint8{0x30, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Negative, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchNotEqual(t *testing.T) {
	bus := newBus([]uint8{0xD0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Zero, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchIfPositive(t *testing.T) {
	bus := newBus([]uint8{0x10, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Negative, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchNotOverflow(t *testing.T) {
	bus := newBus([]uint8{0x50, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Overflow, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBranchIfOverflow(t *testing.T) {
	bus := newBus([]uint8{0x70, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Overflow, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("Incorrect value in PC: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}
