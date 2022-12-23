package classpath

import (
	"os"
	"strings"
)

// 路径分隔符,windows是 ;类UNIX(包括Linux、Mac OS X等)是 :
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 参数是class文件的相对路径
	// 路径之间用斜线（/）分隔
	// 文件名有.class后缀
	// 返回值是读取到的字节数据、最终定位到class文件的Entry,以及错误信息
	readClass(className string) ([]byte, Entry, error) // 寻找和加载class文件的方法
	String() string                                    // 类似java里的toString方法
}

// 根据参数创建不同类型的Entry实例
func newEntry(path string) Entry {
	// Go结构体不需要显示实现接口，只要方法匹配即可。
	// Go没有专门的构造函数,统一使用new开头的函数来创建结构体实例,并把这类函数称为构造函数
	if strings.Contains(path, pathListSeparator) {
		// 目录形式的类路径
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		// zip或jar形式的类路径
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
