package ygggo_env

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
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
