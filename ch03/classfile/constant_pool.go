package classfile

type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUnit16()) //读出有多少个常量
	cp := make([]ConstantInfo, cpCount)

}
