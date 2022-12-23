package main

import "fmt"
import "strings"
import "classpath"

func main() {
	// 解析命令行参数,存入cmd结构体
	cmd := parseCmd()
	Start(cmd)
}
func Start(cmd *Cmd) {
	if cmd.versionFlag {
		fmt.Println("version 0.0.1")
	} else if cmd.helpFlag || cmd.class == "" {
		// 如果指令是帮助(-? -help)或class参数为空,则输出提示
		printUsage()
	} else {
		// 启动jvm
		startJVM(cmd)
	}
}

// 启动jvm(现在还未实现)
func startJVM(cmd *Cmd) {
	// 解析-Xjre 和 cp或classpath
	fmt.Printf("%v\n", cmd.XjreOption)
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	fmt.Printf("classpath:%v class:%v args:%v\n",
		cp, cmd.class, cmd.args)
	// 把命令行参数传入的全类名的 . 替换为 /
	className := strings.Replace(cmd.class, ".", "/", -1)
	// 读取文件
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Printf("找不到或无法加载主函数 %s\n", cmd.class)
		return
	}
	fmt.Printf("class data:%v\n", classData)
}
