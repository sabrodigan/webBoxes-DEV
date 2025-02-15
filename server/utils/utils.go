package utils

import (
	"encoding/json"
)

func MapToStruct(data map[string]interface{}, result interface{}) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = json.Unmarshal(jsonData, result)
	if err != nil {
		return err
	}
	return nil
}
