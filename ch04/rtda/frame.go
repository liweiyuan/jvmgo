package rtda

/**
lower字段用来实现链表数据结构，localVars字段保存
局部变量表指针，operandStack字段保存操作数栈指针。
 */
type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

//执行方法所需的局部变量表大小和操作数栈深度是由编译器预先计算好的，存储在class文件method_info结构的Code属性中
func newFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:newLocalVars(maxLocals),
		operandStack:newOperandStack(maxStack),
	}
}
