package service_framework

import "go.uber.org/zap"

func (m *Manager) Logger() *zap.SugaredLogger {
	return m.logger
}
