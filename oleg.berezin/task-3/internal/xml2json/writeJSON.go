package xml2json

import "encoding/json"

func WriteJSON(valCurs []Format) []byte {
	jsonData, err := json.MarshalIndent(valCurs, "", "    ")
	if err != nil {
		panic("Error during marshal json")
	}

	return jsonData
}
