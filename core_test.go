package moscore

import "testing"

type TestingBus struct {
	mem [0x10000]uint8
}

func newBus(program []uint8) TestingBus {
	bus := TestingBus{}
	bus.mem[0xfffc] = 0x00
	bus.mem[0xfffd] = 0x80

	offset := 0x8000
	for _, byte := range program {
		bus.mem[offset] = byte
		offset++
	}

	return bus
}

func (bus *TestingBus) Read(addr uint16) uint8 {
	return bus.mem[addr]
}

func (bus *TestingBus) Write(addr uint16, byte uint8) {
	bus.mem[addr] = byte
}

func TestJMP(t *testing.T) {
	progAbs := []uint8{0x4C, 0x37, 0x13}
	progInd := []uint8{0x6C, 0x42, 0x42}

	bus := newBus(progAbs)
	bus.Write(0x1337, 0x05)
	core := New(&bus)
	core.Reset()
	core.Step()

	if core.pc != 0x8008 {
		t.Errorf("[abs] incorrect value in PC: wanted=%#04x, got=%#04x", 0x8008, core.pc)
	}

	bus = newBus(progInd)
	bus.Write(0x4242, 0x37)
	bus.Write(0x4243, 0x13)
	core = New(&bus)
	core.Reset()
	core.Step()

	if core.pc != 0x1337 {
		t.Errorf("[ind] incorrect value in PC: wanted=%#04x, got=%#04x", 0x1337, core.pc)
	}
}
