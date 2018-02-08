package stores

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/**
和加载指令刚好相反，存储指令把变量从操作数栈顶弹出，然
后存入局部变量表
 */
//// Store long into local variable
type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	_store(frame, self.Index)
}

type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	_store(frame, 0)
}

type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	_store(frame, 1)
}

type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	_store(frame, 2)
}

type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	_store(frame, 3)
}

//同样定义一个函数供5条指令使用
//todo
func _store(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}
