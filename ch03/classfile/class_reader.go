package classfile

import "encoding/binary"

type ClassReader struct {
	data []uint8
}

// u1
func (c *ClassReader) readUnit8() uint8 {
	r := c.data[0]
	c.data = c.data[1:]
	return r
}

// u2，多字节数据考虑大小端问题
func (c *ClassReader) readUnit16() uint16 {
	r := binary.BigEndian.Uint16(c.data)
	c.data = c.data[2:]
	return r
}

// u4
func (c *ClassReader) readUnit32() uint32 {
	r := binary.BigEndian.Uint32(c.data)
	c.data = c.data[4:]
	return r
}

func (c *ClassReader) readUnit64() uint64 {
	r := binary.BigEndian.Uint64(c.data)
	c.data = c.data[8:]
	return r
}

//先读出数组的size(2个字节)，然后后面填充值
func (c *ClassReader) readUnit16s() []uint16 {
	n := c.readUnit16()
	s := make([]uint16, n)
	for _, i := range s {
		s[i] = c.readUnit16()
	}
	return s

}

//n居然设置这么大
func (c *ClassReader) readBytes(n uint32) []byte {
	bytes := c.data[:n]
	c.data = c.data[n:]
	return bytes
}
