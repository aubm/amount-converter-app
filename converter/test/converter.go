package test

import (
	"github.com/aubm/amount-converter-app/converter"
	"github.com/stretchr/testify/mock"
)

type MockConverterService struct {
	mock.Mock
}

func (m *MockConverterService) Convert(amount float64) map[string]float64 {
	args := m.Called(amount)
	return args.Get(0).(map[string]float64)
}

func (m *MockConverterService) RatesDef() *converter.RatesDef {
	args := m.Called()
	return args.Get(0).(*converter.RatesDef)
}

func (m *MockConverterService) FetchConfiguration(url string) error {
	args := m.Called(url)
	return args.Error(0)
}
