package beerJSON

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestWater(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Water Suite")
}