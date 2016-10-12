package test

import "golang.org/x/net/context"

type MockLogger struct{}

func (l *MockLogger) Infof(ctx context.Context, format string, args ...interface{}) {
}

func (l *MockLogger) Errorf(ctx context.Context, format string, args ...interface{}) {
}
