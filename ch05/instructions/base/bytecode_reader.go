package base

//code字段存放字节码，pc字段记录读取到了哪个字节。
type BytecodeReader struct {
	code []byte
	pc   int
}

//为了避免每次解码指令都新创建一个BytecodeReader实例，给它定义一个Reset（）方法
func (self *BytecodeReader) Rest(code []byte, pc int) {
	self.code = code
	self.pc = pc
}

//读取指令(uint8)
func (self *BytecodeReader) ReadUint8() uint8 {
	i := self.code[self.pc]
	self.pc++
	return i
}

//用过uint8转化得到
func (self *BytecodeReader) ReadInt8() int8 {
	return int8(self.ReadUint8())
}

//ReadUint16连续读取2个字节
func (self *BytecodeReader) ReadUint16() uint16 {
	byte1 := uint16(self.ReadUint8())
	byte2 := uint16(self.ReadUint8())
	return (byte1 << 8) | byte2
}

func (self *BytecodeReader) ReadInt16() int16 {
	return int16(self.ReadUint16())
}

//ReadInt32（）方法连续读取4字节
func (self *BytecodeReader) ReadInt32() int32 {
	byte1 := int32(self.ReadUint8())
	byte2 := int32(self.ReadUint8())
	byte3 := int32(self.ReadUint8())
	byte4 := int32(self.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4

}
