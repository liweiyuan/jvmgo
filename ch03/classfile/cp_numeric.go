package classfile

import "math"

//
type ConstantIntegerInfo struct {
	val int32
}

/**
CONSTANT_Integer_info {
u1 tag;
u4 bytes;
}
 */
//readInfo（）先读取一个uint32数据，然后把它转型成int32类型
func (self *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

//
type ConstantFloatInfo struct {
	val float32
}

/**
CONSTANT_Float_info {
u1 tag;
u4 bytes;
}*/
//readInfo（）先读取一个uint32数据，然后把它转型成float32类型
func (self *ConstantFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = math.Float32frombits(bytes)
}
