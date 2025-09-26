package configm

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func getConfigPath() string {
	configPath, err := os.UserConfigDir()
	if err != nil {
		panic(err)
	}

	dirPath := filepath.Join(configPath, "nodebus")
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		if err := os.MkdirAll(dirPath, 0750); err != nil {
			panic(err)
		}
	}

	return filepath.Join(dirPath, "nodebus.json")
}

func (m *Manager) LoadJSON() error {
	configPath := getConfigPath()

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return fmt.Errorf("无法加载 json 配置: %v", err)
	}

	f, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("无法打开 json 配置: %v", err)
	}
	defer f.Close()

	var items map[string]Item
	if err := json.NewDecoder(f).Decode(&items); err != nil {
		return fmt.Errorf("无法解析 json 配置: %v", err)
	}

	m.mu.Lock()
	defer m.mu.Unlock()

	m.items = items
	return nil
}

func (m *Manager) SaveJSON() error {
	configPath := getConfigPath()

	f, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("无法创建 json 配置: %v", err)
	}
	defer f.Close()

	m.mu.RLock()
	defer m.mu.RUnlock()

	if err := json.NewEncoder(f).Encode(m.items); err != nil {
		return fmt.Errorf("无法保存 json 配置: %v", err)
	}

	return nil
}
