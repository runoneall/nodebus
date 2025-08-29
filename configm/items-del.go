package configm

func (m *Manager) ItemDel(
	names []string,
	success_call func(name string),
	failed_call func(name string),
) {
	m.mu.RLock()
	items := m.items
	m.mu.RUnlock()

	for _, name := range names {
		if _, ok := items[name]; !ok {
			failed_call(name)
		} else {
			delete(items, name)
			success_call(name)
		}
	}

	m.mu.Lock()
	m.items = items
	m.mu.Unlock()
}
