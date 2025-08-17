package ygggo_env

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadEnv_FileExists(t *testing.T) {
	// 创建临时目录和 .env 文件
	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")
	
	// 创建测试用的 .env 文件
	envContent := `YGGGO_MYSQL_HOST=localhost
YGGGO_MYSQL_PORT=3306
YGGGO_MYSQL_USERNAME=root
YGGGO_MYSQL_PASSWORD=password123
YGGGO_MYSQL_DATABASE=testdb
`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test .env file: %v", err)
	}
	
	// 保存当前工作目录
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	
	// 切换到临时目录
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	
	// 清理环境变量
	testVars := []string{
		"YGGGO_MYSQL_HOST",
		"YGGGO_MYSQL_PORT", 
		"YGGGO_MYSQL_USERNAME",
		"YGGGO_MYSQL_PASSWORD",
		"YGGGO_MYSQL_DATABASE",
	}
	for _, v := range testVars {
		os.Unsetenv(v)
	}
	
	// 调用 LoadEnv
	err = LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv() failed: %v", err)
	}
	
	// 验证环境变量是否正确设置
	expected := map[string]string{
		"YGGGO_MYSQL_HOST":     "localhost",
		"YGGGO_MYSQL_PORT":     "3306",
		"YGGGO_MYSQL_USERNAME": "root",
		"YGGGO_MYSQL_PASSWORD": "password123",
		"YGGGO_MYSQL_DATABASE": "testdb",
	}
	
	for key, expectedValue := range expected {
		actualValue := os.Getenv(key)
		if actualValue != expectedValue {
			t.Errorf("Expected %s=%s, got %s=%s", key, expectedValue, key, actualValue)
		}
	}
}

func TestLoadEnv_FileNotExists(t *testing.T) {
	// 创建临时目录（不包含 .env 文件）
	tempDir := t.TempDir()
	
	// 保存当前工作目录
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	
	// 切换到临时目录
	err := os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	
	// 调用 LoadEnv，应该不报错（文件不存在是正常情况）
	err = LoadEnv()
	if err != nil {
		t.Errorf("LoadEnv() should not fail when .env file doesn't exist, got: %v", err)
	}
}

func TestLoadEnv_SearchParentDirectories(t *testing.T) {
	// 创建临时目录结构
	tempDir := t.TempDir()
	subDir := filepath.Join(tempDir, "subdir")
	err := os.Mkdir(subDir, 0755)
	if err != nil {
		t.Fatalf("Failed to create subdirectory: %v", err)
	}
	
	// 在父目录创建 .env 文件
	envFile := filepath.Join(tempDir, ".env")
	envContent := `YGGGO_TEST_VAR=parent_value`
	err = os.WriteFile(envFile, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test .env file: %v", err)
	}
	
	// 保存当前工作目录
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	
	// 切换到子目录
	err = os.Chdir(subDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	
	// 清理环境变量
	os.Unsetenv("YGGGO_TEST_VAR")
	
	// 调用 LoadEnv
	err = LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv() failed: %v", err)
	}
	
	// 验证环境变量是否正确设置
	actualValue := os.Getenv("YGGGO_TEST_VAR")
	expectedValue := "parent_value"
	if actualValue != expectedValue {
		t.Errorf("Expected YGGGO_TEST_VAR=%s, got %s", expectedValue, actualValue)
	}
}

func TestLoadEnv_IgnoreComments(t *testing.T) {
	// 创建临时目录和 .env 文件
	tempDir := t.TempDir()
	envFile := filepath.Join(tempDir, ".env")
	
	// 创建包含注释的 .env 文件
	envContent := `# This is a comment
YGGGO_VAR1=value1
# Another comment
YGGGO_VAR2=value2

# Empty line above and below

YGGGO_VAR3=value3
`
	err := os.WriteFile(envFile, []byte(envContent), 0644)
	if err != nil {
		t.Fatalf("Failed to create test .env file: %v", err)
	}
	
	// 保存当前工作目录
	originalWd, _ := os.Getwd()
	defer os.Chdir(originalWd)
	
	// 切换到临时目录
	err = os.Chdir(tempDir)
	if err != nil {
		t.Fatalf("Failed to change directory: %v", err)
	}
	
	// 清理环境变量
	testVars := []string{"YGGGO_VAR1", "YGGGO_VAR2", "YGGGO_VAR3"}
	for _, v := range testVars {
		os.Unsetenv(v)
	}
	
	// 调用 LoadEnv
	err = LoadEnv()
	if err != nil {
		t.Fatalf("LoadEnv() failed: %v", err)
	}
	
	// 验证环境变量是否正确设置
	expected := map[string]string{
		"YGGGO_VAR1": "value1",
		"YGGGO_VAR2": "value2",
		"YGGGO_VAR3": "value3",
	}
	
	for key, expectedValue := range expected {
		actualValue := os.Getenv(key)
		if actualValue != expectedValue {
			t.Errorf("Expected %s=%s, got %s=%s", key, expectedValue, key, actualValue)
		}
	}
}
