package jsonemitter

import (
	"encoding/json"
	"io"

	"github.com/realFrogboy/task-3/internal/cursesparser"
)

func Emit(outputFile io.Writer, valutes []cursesparser.Valute) error {
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
