package encoding

import (
	"encoding/json"
	"io"
)

func EncodeJson(w io.Writer, data interface{}) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(data)
}

func DecodeJson(r io.Reader) (interface{}, error) {
	var data map[string]interface{}
	err := json.NewDecoder(r).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
