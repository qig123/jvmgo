package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributesInfo
}

// 读取字段表或方法表
func readMembers(reader *ClassReader) {

}
func readMember(reader *ClassReader, cp ConstantPool) {

}

// getter
func (m *MemberInfo) AccessFlags() uint16 {
	return m.accessFlags
}
func (m *MemberInfo) Name() {
	//return
}
