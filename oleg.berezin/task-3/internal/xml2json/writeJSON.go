package xml2json

import (
	"encoding/json"
	"errors"
)

func WriteJSON(valCurs []Format) ([]byte, error) {
	jsonData, err := json.MarshalIndent(valCurs, "", "    ")
	if err != nil {
		return []byte{}, errors.New("error during writing json")
	}

	return jsonData, nil
}
