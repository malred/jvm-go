package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// ZIP或JAR文件形式的类路径
type ZipEntry struct {
	absPath string // 绝对路径
}

// 读取jar或zip文件
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开目录
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	// 对于该目录下的文件 for
	for _, f := range r.File {
		// 如果文件名和传入的className匹配
		if f.Name == className {
			// 打开文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			// 读取文件内容
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("未找到class文件: " + className)
}

// 构造函数
func newZipEntry(path string) *ZipEntry {
	// path转绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(any(err))
	}
	return &ZipEntry{absPath}
}

// 返回绝对路径
func (self *ZipEntry) String() string {
	return self.absPath
}
