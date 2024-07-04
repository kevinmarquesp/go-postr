package utils

import "encoding/json"

func JsonMarshalString(input interface{}) (string, error) {
	result, err := json.Marshal(input)
	if err != nil {
		return "", nil
	}

	return string(result), nil
}
