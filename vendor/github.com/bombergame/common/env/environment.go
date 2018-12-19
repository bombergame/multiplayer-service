package env

import (
	"os"
)

//GetVar returns environment variable value
func GetVar(name, defaultValue string) string {
	v := os.Getenv(name)
	if v == "" {
		v = defaultValue
	}
	return v
}

//SetVar sets environment variable value and returns previous value
func SetVar(name, value string) string {
	v := os.Getenv(name)
	if err := os.Setenv(name, value); err != nil {
		panic(err)
	}
	return v
}
