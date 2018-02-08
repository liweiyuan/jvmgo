/**
常量指令把常量推入操作数栈顶。常量可以来自三个地方：隐
含在操作码里、操作数和运行时常量池。
 */
package constants

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
)

//do nothing指令
type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame rtda.Frame)  {
	//什么也不做
}
