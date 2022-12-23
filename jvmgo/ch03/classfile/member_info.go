package classfile

// 统一表示字段和方法
type MemberInfo struct {
	cp              ConstantPool // 保存常量池指针
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

// readMembers 读取字段表或方法表
func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16() // 得到字段表或方法表的长度
	members := make([]*MemberInfo, memberCount)
	for i := range members {
		members[i] = readMember(reader, cp) // 读取字段或方法
	}
	return members
}

// readMember 函数读取字段或方法数据
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp), // 见 3.4
	}
}
func (self *MemberInfo) AccessFlags() uint16 { ... } // getter
func (self *MemberInfo) Name() string        { ... }
func (self *MemberInfo) Descriptor() string  { ... }
