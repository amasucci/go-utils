package utils

import (
	"errors"
	"log"
	"os"
)

//GetEnvVar returns a string containing the env var value and an error if the var is required but not found
func GetEnvVar(varName string, required bool) (string, error) {
	env := os.Getenv(varName)
	if !required || len(env) > 0 {
		log.Println("Using environment var for " + varName)
		return env, nil
	}

	return env, errors.New("Variable is required")
}

//GetEnvVarOrDefault returns a string containing the env var value and if empty returns the defaultValue
func GetEnvVarOrDefault(varName string, defaultValue string) string {
	env := os.Getenv(varName)
	if len(env) > 0 {
		log.Println("Using environment var for " + varName)
		return env
	}
	log.Println("Using default value: " + defaultValue + " for " + varName)
	return defaultValue
}
