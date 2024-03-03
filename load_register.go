package moscore

type Register *uint8

func (core *Core) loadImmediate(reg Register) {
	*reg = core.getImmediate()
}

func (core *Core) loadZeroPage(reg, index Register) {
	*reg = core.getZeroPage(index)
	core.setNZ(*reg)
}

func (core *Core) loadAbsolute(reg, index Register) {
	*reg = core.getAbsolute(index)
	core.setNZ(*reg)
}

func (core *Core) loadIndexedIndirect(reg Register) {
	*reg = core.getIndexedIndirect()
	core.setNZ(*reg)
}

func (core *Core) loadIndirectIndexed(reg Register) {
	*reg = core.getIndirectIndexed()
	core.setNZ(*reg)
}
