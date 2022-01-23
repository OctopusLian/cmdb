package auth

import (
	"fmt"

	"github.com/astaxie/beego/context"

	"cmdb/models"
)

type Auth interface {
	name() string
	is(*context.Context) bool
	login(*AuthController) bool
	isLogin(*LoginRequiredController) *models.User
	logout(*AuthController) bool
	goLoginPage(*context.Context, string) bool
	goHomePage(*context.Context) bool
}

type Manager struct {
	plugins map[string]Auth
}

func NewManager() *Manager {
	return &Manager{
		plugins: map[string]Auth{},
	}
}

func (m *Manager) Register(a Auth) error {
	name := a.name()
	if name == defaultPlugin.name() {
		return nil
	}

	if _, ok := m.plugins[name]; ok {
		return fmt.Errorf("plugin %s is exists.", name)
	}
	m.plugins[name] = a
	return nil
}

func (m *Manager) GetPlugin(c *context.Context) Auth {
	for _, plugin := range m.plugins {
		if plugin.is(c) {
			return plugin
		}
	}
	return defaultPlugin
}

func (m *Manager) Login(c *AuthController) bool {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.login(c)
	}
	return defaultPlugin.login(c)
}

func (m *Manager) IsLogin(c *LoginRequiredController) *models.User {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.isLogin(c)
	}
	return defaultPlugin.isLogin(c)
}

func (m *Manager) Logout(c *AuthController) bool {
	if plugin := m.GetPlugin(c.Ctx); plugin != nil {
		return plugin.logout(c)
	}
	return defaultPlugin.logout(c)
}

func (m *Manager) GoHomePage(c *context.Context) bool {
	if plugin := m.GetPlugin(c); plugin != nil {
		return plugin.goHomePage(c)
	}
	return defaultPlugin.goHomePage(c)
}

func (m *Manager) GoLoginPage(c *context.Context, next string) bool {
	if plugin := m.GetPlugin(c); plugin != nil {
		return plugin.goLoginPage(c, next)
	}
	return defaultPlugin.goLoginPage(c, next)
}

var DefaultManager = NewManager()
