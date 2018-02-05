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

type ConstantLongInfo struct {
	val int64
}

/**
CONSTANT_Long_info {
u1 tag;
u4 high_bytes;
u4 low_bytes;
}
 */
//readInfo（）先读取一个uint64数据，然后把它转型成int64类型
func (self *ConstantLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

type ConstantDoubleInfo struct {
	val float64
}
/**
CONSTANT_Double_info {
u1 tag;
u4 high_bytes;
u4 low_bytes;
}
 */
func (self *ConstantDoubleInfo) readInfo(reader *ClassReader)  {
	bytes:=reader.readUint64()
	self.val=math.Float64frombits(bytes)
}
