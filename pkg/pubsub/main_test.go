package pubsub

import (
	"os"
	"testing"
)

var testService *Service

func TestMain(m *testing.M) {
	testService = New(
		"localhost:6379",
		"",
	)

	os.Exit(m.Run())
}
