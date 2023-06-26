package classfile

import (
	"fmt"
)

type ClassFile struct {
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16        //一个经过某种算法得出的bitmark
	thisClass    uint16        //常量池索引
	superClass   uint16        //常量池索引
	interfaces   []uint16      //接口索引表，表中存放的是常量池索引，给出该类实现的所有接口的名字
	fieles       []*MemberInfo //结构体指针数组，保存的是地址
	methods      []*MemberInfo
	attributes   []AttributesInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}
func (c *ClassFile) read(reader *ClassReader) {
	c.readAndCheckMagic(reader)
	c.readAndCheckVersion(reader)
	//class文件是类还是接口,访问级别
	c.accessFlags = reader.readUnit16()
	//类索引
	c.thisClass = reader.readUnit16()
	c.superClass = reader.readUnit16()
	c.interfaces = reader.readUnit16s()
}
func (c *ClassFile) readAndCheckMagic(reader *ClassReader) {
	m := reader.readUnit32()
	if m != 0xCAFEBABE {
		panic("jave.lang.ClassFormatError:magic!")
	}
	c.magic = m
}
func (c *ClassFile) readAndCheckVersion(reader *ClassReader) {
	c.minorVersion = reader.readUnit16()
	c.majorVersion = reader.readUnit16()
}

// 从常量池里面找 类名
func (c *ClassFile) ClassName() {
	//return c.constantPool.
}

// 从常量池查找超类名
func (c *ClassFile) SuperClassName() {

}

// 从常量池查找接口名
func (c *ClassFile) InterfaceNames() {

}

// expose getter method
func (c *ClassFile) AccessFlags() uint16 {
	return c.accessFlags
}

func (c *ClassFile) MinorVersion() uint16 {
	return c.minorVersion
}

func (c *ClassFile) MajorVersion() uint16 {
	return c.majorVersion
}
func (c *ClassFile) ConstantPool() ConstantPool {
	return c.constantPool
}
func (c *ClassFile) Fields() []*MemberInfo {
	return c.fieles
}
func (c *ClassFile) Methods() []*MemberInfo {
	return c.methods
}
