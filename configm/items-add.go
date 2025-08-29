package configm

func (m *Manager) ItemAdd(item Item) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.items[item.Name] = item
}
