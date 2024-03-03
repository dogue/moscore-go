package moscore

func (core *Core) addImmediate() {
	byte := core.fetch()
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addZeroPage() {
	zp := core.fetch()
	addr := addrFromBytes(zp, 0x00)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addZeroPageX() {
	zp := core.fetch()
	zp += core.idx
	addr := addrFromBytes(zp, 0x00)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addAbsolute() {
	adl := core.fetch()
	adh := core.fetch()
	addr := addrFromBytes(adl, adh)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addAbsoluteIndexed(index Register) {
	adl := core.fetch()
	adh := core.fetch()
	addr := addrFromBytes(adl, adh)
	addr += uint16(*index)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addIndexedIndirect() {
	zp := core.fetch()
	zp += core.idx
	addr := addrFromBytes(zp, 0x00)
	adl := core.bus.Read(addr)
	adh := core.bus.Read(addr + 1)
	addr = addrFromBytes(adl, adh)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}

func (core *Core) addIndirectIndexed() {
	zp := core.fetch()
	addr := addrFromBytes(zp, 0x00)
	adl := core.bus.Read(addr)
	adh := core.bus.Read(addr + 1)
	addr = addrFromBytes(adl, adh)
	addr += uint16(core.idy)
	byte := core.bus.Read(addr)
	carry := core.getFlagUint8(Carry)
	sum, carry := add8(core.acc, byte, carry)
	core.acc = sum
	core.setFlag(Carry, carry == 1)
	core.setNZ(core.acc)
}
