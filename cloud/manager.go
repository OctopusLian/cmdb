package cloud

import "fmt"

type Manager struct {
	plugins map[string]Cloud
}

func NewManager() *Manager {
	return &Manager{
		plugins: map[string]Cloud{},
	}
}

func (m *Manager) Register(c Cloud) error {
	name := c.Name()
	if _, ok := m.plugins[name]; ok {
		return fmt.Errorf("plugin %s is exists.", name)
	}
	m.plugins[name] = c
	return nil
}

func (m *Manager) Cloud(name string) (Cloud, error) {
	if c, ok := m.plugins[name]; ok {
		return c, nil
	}
	return nil, fmt.Errorf("cloud %s is not found", name)
}

var DefaultManager = NewManager()
