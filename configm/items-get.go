package configm

import "fmt"

func (m *Manager) ItemGetAll() map[string]Item {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.items
}

func (m *Manager) ItemGet(name string) (Item, error) {
	allItem := m.ItemGetAll()

	if item, ok := allItem[name]; ok {
		return item, nil
	}

	return Item{}, fmt.Errorf("未找到节点 %s", name)
}

func (m *Manager) ItemExists(name string) bool {
	allItem := m.ItemGetAll()

	if _, ok := allItem[name]; ok {
		return true
	}

	return false
}

func (m *Manager) ItemGetAllName() []string {
	allItem := m.ItemGetAll()

	keys := make([]string, 0, len(allItem))
	for k := range allItem {
		keys = append(keys, k)
	}

	return keys
}
