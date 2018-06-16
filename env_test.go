package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVarReturnsErrorWhenRequiredAndNotPresent(t *testing.T) {
	expectedErrorString := "Variable is required"
	actualObj, err := GetEnvVar("NOT_SET_VAR", true)
	assert.EqualError(t, err, expectedErrorString)
	assert.Equal(t, "", actualObj)
}

func TestGetEnvVarReturnsValueWhenRequiredAndPresent(t *testing.T) {
	expectedValue := "VAR_VALUE"
	os.Setenv("VAR_NAME", expectedValue)
	actualObj, err := GetEnvVar("VAR_NAME", true)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedValue, actualObj)
	os.Unsetenv("VAR_NAME")
}

func TestGetEnvVarReturnsValueWhenNotRequiredAndPresent(t *testing.T) {
	expectedValue := "VAR_VALUE"
	os.Setenv("VAR_NAME", expectedValue)
	actualObj, err := GetEnvVar("VAR_NAME", false)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedValue, actualObj)
	os.Unsetenv("VAR_NAME")
}

func TestGetEnvVarReturnsEmptyWhenNotRequiredAndPresent(t *testing.T) {
	expectedValue := ""
	actualObj, err := GetEnvVar("VAR_NAME", false)
	assert.Equal(t, err, nil)
	assert.Equal(t, expectedValue, actualObj)
	os.Unsetenv("VAR_NAME")
}

func TestGetEnvVarOrDefaultReturnsDefaultWhenValueNotPresent(t *testing.T) {
	expectedValue := "dodo"
	actualObj := GetEnvVarOrDefault("VAR_NAME", expectedValue)
	assert.Equal(t, expectedValue, actualObj)
}

func TestGetEnvVarOrDefaultReturnsExpectedWhenValueIsPresent(t *testing.T) {
	expectedValue := "dodo"
	defaultValue := "not_considered"
	os.Setenv("VAR_NAME", expectedValue)
	actualObj := GetEnvVarOrDefault("VAR_NAME", defaultValue)
	assert.Equal(t, expectedValue, actualObj)
	os.Unsetenv("VAR_NAME")
}
