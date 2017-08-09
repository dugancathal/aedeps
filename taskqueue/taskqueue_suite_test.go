package taskqueue_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTaskqueue(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Taskqueue Suite")
}
