package classfile

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags uint16
	thisClass uint16
	superClass uint16
	interfaces []uint16
	fields []*MemberInfo
	methods []*MemeberInfo
	attributes []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile,err error) {

}

func (self *ClassFile) read(reader *ClassReader){

}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader){

}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {

}
//getter() minor
func (self *ClassFile) MinorVersion() uint16  {

}

//getter() major
func (self *ClassFile) MajorVersion() uint16{

}

//getter() contantPool
func (self *ClassFile) ConstantPool() ConstantPool{

}

//getter() accessFlags
func (self *ClassFile) AccseeFlags() uint16{

}
//getter() fields
func (self *ClassFile) Fields() []*memberInfo{
}

//getter() methods
func (self *ClassFile) Methods() []*MemberInfo  {

}

func (self *ClassFile) ClassName() string  {
}

func (self *ClassFile) SuperClassName() string{

}

func (self *ClassFile) InterfaceNames() []string{

}
