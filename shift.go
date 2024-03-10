package moscore

func (core *Core) aslAccumulator() {
	core.setFlag(Carry, (core.acc&(1<<7) != 0))
	core.acc <<= 1
	core.setNZ(core.acc)
}

func (core *Core) aslZeroPage(offset uint8) {
	addr := core.getZeroPageAddr(offset)
	byte := core.bus.Read(addr)
	core.setFlag(Carry, (byte&(1<<7) != 0))
	byte <<= 1
	core.setNZ(byte)
	core.bus.Write(addr, byte)
}

func (core *Core) aslAbsolute(offset uint8) {
	addr := core.getAbsoluteAddr(offset)
	byte := core.bus.Read(addr)
	core.setFlag(Carry, (byte&(1<<7) != 0))
	byte <<= 1
	core.setNZ(byte)
	core.bus.Write(addr, byte)
}
