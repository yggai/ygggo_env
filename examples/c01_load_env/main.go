package main

import (
	"fmt"
	"os"

	gge "github.com/yggai/ygggo_env"
)

func main() {
	// 自动查找并加载环境变量
	gge.LoadEnv()

	// 获取环境变量的值
	fmt.Println(os.Getenv("YGGGO_MYSQL_HOST"))
	fmt.Println(os.Getenv("YGGGO_MYSQL_PORT"))
	fmt.Println(os.Getenv("YGGGO_MYSQL_USERNAME"))
	fmt.Println(os.Getenv("YGGGO_MYSQL_PASSWORD"))
	fmt.Println(os.Getenv("YGGGO_MYSQL_DATABASE"))
}
