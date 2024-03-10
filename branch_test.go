package moscore

import "testing"

func TestBCC(t *testing.T) {
	bus := newBus([]uint8{0x90, 0x05})
	core := New(&bus)
	core.Reset()
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBCS(t *testing.T) {
	bus := newBus([]uint8{0xB0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Carry, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBEQ(t *testing.T) {
	bus := newBus([]uint8{0xF0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Zero, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBMI(t *testing.T) {
	bus := newBus([]uint8{0x30, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Negative, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBNE(t *testing.T) {
	bus := newBus([]uint8{0xD0, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Zero, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBPL(t *testing.T) {
	bus := newBus([]uint8{0x10, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Negative, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBVC(t *testing.T) {
	bus := newBus([]uint8{0x50, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Overflow, false)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}

func TestBVS(t *testing.T) {
	bus := newBus([]uint8{0x70, 0x05})
	core := New(&bus)
	core.Reset()
	core.setFlag(Overflow, true)
	core.Step()

	if core.pc != 0x8007 {
		t.Errorf("incorrect branch offset: wanted=%#04x, got=%#04x", 0x8007, core.pc)
	}
}
