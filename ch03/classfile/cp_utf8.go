package classfile

type ConstantUtf8Info struct {
	str string
}

func (c *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUnit16())
	bytes := reader.readBytes(length)
	c.str = decodeMUTF8(bytes)
}

// 简化版，不能处理null字符
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
