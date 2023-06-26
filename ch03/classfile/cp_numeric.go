package classfile

import "math"

/*
CONSTANT_Integer
CONSTANT_Float
CONSTANT_Long
结构如下
CONSTANT_Integer_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Float_info {
    u1 tag;
    u4 bytes;
}
CONSTANT_Long_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
CONSTANT_Double_info {
    u1 tag;
    u4 high_bytes;
    u4 low_bytes;
}
*/
type ConstantIntegerInfo struct {
	value int32
}
type ConstatnFloatInfo struct {
	value float32
}
type ConstatnLongInfo struct {
	value int64
}
type ConstatnDoubleInfo struct {
	value float64
}

func (c *ConstantIntegerInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	c.value = int32(bytes)
}
func (c *ConstatnFloatInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit32()
	c.value = math.Float32frombits(bytes)
}
func (c *ConstatnLongInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit64()
	c.value = int64(bytes)
}
func (c *ConstatnDoubleInfo) readInfo(reader *ClassReader) {
	bytes := reader.readUnit64()
	c.value = math.Float64frombits(bytes)
}
