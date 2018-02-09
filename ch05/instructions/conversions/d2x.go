package conversions

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//类型转换指令大致对应Java语言中的基本类型强制转换操作。类型转换指令有共15条，将全部在本节实现。
/**
按照被转换变量的类型，类型转换指令可以分为3种：i2x系列
指令把int变量强制转换成其他类型；12x系列指令把long变量强制
转换成其他类型；f2x系列指令把float变量强制转换成其他类型；d2x
系列指令把double变量强制转换成其他类型。
 */

type D2F struct {
	base.NoOperandsInstruction
}
func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	f := float32(d)
	stack.PushFloat(f)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	i := int32(d)
	stack.PushInt(i)
}

type D2L struct {
	base.NoOperandsInstruction
}


func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d := stack.PopDouble()
	l := int64(d)
	stack.PushLong(l)
}
