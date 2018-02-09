package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//布尔运算
//boolean--Boolean
/**
布尔运算指令只能操作int和long变量，分为按位与（and）、按位
或（or）、按位异或（xor）3种
*/

type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i1 := stack.PopInt()
	i2 := stack.PopInt()
	result := i1 | i2
	stack.PushInt(result)
}

type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l1 := stack.PopLong()
	l2 := stack.PopLong()
	result := l1 | l2
	stack.PushLong(result)
}
