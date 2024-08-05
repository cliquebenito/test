package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"test/internal/models"
)

func GetInfoForId(ids []int) ([]models.Data, error) {
	err := convertCsvtoJson()
	if err != nil {
		return []models.Data{}, err
	}

	filename := "newdata.json"
	file, err := os.Open(filename)
	if err != nil {
		return []models.Data{}, err
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return []models.Data{}, err
	}
	var people []models.Data
	err = json.Unmarshal(data, &people)
	if err != nil {
		return []models.Data{}, err
	}
	var result []models.Data
	for _, person := range people {
		for _, id := range ids {
			if person.ID == id {
				result = append(result, person)
				break
			}
		}
	}

	if len(result) == 0 {
		return []models.Data{}, errors.New("Info not found")
	}

	return result, nil
}
