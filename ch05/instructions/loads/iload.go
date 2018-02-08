package loads

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/**
加载指令从局部变量表获取变量，然后推入操作数栈顶。加载
指令共33条，按照所操作变量的类型可以分为6类：aload系列指令
操作引用类型变量、dload系列操作double类型变量、fload系列操作
float变量、iload系列操作int变量、lload系列操作long变量、xaload操
作数组。本节实现其中的25条，数组和xaload系列指令将在第8章讨
论。
 */

//// Load int from local variable
type ILOAD struct {
	base.Index8Instruction
}

//iload指令的索引来自操作数，其Execute（）方法如下
func (self *ILOAD) Execute(frame *rtda.Frame) {
	_iload(frame, uint(self.Index))
}

type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	_iload(frame, 0)
}

type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	_iload(frame, 1)
}

type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	_iload(frame, 2)
}

type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	_iload(frame, 3)
}


//为了避免重复代码，定义一个函数供iload系列指令使用
//todo
//这里与源代码有差别，后续研究一下
func _iload(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}
