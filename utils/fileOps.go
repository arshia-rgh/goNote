package utils

import (
	"encoding/json"
	"goNote/note"
	"os"
	"strings"
)

func WriteToFile(fileName string, note note.Note) error {

	fileName = ensureJsonSuffix(fileName)
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

func ReadFromFile(fileName string) (note.Note, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return note.Note{}, err
	}

	var n note.Note

	err = json.Unmarshal(data, &n)
	if err != nil {
		return note.Note{}, err
	}

	return n, nil
}

func ensureJsonSuffix(fileName string) string {
	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	return fileName
}
