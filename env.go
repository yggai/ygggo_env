package ygggo_env

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// LoadEnv 自动查找并加载环境变量文件
// 从当前目录开始向上查找 .env 文件，找到后解析并设置环境变量
func LoadEnv() error {
	envFile, err := findEnvFile()
	if err != nil {
		return err
	}

	// 如果没有找到 .env 文件，不报错（这是正常情况）
	if envFile == "" {
		return nil
	}

	return loadEnvFile(envFile)
}

// findEnvFile 从当前目录开始向上查找 .env 文件
func findEnvFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get current directory: %w", err)
	}

	// 从当前目录开始向上查找
	for {
		envPath := filepath.Join(currentDir, ".env")

		// 检查文件是否存在
		if _, err := os.Stat(envPath); err == nil {
			return envPath, nil
		}

		// 获取父目录
		parentDir := filepath.Dir(currentDir)

		// 如果已经到达根目录，停止查找
		if parentDir == currentDir {
			break
		}

		currentDir = parentDir
	}

	// 没有找到 .env 文件
	return "", nil
}

// loadEnvFile 加载指定的环境变量文件
func loadEnvFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("failed to open env file %s: %w", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())

		// 跳过空行和注释行
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		// 解析键值对
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("invalid line %d in %s: %s", lineNum, filename, line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// 设置环境变量
		err := os.Setenv(key, value)
		if err != nil {
			return fmt.Errorf("failed to set environment variable %s: %w", key, err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading env file %s: %w", filename, err)
	}

	return nil
}

// GetStr 获取字符串类型的环境变量
// 如果环境变量不存在或为空，返回默认值
func GetStr(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// GetInt 获取整数类型的环境变量
// 如果环境变量不存在、为空或无法转换为整数，返回默认值
func GetInt(key string, defaultValue int) int {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	intValue, err := strconv.Atoi(value)
	if err != nil {
		return defaultValue
	}

	return intValue
}

// GetFloat 获取浮点数类型的环境变量
// 如果环境变量不存在、为空或无法转换为浮点数，返回默认值
func GetFloat(key string, defaultValue float64) float64 {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	floatValue, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return defaultValue
	}

	return floatValue
}

// GetBool 获取布尔类型的环境变量
// 支持多种布尔值表示：true/false, 1/0, yes/no, on/off (不区分大小写)
// 如果环境变量不存在、为空或无法识别为布尔值，返回默认值
func GetBool(key string, defaultValue bool) bool {
	value := strings.ToLower(strings.TrimSpace(os.Getenv(key)))
	if value == "" {
		return defaultValue
	}

	switch value {
	case "true", "1", "yes", "on":
		return true
	case "false", "0", "no", "off":
		return false
	default:
		return defaultValue
	}
}

// GetMap 获取字典类型的环境变量
// 环境变量值应该是有效的 JSON 格式
// 如果环境变量不存在、为空或无法解析为 JSON，返回默认值
func GetMap(key string, defaultValue map[string]interface{}) map[string]interface{} {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}

	var result map[string]interface{}
	err := json.Unmarshal([]byte(value), &result)
	if err != nil {
		return defaultValue
	}

	return result
}

// GetArr 获取数组类型的环境变量
// 支持两种格式：
// 1. 逗号分隔的字符串：value1,value2,value3
// 2. JSON 数组格式：["value1", "value2", "value3"]
// 如果环境变量不存在或为空，返回默认值
func GetArr(key string, defaultValue []string) []string {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return defaultValue
	}

	// 尝试解析为 JSON 数组
	if strings.HasPrefix(value, "[") {
		var result []string
		err := json.Unmarshal([]byte(value), &result)
		if err == nil {
			return result
		}
		// 如果 JSON 解析失败，返回默认值
		return defaultValue
	}

	// 按逗号分隔处理
	parts := strings.Split(value, ",")
	result := make([]string, len(parts))
	for i, part := range parts {
		result[i] = strings.TrimSpace(part)
	}

	return result
}
