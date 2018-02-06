package rtda

//实现线程私有的运行时数据区
type Thread struct {
	pc    int
	stack *Stack
}

//newStack（）函数创建Stack结构体实例，它的参数表示要创建的Stack最多可以容纳多少帧，4.3.2节将给出这个函数的代码
func NewThread(maxLocals, maxStack uint) *Thread {
	return &Thread{
		stack: newStack(1024),
	}
	/*return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}*/
}

//getter
func (self *Thread) PC() int {
	return self.pc
}

//setter
func (self *Thread) SetPc(pc int) {
	self.pc = pc
}

//PushFrame（）和PopFrame（）方法只是调用Stack结构体的相应方法而已
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	self.stack.pop()
}
//CurrentFrame（）方法返回当前帧
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
