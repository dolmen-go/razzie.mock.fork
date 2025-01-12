package mock

const (
	// RAX, RBX, RCX, RDI, RSI, R8, R9, R10, R11.
	intArgRegs = 9

	// X0 -> X14.
	floatArgRegs = 15

	// We use SSE2 registers which support 64-bit float operations.
	floatRegSize = 8

	// PtrSize is the size of a pointer in bytes - unsafe.Sizeof(uintptr(0)) but as an ideal constant.
	// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
	ptrSize = 4 << (^uintptr(0) >> 63)

	bigEndian = false
)
