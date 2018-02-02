package classfile

import "fmt"

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
	methods []*MemberInfo
	attributes []AttributeInfo
}

//把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile,err error) {
	defer func() {
		if r:=recover();r!=nil{
			var ok bool
			err,ok=r.(error)
			if !ok{
				err=fmt.Errorf("%v",r)
			}
		}
	}()
	cr:=&ClassReader{classData}
	cf=&ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader){
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool=readConstantPool(reader)
	self.accessFlags=reader.readUint16()
	self.thisClass=reader.readUint16()
	self.superClass=reader.readUint16()
	self.interfaces=reader.readUint16s()
	self.fields=readMembers(reader,self.constantPool)
	self.methods=readMembers(reader,self.constantPool)
	self.attributes=readAttributes(reader,self.constantPool)
}

//读取魔数
func (self *ClassFile) readAndCheckMagic(reader *ClassReader){
	magic:=reader.readUint32()
	if magic!=0xCAFEBABE{
		panic("java.lang.ClassFormatError: magic!")
	}
}
//读取版本检测
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion=reader.readUint16()
	self.majorVersion=reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46,47,48,49.50,51,52:
		if self.minorVersion==0{
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}
//getter() minor
func (self *ClassFile) MinorVersion() uint16  {
	return self.minorVersion
}

//getter() major
func (self *ClassFile) MajorVersion() uint16{
	return self.majorVersion
}

//getter() contantPool
func (self *ClassFile) ConstantPool() ConstantPool{
	return self.constantPool
}

//getter() accessFlags
func (self *ClassFile) AccseeFlags() uint16{
	return self.accessFlags
}
//getter() fields
func (self *ClassFile) Fields() []*memberInfo{
	return self.fields
}

//getter() methods
func (self *ClassFile) Methods() []*MemberInfo  {
	return self.methods
}

func (self *ClassFile) ClassName() string  {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string{
	if self.superClass>0{
		return self.constantPool.getClassName(self.superClass)
	}
	return ""//没有父类  java.lang.Object
}

func (self *ClassFile) InterfaceNames() []string{
	interfaceNames:=make([]string,len(self.interfaces))
	for i,cpIndex:=range interfaceNames{
		interfaceNames[i]=self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
