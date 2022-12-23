package classfile

import "encoding/binary"

// 解析class的第一步是读取,这里利用结构体来帮助读取数据
// ClassReader []byte类型的包装
type ClassReader struct {
	data []byte
}

// readUint8 读取u1类型数据
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// readUint16 读取u2类型数据
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data) // 读取
	self.data = self.data[2:]                 // 读完就丢弃
	return val
}

// readUint32 读取u4类型数据
func (self *ClassReader) readUint32() uint32 {
	// encoding/binary包中定义了一个变量BigEndian,可以从[]byte中解码多字节数据
	val := binary.BigEndian.Uint32(self.data) // 读取
	self.data = self.data[4:]                 // 读完就丢弃
	return val
}

// readUint64 读取uint64（Java虚拟机规范并没有定义u8）类型数据
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data) // 读取
	self.data = self.data[8:]                 // 读完就丢弃
	return val
}

// readUint16s 读取uint16表，表的大小由开头的uint16数据指出
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() // 读取uint16类型和长度的数据
	s := make([]uint16, n)
	for i := range s {
		// 把读到的数据放入uint16切片(数组)中
		s[i] = self.readUint16()
	}
	return s
}

// readBytes 从当前结构体中,读取指定数量的字节
func (self *ClassReader) readBytes(n uint32) []byte {
	bytes := self.data[:n]    // 读取
	self.data = self.data[n:] // 读完就丢弃
	return bytes
}
