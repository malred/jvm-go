package classpath

import (
	"errors"
	"strings"
)

/**
* CompositeEntry由更小的Entry组成
* 在Go语言中，数组属于比较低层的数据结构，很少直接使用。大部分情况下，使用更便利的slice类型
 */
type CompositeEntry []Entry

// 构造函数
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	// 把参数（路径列表）按分隔符分成小路径
	for _, path := range strings.Split(pathList, pathListSeparator) {
		// 把每个小路径都转换成具体的Entry实例
		entry := newEntry(path) // entry根据路径类型寻找对应实现并实例化
		// 拼接到compositeEntry
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

// 读取文件
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	// 遍历读取文件内容
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

//用每一个子路径的String（）方法，然后把得到的字符串用路径分隔符拼接起来即可
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
