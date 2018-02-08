package stack

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

/**
pop和pop2指令将栈顶变量弹出，dup系列指令复制栈顶变量，swap指令交换栈顶的两
个变量。
 */
type POP struct {
	base.NoOperandsInstruction
}
/**
pop指令只能用于弹出int、float等占用一个操作数栈位置的变
量。
 */
func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	stack.PopSlot()
}

type POP2 struct {
	base.NoOperandsInstruction
}
/**
double和long变量在操作数栈中占据两个位置，需要使用pop2
指令弹出
 */
func (self *POP2) Execute(frame *rtda.Frame)  {
	stack:=frame.OperandStack()
	stack.PopSlot()
	stack.PopSlot()
}
