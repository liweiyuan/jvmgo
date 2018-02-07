package rtda

//局部变量表是按索引访问的，所以很自然，可以把它想象成一个数组

type Slot struct {
	num int32
	ref *Object
}
/**
num字段存放整数，ref字段存放引用，刚好满足我们的需求。
下面用它来实现局部变量表。
 */