package api

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestConverterService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "api")
}
