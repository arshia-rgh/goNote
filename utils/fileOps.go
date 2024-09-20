package utils

import (
	"encoding/json"
	"goNote/note"
	"os"
)

func WriteToFile(fileName string, note note.Note) error {
	data, err := json.Marshal(note)

	if err != nil {
		return err
	}

	err = os.WriteFile(fileName, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func ReadFromFile() {

}
