package config

import (
	"os"
	"strconv"
)

func readConfigString(key string, defaultValue string) string {

	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	return val
}

func readConfigInt(key string, defaultValue int) int {

	val := os.Getenv(key)
	if val == "" {
		return defaultValue
	}

	value, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}

	return value
}
