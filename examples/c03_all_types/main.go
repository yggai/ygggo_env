package main

import (
	"fmt"

	gge "github.com/yggai/ygggo_env"
)

func main() {
	// 自动查找并加载环境变量
	gge.LoadEnv()

	fmt.Println("=== 字符串类型 ===")
	fmt.Printf("YGGGO_MYSQL_HOST: %s\n", gge.GetStr("YGGGO_MYSQL_HOST", "localhost"))
	fmt.Printf("YGGGO_MYSQL_USERNAME: %s\n", gge.GetStr("YGGGO_MYSQL_USERNAME", "root"))
	fmt.Printf("NON_EXISTING_STR: %s\n", gge.GetStr("NON_EXISTING_STR", "default_value"))

	fmt.Println("\n=== 整数类型 ===")
	fmt.Printf("YGGGO_MYSQL_PORT: %d\n", gge.GetInt("YGGGO_MYSQL_PORT", 3306))
	fmt.Printf("NON_EXISTING_INT: %d\n", gge.GetInt("NON_EXISTING_INT", 8080))

	fmt.Println("\n=== 浮点数类型 ===")
	fmt.Printf("YGGGO_TIMEOUT: %.1f\n", gge.GetFloat("YGGGO_TIMEOUT", 10.0))
	fmt.Printf("NON_EXISTING_FLOAT: %.2f\n", gge.GetFloat("NON_EXISTING_FLOAT", 3.14))

	fmt.Println("\n=== 布尔类型 ===")
	fmt.Printf("YGGGO_DEBUG: %t\n", gge.GetBool("YGGGO_DEBUG", false))
	fmt.Printf("NON_EXISTING_BOOL: %t\n", gge.GetBool("NON_EXISTING_BOOL", true))

	fmt.Println("\n=== 数组类型 ===")
	servers := gge.GetArr("YGGGO_SERVERS", []string{"default_server"})
	fmt.Printf("YGGGO_SERVERS: %v\n", servers)
	
	features := gge.GetArr("YGGGO_FEATURES", []string{"default_feature"})
	fmt.Printf("YGGGO_FEATURES: %v\n", features)
	
	nonExistingArr := gge.GetArr("NON_EXISTING_ARR", []string{"default1", "default2"})
	fmt.Printf("NON_EXISTING_ARR: %v\n", nonExistingArr)

	fmt.Println("\n=== 字典类型 ===")
	config := gge.GetMap("YGGGO_CONFIG", map[string]interface{}{"default": "value"})
	fmt.Printf("YGGGO_CONFIG: %v\n", config)
	
	nonExistingMap := gge.GetMap("NON_EXISTING_MAP", map[string]interface{}{"key": "default"})
	fmt.Printf("NON_EXISTING_MAP: %v\n", nonExistingMap)
}
