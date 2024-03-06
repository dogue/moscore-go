package moscore

type Bus interface {
	Read(uint16) uint8
	Write(uint16, uint8)
}

type Core struct {
	acc    uint8
	idx    uint8
	idy    uint8
	sp     uint8
	pc     uint16
	p      uint8
	halted bool
	bus    Bus
}

func New(bus Bus) *Core {
	return &Core{bus: bus}
}

func (core *Core) Reset() {
	core.acc = 0
	core.idx = 0
	core.idy = 0
	core.sp = 0xff
	core.p = 0

	// 6502 reset vector
	pcl := core.bus.Read(0xfffc)
	pch := core.bus.Read(0xfffd)
	core.pc = addrFromBytes(pcl, pch)
	core.halted = false
}

func (core *Core) Step() {
	opcode := core.fetch()
	core.decode(opcode)
}

func (core *Core) fetch() uint8 {
	byte := core.bus.Read(core.pc)
	core.pc++
	return byte
}

func (core *Core) push(byte uint8) {
	core.sp--
	addr := addrFromBytes(core.sp, 0x01)
	core.bus.Write(addr, byte)
}

func (core *Core) pull() uint8 {
	addr := addrFromBytes(core.sp, 0x01)
	core.sp++
	return core.bus.Read(addr)
}

func (core *Core) decode(opcode uint8) {
	switch opcode {
	case 0x06:
		core.aslZeroPage(nil)

	case 0x0A:
		core.aslAccumulator()

	case 0x0E:
		core.aslAbsolute(nil)

	case 0x10:
		core.branchIfPositive()

	case 0x16:
		core.aslZeroPage(&core.idx)

	case 0x18:
		core.setFlag(Carry, false)

	case 0x1E:
		core.aslAbsolute(&core.idx)

	case 0x21:
		core.andIndexedIndirect()

	case 0x24:
		core.bitZeroPage()

	case 0x25:
		core.andZeroPage(nil)

	case 0x29:
		core.andImmediate()

	case 0x2C:
		core.bitAbsolute()

	case 0x2D:
		core.andAbsolute(nil)

	case 0x30:
		core.branchIfMinus()

	case 0x31:
		core.andIndirectIndexed()

	case 0x35:
		core.andZeroPage(&core.idx)

	case 0x38:
		core.setFlag(Carry, true)

	case 0x39:
		core.andAbsolute(&core.idy)

	case 0x3D:
		core.andAbsolute(&core.idx)

	case 0x50:
		core.branchNotOverflow()

	case 0x58:
		core.setFlag(IRQD, false)

	case 0x65:
		core.addZeroPage(nil)

	case 0x69:
		core.addImmediate()

	case 0x6D:
		core.addAbsolute(nil)

	case 0x70:
		core.branchIfOverflow()

	case 0x75:
		core.addZeroPage(&core.idx)

	case 0x78:
		core.setFlag(IRQD, true)

	case 0x79:
		core.addAbsolute(&core.idy)

	case 0x7D:
		core.addAbsolute(&core.idx)

	case 0x81:
		addr := core.getIndexedIndirectAddr()
		core.bus.Write(addr, core.acc)

	case 0x85:
		addr := core.getZeroPageAddr(nil)
		core.bus.Write(addr, core.acc)

	case 0x8A:
		core.acc = core.idx
		core.setNZ(core.acc)

	case 0x8D:
		addr := core.getAbsoluteAddr(nil)
		core.bus.Write(addr, core.acc)

	case 0x90:
		core.branchCarryClear()

	case 0x91:
		addr := core.getIndirectIndexedAddr()
		core.bus.Write(addr, core.acc)

	case 0x95:
		addr := core.getZeroPageAddr(&core.idx)
		core.bus.Write(addr, core.acc)

	case 0x98:
		core.acc = core.idy
		core.setNZ(core.acc)

	case 0x99:
		addr := core.getAbsoluteAddr(&core.idy)
		core.bus.Write(addr, core.acc)

	case 0x9A:
		core.sp = core.idx

	case 0x9D:
		addr := core.getAbsoluteAddr(&core.idx)
		core.bus.Write(addr, core.acc)

	case 0xA0:
		core.loadImmediate(&core.idy)

	case 0xA1:
		core.loadIndexedIndirect(&core.acc)

	case 0xA2:
		core.loadImmediate(&core.idx)

	case 0xA4:
		core.loadZeroPage(&core.idy, nil)

	case 0xA5:
		core.loadZeroPage(&core.acc, nil)

	case 0xA6:
		core.loadZeroPage(&core.idx, nil)

	case 0xA8:
		core.idy = core.acc
		core.setNZ(core.idy)

	case 0xA9:
		core.loadImmediate(&core.acc)

	case 0xAA:
		core.idx = core.acc
		core.setNZ(core.idx)

	case 0xAC:
		core.loadAbsolute(&core.idy, nil)

	case 0xAD:
		core.loadAbsolute(&core.acc, nil)

	case 0xAE:
		core.loadAbsolute(&core.idx, nil)

	case 0xB0:
		core.branchCarrySet()

	case 0xB1:
		core.loadIndirectIndexed(&core.acc)

	case 0xB4:
		core.loadZeroPage(&core.idy, &core.idx)

	case 0xB5:
		core.loadZeroPage(&core.acc, &core.idx)

	case 0xB6:
		core.loadZeroPage(&core.idx, &core.idy)

	case 0xB8:
		core.setFlag(Overflow, false)

	case 0xB9:
		core.loadAbsolute(&core.acc, &core.idy)

	case 0xBA:
		core.idx = core.sp
		core.setNZ(core.idx)

	case 0xBC:
		core.loadAbsolute(&core.idy, &core.idx)

	case 0xBD:
		core.loadAbsolute(&core.acc, &core.idx)

	case 0xBE:
		core.loadAbsolute(&core.idx, &core.idy)

	case 0xD0:
		core.branchNotEqual()

	case 0xD8:
		core.setFlag(Decimal, false)

	case 0xF0:
		core.branchIfEqual()

	case 0xF8:
		core.setFlag(Decimal, true)
	}
}
