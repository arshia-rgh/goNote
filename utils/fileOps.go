package utils

import (
	"bufio"
	"encoding/json"
	"goNote/note"
	"os"
	"strings"
)

func WriteToFile(fileName string, note note.Note) error {

	fileName = ensureJsonSuffix(fileName)

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.Marshal(note)

	if err != nil {
		return err
	}

	_, err = file.Write(append(data, '\n'))

	if err != nil {
		return err
	}
	return nil
}

//func ReadFromFile(fileName string) ([]note.Note, error) {
//	fileName = ensureJsonSuffix(fileName)
//
//	data, err := os.ReadFile(fileName)
//
//	if err != nil {
//		return nil, err
//	}
//
//	var notes []note.Note
//
//	entries := strings.Split(string(data), "\n")
//
//	for _, entry := range entries {
//		if entry == "" {
//			continue
//		}
//		var n note.Note
//		err = json.Unmarshal([]byte(entry), &n)
//		if err != nil {
//			return nil, err
//		}
//
//		notes = append(notes, n)
//	}
//
//	return notes, nil
//}

func ReadFromFile(fileName string) ([]note.Note, error) {
	file, err := os.Open(fileName)

	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	var notes []note.Note

	for scanner.Scan() {
		var n note.Note
		err := json.Unmarshal([]byte(scanner.Text()), &n)
		if err != nil {
			return nil, err
		}
		notes = append(notes)

	}

	return notes, nil
}

func ensureJsonSuffix(fileName string) string {
	if !strings.HasSuffix(fileName, ".json") {
		fileName += ".json"
	}

	return fileName
}
