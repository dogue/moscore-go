package moscore

import ()

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
