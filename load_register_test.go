package moscore

import "testing"

type LoadTest struct {
	program     []uint8
	expectedAcc uint8
	expectedX   uint8
	expectedY   uint8
}

func TestLoadImmediate(t *testing.T) {
	tests := []LoadTest{
		{[]uint8{0xA9, 0x69}, 0x69, 0x00, 0x00}, // LDA
		{[]uint8{0xA2, 0x69}, 0x00, 0x69, 0x00}, // LDX
		{[]uint8{0xA0, 0x69}, 0x00, 0x00, 0x69}, // LDY
	}

	for i, test := range tests {
		bus := newBus(test.program)
		core := New(&bus)
		core.Reset()
		core.Step()

		if core.acc != test.expectedAcc {
			t.Errorf("[%d] incorrect value in ACC: wanted=%#02x, got=%#02x", i, test.expectedAcc, core.acc)
		}

		if core.idx != test.expectedX {
			t.Errorf("[%d] incorrect value in IDX: wanted=%#02x, got=%#02x", i, test.expectedX, core.idx)
		}

		if core.idy != test.expectedY {
			t.Errorf("[%d] incorrect value in IDY: wanted=%#02x, got=%#02x", i, test.expectedY, core.idy)
		}
	}
}

func TestLoadZeroPage(t *testing.T) {
	tests := []LoadTest{
		{[]uint8{0xA5, 0x42}, 0x69, 0x00, 0x00}, // LDA
		{[]uint8{0xA6, 0x42}, 0x00, 0x69, 0x00}, // LDX
		{[]uint8{0xA4, 0x42}, 0x00, 0x00, 0x69}, // LDY
	}

	for i, test := range tests {
		bus := newBus(test.program)
		bus.Write(0x0042, 0x69)
		core := New(&bus)
		core.Reset()
		core.Step()

		if core.acc != test.expectedAcc {
			t.Errorf("[%d] incorrect value in ACC: wanted=%#02x, got=%#02x", i, test.expectedAcc, core.acc)
		}

		if core.idx != test.expectedX {
			t.Errorf("[%d] incorrect value in IDX: wanted=%#02x, got=%#02x", i, test.expectedX, core.idx)
		}

		if core.idy != test.expectedY {
			t.Errorf("[%d] incorrect value in IDY: wanted=%#02x, got=%#02x", i, test.expectedY, core.idy)
		}
	}
}

func TestLoadZeroPageIndexed(t *testing.T) {
	tests := []LoadTest{
		{[]uint8{0xB5, 0x40}, 0x69, 0x02, 0x02},
		{[]uint8{0xB6, 0x40}, 0x00, 0x69, 0x02},
		{[]uint8{0xB4, 0x40}, 0x00, 0x02, 0x69},
	}

	for i, test := range tests {
		bus := newBus(test.program)
		bus.Write(0x0042, 0x69)
		core := New(&bus)
		core.Reset()
		core.idx = 0x02
		core.idy = 0x02
		core.Step()

		if core.acc != test.expectedAcc {
			t.Errorf("[%d] incorrect value in ACC: wanted=%#02x, got=%#02x", i, test.expectedAcc, core.acc)
		}

		if core.idx != test.expectedX {
			t.Errorf("[%d] incorrect value in IDX: wanted=%#02x, got=%#02x", i, test.expectedX, core.idx)
		}

		if core.idy != test.expectedY {
			t.Errorf("[%d] incorrect value in IDY: wanted=%#02x, got=%#02x", i, test.expectedY, core.idy)
		}
	}
}

func TestLoadAbsolute(t *testing.T) {
	tests := []LoadTest{
		{[]uint8{0xAD, 0x37, 0x13}, 0x69, 0x00, 0x00},
		{[]uint8{0xAE, 0x37, 0x13}, 0x00, 0x69, 0x00},
		{[]uint8{0xAC, 0x37, 0x13}, 0x00, 0x00, 0x69},
	}

	for i, test := range tests {
		bus := newBus(test.program)
		bus.Write(0x1337, 0x69)
		core := New(&bus)
		core.Reset()
		core.Step()

		if core.acc != test.expectedAcc {
			t.Errorf("[%d] incorrect value in ACC: wanted=%#02x, got=%#02x", i, test.expectedAcc, core.acc)
		}

		if core.idx != test.expectedX {
			t.Errorf("[%d] incorrect value in IDX: wanted=%#02x, got=%#02x", i, test.expectedX, core.idx)
		}

		if core.idy != test.expectedY {
			t.Errorf("[%d] incorrect value in IDY: wanted=%#02x, got=%#02x", i, test.expectedY, core.idy)
		}
	}
}

func TestLoadAbsoluteIndexed(t *testing.T) {
	tests := []LoadTest{
		{[]uint8{0xBD, 0x00, 0x13}, 0x69, 0x37, 0x37}, // LDA, X
		{[]uint8{0xB9, 0x00, 0x13}, 0x69, 0x37, 0x37}, // LDA, Y
		{[]uint8{0xBE, 0x00, 0x13}, 0x00, 0x69, 0x37}, // LDX, Y
		{[]uint8{0xBC, 0x00, 0x13}, 0x00, 0x37, 0x69}, // LDY, X
	}

	for i, test := range tests {
		bus := newBus(test.program)
		bus.Write(0x1337, 0x69)
		core := New(&bus)
		core.Reset()
		core.idx = 0x37
		core.idy = 0x37
		core.Step()

		if core.acc != test.expectedAcc {
			t.Errorf("[%d] incorrect value in ACC: wanted=%#02x, got=%#02x", i, test.expectedAcc, core.acc)
		}

		if core.idx != test.expectedX {
			t.Errorf("[%d] incorrect value in IDX: wanted=%#02x, got=%#02x", i, test.expectedX, core.idx)
		}

		if core.idy != test.expectedY {
			t.Errorf("[%d] incorrect value in IDY: wanted=%#02x, got=%#02x", i, test.expectedY, core.idy)
		}
	}
}

func TestLoadIndexedIndirect(t *testing.T) {
	bus := newBus([]uint8{0xA1, 0x40})
	bus.Write(0x0042, 0x37)
	bus.Write(0x0043, 0x13)
	bus.Write(0x1337, 0x69)
	core := New(&bus)
	core.Reset()
	core.idx = 0x02
	core.Step()

	if core.acc != 0x69 {
		t.Errorf("Incorrect value in register: expected=%#02x, got=%#02x", 0x69, core.acc)
	}
}

func TestLoadIndirectIndexed(t *testing.T) {
	bus := newBus([]uint8{0xB1, 0x42})
	bus.Write(0x0042, 0x33)
	bus.Write(0x0043, 0x13)
	bus.Write(0x1337, 0x69)
	core := New(&bus)
	core.Reset()
	core.idy = 0x04
	core.Step()

	if core.acc != 0x69 {
		t.Errorf("Incorrect value in register: exepected=%#02x, got=%#02x", 0x69, core.acc)
	}
}
