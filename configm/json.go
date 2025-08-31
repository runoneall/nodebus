package configm

import (
	"encoding/json"
	"fmt"
	"os"
)

func (m *Manager) LoadJSON() error {
	if _, err := os.Stat("nodebus.json"); os.IsNotExist(err) {
		return fmt.Errorf("无法加载 json 配置: %v", err)
	}

	f, err := os.Open("nodebus.json")
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
	f, err := os.Create("nodebus.json")
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
