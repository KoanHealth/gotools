package docker_helper

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFix(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Docker Helper Suite")
}
