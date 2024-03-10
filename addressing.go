package moscore

// Fetch next byte as operand
func (core *Core) getImmediateByte() uint8 {
	return core.fetch()
}

func (core *Core) getZeroPageAddr(offset uint8) uint16 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpLow += offset
	return addrFromBytes(zpLow, zpHigh)
}

func (core *Core) getZeroPageByte(offset uint8) uint8 {
	addr := core.getZeroPageAddr(offset)
	return core.bus.Read(addr)
}

func (core *Core) getAbsoluteAddr(offset uint8) uint16 {
	low := core.fetch()
	high := core.fetch()
	addr := addrFromBytes(low, high)
	addr += uint16(offset)
	return addr
}

func (core *Core) getAbsoluteByte(offset uint8) uint8 {
	addr := core.getAbsoluteAddr(offset)
	return core.bus.Read(addr)
}

func (core *Core) getIndirectXAddr() uint16 {

	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpLow += core.idx
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	return addrFromBytes(targetLow, targetHigh)
}

func (core *Core) getIndirectXByte() uint8 {
	addr := core.getIndirectXAddr()
	return core.bus.Read(addr)
}

func (core *Core) getIndirectYAddr() uint16 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	targetAddr := addrFromBytes(targetLow, targetHigh)
	targetAddr += uint16(core.idy)

	return targetAddr
}

func (core *Core) getIndirectYByte() uint8 {
	addr := core.getIndirectYAddr()
	return core.bus.Read(addr)
}
