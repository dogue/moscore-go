package moscore

func (core *Core) andImmediate() {
	byte := core.getImmediate()
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andZeroPage(index Register) {
	byte := core.getZeroPageByte(index)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andAbsolute(index Register) {
	byte := core.getAbsoluteByte(index)
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndexedIndirect() {
	byte := core.getIndexedIndirectByte()
	core.acc &= byte
	core.setNZ(core.acc)
}

func (core *Core) andIndirectIndexed() {
	byte := core.getIndirectIndexedByte()
	core.acc &= byte
	core.setNZ(core.acc)
}
