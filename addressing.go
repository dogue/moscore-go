package moscore

func (core *Core) getImmediate() uint8 {
	return core.fetch()
}

func (core *Core) getZeroPage(index Register) uint8 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)

	if index != nil {
		zpLow += *index
	}

	addr := addrFromBytes(zpLow, zpHigh)
	return core.bus.Read(addr)
}

func (core *Core) getAbsolute(index Register) uint8 {
	low := core.fetch()
	high := core.fetch()
	addr := addrFromBytes(low, high)

	if index != nil {
		addr += uint16(*index)
	}

	return core.bus.Read(addr)
}

func (core *Core) getIndexedIndirect() uint8 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpLow += core.idx
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	targetAddr := addrFromBytes(targetLow, targetHigh)

	return core.bus.Read(targetAddr)
}

func (core *Core) getIndirectIndexed() uint8 {
	zpLow := core.fetch()
	zpHigh := uint8(0x00)
	zpAddr := addrFromBytes(zpLow, zpHigh)

	targetLow := core.bus.Read(zpAddr)
	targetHigh := core.bus.Read(zpAddr + 1)
	targetAddr := addrFromBytes(targetLow, targetHigh)
	targetAddr += uint16(core.idy)

	return core.bus.Read(targetAddr)
}
