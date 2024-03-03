package moscore

func (core *Core) andImmediate() {
	byte := core.getImmediate()
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andZeroPage(index Register) {
	byte := core.getZeroPage(index)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andAbsolute(index Register) {
	byte := core.getAbsolute(index)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndexedIndirect() {
	byte := core.getIndexedIndirect()
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndirectIndexed() {
	byte := core.getIndirectIndexed()
	core.acc &= byte
	core.setNZ(core.acc)
}
