package ygggo_env

import (
	"os"
	"testing"
)

func TestGetStr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue string
		envValue     string
		expected     string
	}{
		{
			name:         "existing environment variable",
			key:          "TEST_STR",
			defaultValue: "default",
			envValue:     "test_value",
			expected:     "test_value",
		},
		{
			name:         "non-existing environment variable",
			key:          "NON_EXISTING_STR",
			defaultValue: "default_value",
			envValue:     "",
			expected:     "default_value",
		},
		{
			name:         "empty environment variable",
			key:          "EMPTY_STR",
			defaultValue: "default",
			envValue:     "",
			expected:     "default",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetStr(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetStr(%s, %s) = %s, want %s", tt.key, tt.defaultValue, result, tt.expected)
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetInt(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue int
		envValue     string
		expected     int
	}{
		{
			name:         "valid integer",
			key:          "TEST_INT",
			defaultValue: 0,
			envValue:     "123",
			expected:     123,
		},
		{
			name:         "negative integer",
			key:          "TEST_NEG_INT",
			defaultValue: 0,
			envValue:     "-456",
			expected:     -456,
		},
		{
			name:         "invalid integer",
			key:          "TEST_INVALID_INT",
			defaultValue: 999,
			envValue:     "not_a_number",
			expected:     999,
		},
		{
			name:         "non-existing variable",
			key:          "NON_EXISTING_INT",
			defaultValue: 42,
			envValue:     "",
			expected:     42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetInt(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetInt(%s, %d) = %d, want %d", tt.key, tt.defaultValue, result, tt.expected)
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetFloat(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue float64
		envValue     string
		expected     float64
	}{
		{
			name:         "valid float",
			key:          "TEST_FLOAT",
			defaultValue: 0.0,
			envValue:     "123.45",
			expected:     123.45,
		},
		{
			name:         "negative float",
			key:          "TEST_NEG_FLOAT",
			defaultValue: 0.0,
			envValue:     "-67.89",
			expected:     -67.89,
		},
		{
			name:         "integer as float",
			key:          "TEST_INT_FLOAT",
			defaultValue: 0.0,
			envValue:     "100",
			expected:     100.0,
		},
		{
			name:         "invalid float",
			key:          "TEST_INVALID_FLOAT",
			defaultValue: 99.9,
			envValue:     "not_a_float",
			expected:     99.9,
		},
		{
			name:         "non-existing variable",
			key:          "NON_EXISTING_FLOAT",
			defaultValue: 3.14,
			envValue:     "",
			expected:     3.14,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetFloat(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetFloat(%s, %f) = %f, want %f", tt.key, tt.defaultValue, result, tt.expected)
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetBool(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue bool
		envValue     string
		expected     bool
	}{
		{
			name:         "true value",
			key:          "TEST_BOOL_TRUE",
			defaultValue: false,
			envValue:     "true",
			expected:     true,
		},
		{
			name:         "false value",
			key:          "TEST_BOOL_FALSE",
			defaultValue: true,
			envValue:     "false",
			expected:     false,
		},
		{
			name:         "1 as true",
			key:          "TEST_BOOL_1",
			defaultValue: false,
			envValue:     "1",
			expected:     true,
		},
		{
			name:         "0 as false",
			key:          "TEST_BOOL_0",
			defaultValue: true,
			envValue:     "0",
			expected:     false,
		},
		{
			name:         "yes as true",
			key:          "TEST_BOOL_YES",
			defaultValue: false,
			envValue:     "yes",
			expected:     true,
		},
		{
			name:         "no as false",
			key:          "TEST_BOOL_NO",
			defaultValue: true,
			envValue:     "no",
			expected:     false,
		},
		{
			name:         "invalid bool",
			key:          "TEST_INVALID_BOOL",
			defaultValue: true,
			envValue:     "maybe",
			expected:     true,
		},
		{
			name:         "non-existing variable",
			key:          "NON_EXISTING_BOOL",
			defaultValue: false,
			envValue:     "",
			expected:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetBool(tt.key, tt.defaultValue)
			if result != tt.expected {
				t.Errorf("GetBool(%s, %t) = %t, want %t", tt.key, tt.defaultValue, result, tt.expected)
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetMap(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue map[string]interface{}
		envValue     string
		expected     map[string]interface{}
	}{
		{
			name:         "valid JSON map",
			key:          "TEST_MAP",
			defaultValue: map[string]interface{}{"default": "value"},
			envValue:     `{"host": "localhost", "port": 3306}`,
			expected:     map[string]interface{}{"host": "localhost", "port": float64(3306)},
		},
		{
			name:         "invalid JSON",
			key:          "TEST_INVALID_MAP",
			defaultValue: map[string]interface{}{"default": "value"},
			envValue:     `{invalid json}`,
			expected:     map[string]interface{}{"default": "value"},
		},
		{
			name:         "non-existing variable",
			key:          "NON_EXISTING_MAP",
			defaultValue: map[string]interface{}{"key": "default"},
			envValue:     "",
			expected:     map[string]interface{}{"key": "default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetMap(tt.key, tt.defaultValue)

			// 比较 map 内容
			if len(result) != len(tt.expected) {
				t.Errorf("GetMap(%s) length = %d, want %d", tt.key, len(result), len(tt.expected))
				return
			}

			for k, v := range tt.expected {
				if result[k] != v {
					t.Errorf("GetMap(%s)[%s] = %v, want %v", tt.key, k, result[k], v)
				}
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}

func TestGetArr(t *testing.T) {
	tests := []struct {
		name         string
		key          string
		defaultValue []string
		envValue     string
		expected     []string
	}{
		{
			name:         "comma separated values",
			key:          "TEST_ARR_COMMA",
			defaultValue: []string{"default"},
			envValue:     "value1,value2,value3",
			expected:     []string{"value1", "value2", "value3"},
		},
		{
			name:         "JSON array",
			key:          "TEST_ARR_JSON",
			defaultValue: []string{"default"},
			envValue:     `["item1", "item2", "item3"]`,
			expected:     []string{"item1", "item2", "item3"},
		},
		{
			name:         "single value",
			key:          "TEST_ARR_SINGLE",
			defaultValue: []string{"default"},
			envValue:     "single_value",
			expected:     []string{"single_value"},
		},
		{
			name:         "empty value",
			key:          "TEST_ARR_EMPTY",
			defaultValue: []string{"default"},
			envValue:     "",
			expected:     []string{"default"},
		},
		{
			name:         "invalid JSON array",
			key:          "TEST_ARR_INVALID",
			defaultValue: []string{"default"},
			envValue:     `[invalid json`,
			expected:     []string{"default"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 清理环境变量
			os.Unsetenv(tt.key)

			// 设置环境变量（如果有值）
			if tt.envValue != "" {
				os.Setenv(tt.key, tt.envValue)
			}

			result := GetArr(tt.key, tt.defaultValue)

			// 比较数组内容
			if len(result) != len(tt.expected) {
				t.Errorf("GetArr(%s) length = %d, want %d", tt.key, len(result), len(tt.expected))
				return
			}

			for i, v := range tt.expected {
				if result[i] != v {
					t.Errorf("GetArr(%s)[%d] = %s, want %s", tt.key, i, result[i], v)
				}
			}

			// 清理
			os.Unsetenv(tt.key)
		})
	}
}
