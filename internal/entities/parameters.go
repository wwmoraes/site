package entities

import (
	"fmt"
)

type Parameters map[string]any

func (params Parameters) GetString(key string) (string, error) {
	raw, ok := params[key]
	if !ok {
		return "", nil
	}

	value, ok := raw.(string)
	if !ok {
		return "", fmt.Errorf("value isn't a valid string")
	}

	return value, nil
}

func (params Parameters) GetInt64(key string) (int64, error) {
	raw, ok := params[key]
	if !ok {
		return 0, nil
	}

	value, ok := raw.(int64)
	if !ok {
		return 0, fmt.Errorf("value isn't a valid int64")
	}

	return value, nil
}
