package moscore

type Register *uint8

func (core *Core) loadImmediate(reg Register) {
	*reg = core.fetch()
}

func (core *Core) loadZeroPage(reg Register) {
	zp := core.fetch()
	addr := addrFromBytes(zp, 0x00)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}

func (core *Core) loadZeroPageIndexed(reg, index Register) {
	zp := core.fetch()
	zp += *index
	addr := addrFromBytes(zp, 0x00)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}

func (core *Core) loadAbsolute(reg Register) {
	adl := core.fetch()
	adh := core.fetch()
	addr := addrFromBytes(adl, adh)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}

func (core *Core) loadAbsoluteIndexed(reg, index Register) {
	adl := core.fetch()
	adh := core.fetch()
	addr := addrFromBytes(adl, adh)
	addr += uint16(*index)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}

func (core *Core) loadIndexedIndirect(reg Register) {
	adl := core.fetch()
	adl += core.idx
	addr := addrFromBytes(adl, 0x00)
	idl := core.bus.Read(addr)
	idh := core.bus.Read(addr + 1)
	addr = addrFromBytes(idl, idh)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}

func (core *Core) loadIndirectIndexed(reg Register) {
	addr := addrFromBytes(core.fetch(), 0x00)
	idl := core.bus.Read(addr)
	idh := core.bus.Read(addr + 1)
	addr = addrFromBytes(idl, idh)
	addr += uint16(core.idy)
	*reg = core.bus.Read(addr)
	core.setNZ(*reg)
}
