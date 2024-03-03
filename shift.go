package moscore

func (core *Core) aslAccumulator() {
	core.setFlag(Carry, (core.acc&(1<<7) != 0))
	core.acc <<= 1
	core.setNZ(core.acc)
}

func (core *Core) aslZeroPage(index Register) {
	addr := core.getZeroPageAddr(index)
	byte := core.bus.Read(addr)
	core.setFlag(Carry, (byte&(1<<7) != 0))
	byte <<= 1
	core.setNZ(byte)
	core.bus.Write(addr, byte)
}

func (core *Core) aslAbsolute(index Register) {
	addr := core.getAbsoluteAddr(index)
	byte := core.bus.Read(addr)
	core.setFlag(Carry, (byte&(1<<7) != 0))
	byte <<= 1
	core.setNZ(byte)
	core.bus.Write(addr, byte)
}
