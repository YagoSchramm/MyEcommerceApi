package foundation

import "go.uber.org/zap"

func NewLogger() (*zap.Logger, error) {
	// on development for now with a change for production
	return zap.NewDevelopment()
}
