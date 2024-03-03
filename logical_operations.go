package moscore

func (core *Core) andImmediate() {
	byte := core.fetch()
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andZeroPage(index Register) {
	zp := core.fetch()
	if index != nil {
		zp += *index
	}
	addr := addrFromBytes(zp, 0x00)
	byte := core.bus.Read(addr)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andAbsolute(index Register) {
	adl := core.fetch()
	adh := core.fetch()
	addr := addrFromBytes(adl, adh)
	if index != nil {
		addr += uint16(*index)
	}
	byte := core.bus.Read(addr)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndexedIndirect() {
	zp := core.fetch()
	zp += core.idx
	indirectAddr := addrFromBytes(zp, 0x00)
	adl := core.bus.Read(indirectAddr)
	adh := core.bus.Read(indirectAddr + 1)
	addr := addrFromBytes(adl, adh)
	byte := core.bus.Read(addr)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndirectIndexed() {
	zp := core.fetch()
	indirect := addrFromBytes(zp, 0x00)
	adl := core.bus.Read(indirect)
	adh := core.bus.Read(indirect + 1)
	addr := addrFromBytes(adl, adh)
	addr += uint16(core.idy)
	byte := core.bus.Read(addr)
	core.acc &= byte
	core.setNZ(core.acc)
}
