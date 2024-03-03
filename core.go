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
	case 0x65:
		core.addZeroPage()

	case 0x69:
		core.addImmediate()

	case 0x6D:
		core.addAbsolute()

	case 0x75:
		core.addZeroPageX()

	case 0x79:
		core.addAbsoluteIndexed(&core.idy)

	case 0x7D:
		core.addAbsoluteIndexed(&core.idx)

	case 0xA0:
		core.loadImmediate(&core.idy)

	case 0xA1:
		core.loadIndexedIndirect(&core.acc)

	case 0xA2:
		core.loadImmediate(&core.idx)

	case 0xA4:
		core.loadZeroPage(&core.idy)

	case 0xA5:
		core.loadZeroPage(&core.acc)

	case 0xA6:
		core.loadZeroPage(&core.idx)

	case 0xA9:
		core.loadImmediate(&core.acc)

	case 0xAC:
		core.loadAbsolute(&core.idy)

	case 0xAD:
		core.loadAbsolute(&core.acc)

	case 0xAE:
		core.loadAbsolute(&core.idx)

	case 0xB1:
		core.loadIndirectIndexed(&core.acc)

	case 0xB4:
		core.loadZeroPageIndexed(&core.idy, &core.idx)

	case 0xB5:
		core.loadZeroPageIndexed(&core.acc, &core.idx)

	case 0xB6:
		core.loadZeroPageIndexed(&core.idx, &core.idy)

	case 0xB9:
		core.loadAbsoluteIndexed(&core.acc, &core.idy)

	case 0xBC:
		core.loadAbsoluteIndexed(&core.idy, &core.idx)

	case 0xBD:
		core.loadAbsoluteIndexed(&core.acc, &core.idx)

	case 0xBE:
		core.loadAbsoluteIndexed(&core.idx, &core.idy)
	}
}
