package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("FOO", "")

	fmt.Println(getEnv("FOO", WithValueIfUnset("hello")))       // -
	fmt.Println(getEnv("FOO", WithValueIfEmpty("hello")))       // :-
	fmt.Println(getEnv("FOO", WithMustBeSet()))                 // ?
	fmt.Println(getEnv("FOO", WithMustBeNonEmpty()))            // :?
	fmt.Println(getEnv("FOO", WithValueIfSet("hello")))         // +
	fmt.Println(getEnv("FOO", WithValueIfNonEmpty("hello")))    // :+
	fmt.Println(getEnv("FOO", WithUpdateValueIfUnset("hello"))) // =
	fmt.Println(getEnv("FOO", WithUpdateValueIfEmpty("hello"))) // :=
}

func getEnv(key string, options ...EnvVarOptionFn) string {
	envCtx := envVarContext{
		key: key,
	}

	return envCtx.Get(options...)
}

// EnvVar provides a way to get environment variables with optional conditions
// and transformations, such as default values when the upstream is unset or
// empty. It also allows to update the value.
type EnvVar interface {
	// Get retrieves the environment variable value and returns a result based on
	// the options rules. Options run in a FIFO order
	Get(options ...EnvVarOptionFn) string

	// Set updates the environment variable, i.e. it changes the OS-provided
	// process memory area, so future calls to os.Getenv returns the new value
	Set(value string)
}

// EnvVarContext represents the state of an environment variable and allows
// overriding the local value e.g. to provide a default one.
type EnvVarContext interface {
	EnvVar

	// Key returns the environment variable name
	Key() string

	// Value returns the environment variable value
	Value() string

	// IsSet returns whether the environment variable is defined
	IsSet() bool

	// SetContextValue replaces the environment variable value within the context
	// only i.e. it does NOT change the OS-provided process memory area. Future
	// calls to os.Getenv returns the current value
	SetContextValue(value string)
}

type envVarContext struct {
	key   string
	value string
	isSet bool
}

func (envVarCtx *envVarContext) Key() string {
	return envVarCtx.key
}

func (envVarCtx *envVarContext) Value() string {
	return envVarCtx.value
}

func (envVarCtx *envVarContext) IsSet() bool {
	return envVarCtx.isSet
}

func (envVarCtx *envVarContext) Set(value string) {
	os.Setenv(envVarCtx.key, value)
	envVarCtx.isSet = true
	envVarCtx.value = value
}

func (envVarCtx *envVarContext) SetContextValue(value string) {
	envVarCtx.value = value
}

func (envVarCtx *envVarContext) Get(options ...EnvVarOptionFn) string {
	envVarCtx.value, envVarCtx.isSet = os.LookupEnv(envVarCtx.key)

	for _, option := range options {
		option(envVarCtx)
	}

	return envVarCtx.value
}

type EnvVarOptionFn func(envCtx EnvVarContext)

func WithValueIfUnset(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.IsSet() {
			return
		}

		envVarCtx.SetContextValue(value)
	}
}

func WithValueIfEmpty(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.Value() != "" {
			return
		}

		envVarCtx.SetContextValue(value)
	}
}

func WithMustBeSet() EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.IsSet() {
			return
		}

		panic(fmt.Errorf("environment variable %s must be set", envVarCtx.Key()))
	}
}

func WithMustBeNonEmpty() EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.Value() != "" {
			return
		}

		panic(fmt.Errorf("environment variable %s must not be empty", envVarCtx.Key()))
	}
}

func WithValueIfSet(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if !envVarCtx.IsSet() {
			return
		}

		envVarCtx.SetContextValue(value)
	}
}

func WithValueIfNonEmpty(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if !envVarCtx.IsSet() || envVarCtx.Value() == "" {
			return
		}

		envVarCtx.SetContextValue(value)
	}
}

func WithUpdateValueIfUnset(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.IsSet() {
			return
		}

		envVarCtx.Set(value)
	}
}

func WithUpdateValueIfEmpty(value string) EnvVarOptionFn {
	return func(envVarCtx EnvVarContext) {
		if envVarCtx.IsSet() && envVarCtx.Value() != "" {
			return
		}

		envVarCtx.Set(value)
	}
}
