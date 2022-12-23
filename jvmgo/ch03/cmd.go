package main

import (
	"flag"
	"fmt"
	"os"
)

//Go源文件一般以.go作为后缀，文件名全部小写，多个单词之间用下划线分隔。Go语言规范
//要求Go源文件必须使用UTF-8编码
// Cmd 命令行相关的结构体
type Cmd struct {
	helpFlag    bool     // 帮助
	versionFlag bool     // 版本
	cpOption    string   // 类路径
	XjreOption  string   // 指定(Java标准库中的类)jre目录的所在路径
	class       string   // 应该是类字节码?
	args        []string // 接收的参数
}

// 解析命令行参数
func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printUsage // Usage打印一条Usage消息，记录所有已定义的命令行标志
	// 设置需要解析的选项   flag.TypeVar(*变量,flag 名, 默认值, 帮助信息) *Type
	// -?或-help对应cmd的helpFlag
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	// -version对应cmd的versionFlag
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	// -cp或-classpath对应cmd的cpOption
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	// -Xjre指定java标准库jre所在目录(比如 你的javahome路径 + \jre
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	// 开始解析
	flag.Parse()
	// flag.Args()函数可以捕获其他没有被解析的参数
	args := flag.Args()
	if len(args) > 0 {
		// 其中第一个参数(用户传递的)就是主类名
		cmd.class = args[0]
		// 剩下的是要传递给主类的参数
		cmd.args = args[1:]
	}
	// 解析完放入cmd结构体,返回给外部使用
	return cmd
}

// 出错时打印提示信息
func printUsage() {
	//fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
	fmt.Printf("标准格式: %s [-options] class [args...]\n", os.Args[0])
	fmt.Printf(" [-options]是可选参数,要加 - ,如果没传,则默认是class参数,即类路径,[args]是传个class的参数\n")
}
