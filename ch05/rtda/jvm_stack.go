package rtda

/**
用经典的链表（linked list）数据结构来实现Java虚拟机栈，这样
栈就可以按需使用内存空间，而且弹出的帧也可以及时被Go的垃
圾收集器回收
 */
//maxSize字段保存栈的容量（最多可以容纳多少帧），size字段保存栈的当前大小，_top字段保存栈顶指针
type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
//如果此时栈是空的，肯定是我们的虚拟机有bug，调用panic（）函数终止程序执行即可
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top:=self._top
	self._top=top.lower
	top.lower=nil
	self.size--
	return top
}
//top（）方法只是返回栈顶帧，但并不弹出
func (self *Stack) top() *Frame {
	if self._top==nil{
		panic("jvm stack is empty..")
	}
	return  self._top
}
