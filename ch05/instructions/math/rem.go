package math

import (
	"jvmgo/ch05/instructions/base"
	"jvmgo/ch05/rtda"
	"math"
)

/**
数学指令大致对应Java语言中的加、减、乘、除等数学运算符。
数学指令包括算术指令、位移指令和布尔运算指令等，共37条
 */
//求余指令的实现

// Remainder double
type DREM struct {
	base.NoOperandsInstruction
}

/**
Go语言没有给浮点数类型定义求余操作符，所以需要使用
math包的Mod（）函数。另外，浮点数类型因为有Infinity（无穷大）
值，所以即使是除零，也不会导致ArithmeticException异常抛出。
 */
func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	d1 := stack.PopDouble()
	d2 := stack.PopDouble()
	result := math.Mod(d1, d2) //todo
	stack.PushDouble(result)
}

// Remainder float
type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	f1 := stack.PopFloat()
	f2 := stack.PopFloat()
	result := math.Mod(float64(f1), float64(f2))
	stack.PushFloat(float32(result))
}

// Remainder int
type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	i1 := stack.PopInt()
	i2 := stack.PopInt()
	if i2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := i1 % i2
	stack.PushInt(result)
}

// Remainder long
type LREM struct{ base.NoOperandsInstruction }

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	l2 := stack.PopLong()
	l1 := stack.PopLong()
	if l2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	result := l1 % l2
	stack.PushLong(result)
}
