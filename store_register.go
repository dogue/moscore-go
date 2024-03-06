package moscore

func (core *Core) storeZeroPage(reg, index Register) {
	addr := core.getZeroPageAddr(index)
	core.bus.Write(addr, *reg)
}
