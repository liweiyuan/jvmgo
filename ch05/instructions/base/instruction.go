package base

import "jvmgo/ch05/rtda"

//定义指令接口，然后定义一个结构体用来辅助指令解码。

/**
FetchOperands（）方法从字节码中提取操作数，Execute（）方法
执行指令逻辑。
 */
type Instruction interface {
	FetchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame)
}


//NoOperandsInstruction表示没有操作数的指令，所以没有定义任何字段
type NoOperandsInstruction struct {

}

func (self *NoOperandsInstruction) FetchOperands(reader *BytecodeReader){
	//表示什么都不做
}

//BranchInstruction表示跳转指令，Offset字段存放跳转偏移量。
type BranchInstruction struct{
	Offset int
}
//FetchOperands（）方法从字节码中读取一个uint16整数，转成int后赋给Offset字段
func (self *BranchInstruction) FetchOperands(reader *BytecodeReader){
	self.Offset=int(reader.ReadInt16())
}

//存储和加载类指令需要根据索引存取局部变量表，索引由单字节操作数给出。把这类指令抽象成Index8Instruction结构体，用Index字段表示局部变量表索引。
type Index8Instruction struct {
	Index uint
}
//FetchOperands（）方法从字节码中读取一个int8整数，转成uint后赋给Index字段
func (self *Index8Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index=uint(reader.ReadUint8())
}


//有一些指令需要访问运行时常量池，常量池索引由两字节操作数给出。把这类指令抽象成Index16Instruction结构体，用Index字段表示常量池索引
type Index16Instruction struct {
	Index uint
}

//FetchOperands（）方法从字节码中读取一个uint16整数，转成uint后赋给Index字段
func (self *Index16Instruction) FetchOperands(reader *BytecodeReader)  {
	self.Index=uint(reader.ReadUint16())
}

