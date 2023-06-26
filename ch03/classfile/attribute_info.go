package classfile

type AttributesInfo interface {
	readInfo(reader *ClassReader)
}
