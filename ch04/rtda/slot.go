package rtda

//局部变量表是按索引访问的，所以很自然，可以把它想象成一个数组

type Slot struct {
	num int32
	ref *Object
}
