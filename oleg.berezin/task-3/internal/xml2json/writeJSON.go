package xml2json

import "encoding/json"

func writeJSON(valCurs Format) []byte {
	jsonData, err := json.Marshal(valCurs)
	if err != nil {
		panic("Error during marshal json")
	}

	return jsonData
}
