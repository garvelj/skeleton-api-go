package utils

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockResponder *Client

func TestNewResponder(t *testing.T) {
	mockResponder = NewClient(
		true,
		true,
		log.Printf,
		alertSimulate,
		"json",
		GinResponder{},
	)

	assert.NotNil(t, mockResponder)
	assert.Equal(t, mockResponder.shouldLog, true)
	assert.Equal(t, mockResponder.shouldAlert, true)
	assert.NotNil(t, mockResponder.logf)
	assert.NotNil(t, mockResponder.alertf)
}

func TestResponderLog(t *testing.T) {
	if mockResponder == nil {
		TestNewResponder(t)
	}
	mockResponder.Log("This %s should be visible in the terminal", "argument")
}

func TestResponderAlert(t *testing.T) {
	if mockResponder == nil {
		TestNewResponder(t)
	}
	mockResponder.Alert("This is %s that should be alerted!", "a MESSAGE")
}

func alertSimulate(message string, args ...any) {
	message = fmt.Sprintf(message, args...)
	fmt.Println("\n!ALERT! ", message, " ⚠️")
}
