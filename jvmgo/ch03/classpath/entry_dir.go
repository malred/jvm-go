package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 目录形式的类路径
type DirEntry struct {
	absDir string // 存放目录的绝对路径
}

// 构造函数
func newDirEntry(path string) *DirEntry {
	// path转换成绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(any(err))
	}
	return &DirEntry{absDir}
}

// 读取文件
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	// 把当absDir和class文件名拼成一个完整的路径
	fileName := filepath.Join(self.absDir, className)
	// 读取文件
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 返回path绝对路径
func (self *DirEntry) String() string {
	return self.absDir
}
