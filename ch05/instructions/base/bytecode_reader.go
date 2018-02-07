package base

//code字段存放字节码，pc字段记录读取到了哪个字节。
type BytecodeReader struct {
	code []byte
	pc   int
}

//为了避免每次解码指令都新创建一个BytecodeReader实例，给它定义一个Reset（）方法
func (self *BytecodeReader) Rest(code []byte, pc int) {
	self.code = code
	self.pc = pc
}
