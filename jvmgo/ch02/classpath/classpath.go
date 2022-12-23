package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry // 启动类路径
	extClasspath  Entry // 拓展类路径
	userClasspath Entry // 用户类路径
}

// 解析命令行
func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	// -Xjre选项解析启动类路径和扩展类路径
	cp.parseBootAndExtClasspath(jreOption)
	// -classpath/-cp选项解析用户类路径
	cp.parseUserClasspath(cpOption)
	return cp
}

// 解析启动类路径和拓展类路径
func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	// 获取jre路径,用于后面new entry
	jreDir := getJreDir(jreOption)
	// jre/lib/*
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	// 启动类路径entry
	self.bootClasspath = newWildcardEntry(jreLibPath)
	//fmt.Println(self.bootClasspath) // ?
	// jre/lib/ext/*
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	// 拓展类路径entry
	self.extClasspath = newWildcardEntry(jreExtPath)
}

// 获取jre路径
func getJreDir(jreOption string) string {
	// 优先使用用户输入的-Xjre选项作为jre目录
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	// 如果没有输入该选项，则在当前目录下寻找jre目录
	if exists("./jre") {
		return "./jre"
	}
	// 如果找不到，尝试使用JAVA_HOME环境变量
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic(any("找不到jre文件夹!"))
}

// 判断目录是否存在
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// 解析用户类路径
func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" { // 如果没有提供 -classpath或-cp选项,就默认当前目录
		cpOption = "."
	}
	// 用户类路径entry
	self.userClasspath = newEntry(cpOption)
}

// 依次从启动类路径、扩展类路径和用户类路径中搜索class文件
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class" // 添加后缀
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClasspath.readClass(className)
}

// 返回用户类路径
func (self *Classpath) String() string {
	return self.userClasspath.String()
}
