package classfile

import "fmt"

package classfile
import "fmt"
// 对应Java虚拟机规范定义的class文件格式
type ClassFile struct {
	//magic uint32
	minorVersion uint16       // 次版本号
	majorVersion uint16       // 主版本号
	constantPool ConstantPool // 常量池
	accessFlags  uint16       // 类访问标志
	thisClass    uint16       // 常量池索引->类名
	superClass   uint16       // 常量池索引->超类名
	interfaces   []uint16     // 接口索引表,存放的也是常量池索引，给出该类实现的所有接口的名字
	fields       []*MemberInfo // 字段表
	methods      []*MemberInfo // 方法表
	attributes   []AttributeInfo
}

// 把[]byte解析成ClassFile结构体
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		// Go语言没有异常处理机制，只有一个panic-recover机制
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
	// read方法依次调用其他方法解析class文件
	cf.read(cr)
	return
}
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)               // 见 3.2.3
	self.readAndCheckVersion(reader)             // 见 3.2.4
	self.constantPool = readConstantPool(reader) // 见 3.3
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool) // 见 3.2.8
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool) // 见 3.4
}

// MajorVersion等6个方法是Getter方法，把结构体的字段暴露给其他包使用
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
} // getter
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
} // getter
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
} // getter
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
} // getter
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
} // getter
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
} // getter
// ClassName 从常量池查找类名
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// SuperClassName 从常量池查找超(全)类名
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" // 只有 java.lang.Object 没有超类
}

// InterfaceNames 从常量池查找接口名
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

// readAndCheckMagic 检查是不是class文件
// 很多文件格式都会规定满足该格式的文件必须以某几个固定字节开头，这几个字节主要起标识作用，叫作魔数（magic number）
// class文件的魔数是“0xCAFEBABE”
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic(any("java.lang.ClassFormatError: magic! --魔数不正确"))
	}
}

// readAndCheckVersion 检查class文件版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		// 次版本号只在J2SE 1.2之前用过，从1.2开始基本上就没什么用了（都是0）
		if self.minorVersion == 0 {
			return
		}
	}
	panic(any("java.lang.UnsupportedClassVersionError! --不支持的class文件版本"))
}
