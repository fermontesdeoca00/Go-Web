package loadJson

import (
	"encoding/json"
	"os"
)

func LoadDataFromFile(filePath string, inter interface{}) error {
	// open json file
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// create a decoder for the json data
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(inter); err != nil {
		return err
	}
	return nil
}
