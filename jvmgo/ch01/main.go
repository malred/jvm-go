package main

import "fmt"

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
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
