package jsonemitter

import (
	"encoding/json"
	"os"

	"github.com/realFrogboy/task-3/internal/cursesparser"
)

func Emit(outputFile *os.File, valutes []cursesparser.Valute) error {
	data, err := json.Marshal(valutes)
	if err != nil {
		return err
	}

	_, err = outputFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}
