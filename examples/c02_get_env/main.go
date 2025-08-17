package main

import (
	"fmt"

	gge "github.com/yggai/ygggo_env"
)

func main() {
	// 自动查找并加载环境变量
	gge.LoadEnv()

	// 获取环境变量的值
	fmt.Println(gge.GetStr("YGGGO_MYSQL_HOST", "localhost"))
	fmt.Println(gge.GetInt("YGGGO_MYSQL_PORT", 3306))
	fmt.Println(gge.GetStr("YGGGO_MYSQL_USERNAME", "root"))
	fmt.Println(gge.GetStr("YGGGO_MYSQL_PASSWORD", ""))
	fmt.Println(gge.GetStr("YGGGO_MYSQL_DATABASE", "test"))
}
