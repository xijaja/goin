package start

import (
	"flag"
	"fmt"
)

// ---------------------------------------------
// 启动配置
// ---------------------------------------------

// S P T 启动变量
var S = flag.Bool("s", false, "true 为生产环境，默认 false 调试环境")
var P = flag.Bool("p", false, "true 为启用多线程，默认 false 不启动")
var T = flag.Bool("t", false, "true 为启动定时任务，默认 false 不启动")

// Config 初始化配置
var Config = projectConfig{}

// 初始化配置信息
func init() {
	// 解析命令行参数
	flag.Parse()

	// 执行编译前端的命令
	// if Config.NeedBuild {
	// 	switch Config.PkgManager {
	// 	case "npm", "pnpm", "yarn", "cnpm":
	// 		cmdrun.RunCmd(fmt.Sprintf("cd ./frontend && %s run build", Config.PkgManager))
	// 	default:
	// 		log.Fatal("请在 .env 或 .env.dev 文件中设置 PKG_MANAGER 的值\n例如: PKG_MANAGER = npm")
	// 	}
	// }

	// 设置为发布模式
	if *S {
		Config = Config.getMyConfig(true) // 赋值为生产环境配置
		fmt.Printf("当前为🔥生产环境🔥 定时任务启动状态:%v\n", *T)
	} else {
		Config = Config.getMyConfig(false) // 赋值为开发环境配置
		fmt.Printf("当前为🌲开发环境🌲 定时任务启动状态:%v\n", *T)
	}
}
