package moscore

func (core *Core) getImmediate() uint8 {
	return core.fetch()
}

func (core *Core) getZeroPageAddr(index Register) uint16 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)

	if index != nil {
		zpLow += *index
	}

	return addrFromBytes(zpLow, zpHigh)
}

func (core *Core) getZeroPageByte(index Register) uint8 {
	addr := core.getZeroPageAddr(index)
	return core.bus.Read(addr)
}

func (core *Core) getAbsoluteAddr(index Register) uint16 {
	low := core.fetch()
	high := core.fetch()
	addr := addrFromBytes(low, high)

	if index != nil {
		addr += uint16(*index)
	}

	return addr
}

func (core *Core) getAbsoluteByte(index Register) uint8 {
	addr := core.getAbsoluteAddr(index)
	return core.bus.Read(addr)
}

func (core *Core) getIndexedIndirectAddr() uint16 {

	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpLow += core.idx
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	return addrFromBytes(targetLow, targetHigh)
}

func (core *Core) getIndexedIndirectByte() uint8 {
	addr := core.getIndexedIndirectAddr()
	return core.bus.Read(addr)
}

func (core *Core) getIndirectIndexedAddr() uint16 {

	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	targetAddr := addrFromBytes(targetLow, targetHigh)
	targetAddr += uint16(core.idy)

	return targetAddr
}

func (core *Core) getIndirectIndexedByte() uint8 {
	addr := core.getIndirectIndexedAddr()
	return core.bus.Read(addr)
}
