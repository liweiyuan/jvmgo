package classfile

//量表示类或者接口的符号引用
type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

/**
CONSTANT_Class_info {
u1 tag;
u2 name_index;
}
 */
func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
